#!/bin/bash

# Build initial access keys
if ! rebar -U rebar -P rebar1 users show machine-install; then
    # read key rest < <(dd if=/dev/urandom bs=64 count=1 2>/dev/null |sha512sum - 2>/dev/null)
    key=`dd if=/dev/urandom bs=64 count=1 2>/dev/null | sha512sum - 2>/dev/null | awk '{ print $1 }'`
    echo "Creating machine-install user"
    machine_user="
{
  \"username\": \"machine-install\",
  \"email\": \"root@localhost.localdomain\",
  \"password\": \"$key\",
  \"password_confirmation\": \"$key\",
  \"remember_me\": false,
  \"is_admin\": true,
  \"digest\": true
}"

    if ! rebar -U rebar -P rebar1 -E https://127.0.0.1:3000 users import "$machine_user"; then
        echo "Could not create machine-install user!"
        exit 1
    fi
    echo "machine-install:$key" >/etc/rebar.install.key
    kv_put digitalrebar/private/api/keys/machine_key </etc/rebar.install.key
else
    kv_get digitalrebar/private/api/keys/machine_key >/etc/rebar.install.key
fi
export REBAR_KEY="$(cat /etc/rebar.install.key)"
