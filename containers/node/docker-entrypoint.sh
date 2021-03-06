#!/bin/bash

count=0
while [[ ! -e /etc/rebar-data/rebar-key.sh ]] ; do
  if [ $((count % 12)) -eq 0 ] ; then
      echo "Waiting for rebar-key.sh to show up"
  fi
  sleep 5
  count=$((count+1))
done

# Wait for the webserver to be ready.
. /etc/rebar-data/rebar-key.sh
count=0
while ! rebar ping; do
  sleep 1
  if [ $((count % 60)) -eq 0 ] ; then
      echo "Waiting for functional rebar-key.sh"
  fi
  . /etc/rebar-data/rebar-key.sh
  count=$((count+1))
done

service ssh start

# This could error if this is the first time.  Ignore it
#set +e
# Node id is harcoded here, and that is a Bad Thing
if blob="$(rebar deployments get 1 attrib rebar-access_keys)"; then
  mkdir -p /root/.ssh
  touch /root/.ssh/authorized_keys
  awk -F\" '{ print $4 }' <<<"$blob" >> /root/.ssh/authorized_keys
  chmod 700 /root/.ssh/authorized_keys
fi

#set -e
ip_re='([0-9a-f.:]+/[0-9]+)'
if ! [[ $(ip -4 -o addr show |grep 'scope global' |grep -v ' lo' |grep -v ' dynamic') =~ $ip_re ]]; then
  echo "Cannot find IP address for the admin node!"
  exit 1
fi
IPADDR="${BASH_REMATCH[1]}"
echo "Using $IPADDR for this host"

DOMAIN="$(rebar nodes get "system-phantom.internal.local" attrib dns-domain | jq -r .value)"
if [ $DOMAIN == "null" ] ; then
  echo "Domain must be set to something"
  exit 1
fi

HOSTNAME=$(hostname).$DOMAIN

# Reset the default route to our forwarder - only needed for docker containers.
if [ "$FORWARDER_IP" != "" ] ; then
    forwarder=$(awk '/forwarder/ { print $1 }' </etc/hosts | sort -u)
    ip route del default
    ip route add default via $forwarder
fi

# Let the other nodes come up
# GREG: wait for it - check instead
sleep 120

# Add node to DigitalRebar
if ! rebar nodes show "$HOSTNAME"; then
  # Create a new node for us,
  # Let the annealer do its thing.
  rebar nodes import "{\"name\": \"$HOSTNAME\", \"admin\": false, \"ip\": \"$IPADDR\", \"bootenv\": \"local\"}"|| {
    echo "We could not create a node for ourself!"
    #exit 1
  }
else
  echo "Node already created, moving on"
fi

# does the rebar-joined-role exist?
if ! rebar nodes roles $HOSTNAME  |grep -q 'rebar-joined-node'; then
  rebar nodes bind "$HOSTNAME" to 'rebar-joined-node'
  rebar nodes commit "$HOSTNAME" || {
    echo "We could not commit the node!"
    #exit 1
  }
else
  echo "Node already committed, moving on"
fi

# Always make sure we are marking the node alive. It will comeback later.
rebar nodes update "$HOSTNAME" "{\"alive\": true, \"bootenv\": \"local\"}"

tail -f /var/log/*
