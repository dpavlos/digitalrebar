package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new networks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for networks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DELETENetwork deletes network
*/
func (a *Client) DELETENetwork(params *DELETENetworkParams) (*DELETENetworkNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDELETENetworkParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DELETE-network",
		Method:             "DELETE",
		PathPattern:        "/networks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DELETENetworkReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DELETENetworkNoContent), nil

}

/*
GETNetwork gets network
*/
func (a *Client) GETNetwork(params *GETNetworkParams) (*GETNetworkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETNetworkParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GET-network",
		Method:             "GET",
		PathPattern:        "/networks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GETNetworkReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GETNetworkOK), nil

}

/*
LISTNetworks lists networks
*/
func (a *Client) LISTNetworks(params *LISTNetworksParams) (*LISTNetworksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLISTNetworksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "LIST-networks",
		Method:             "GET",
		PathPattern:        "/networks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LISTNetworksReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*LISTNetworksOK), nil

}

/*
POSTNetwork creates network
*/
func (a *Client) POSTNetwork(params *POSTNetworkParams) (*POSTNetworkCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPOSTNetworkParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "POST-network",
		Method:             "POST",
		PathPattern:        "/networks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &POSTNetworkReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*POSTNetworkCreated), nil

}

/*
PUTNetwork updates network
*/
func (a *Client) PUTNetwork(params *PUTNetworkParams) (*PUTNetworkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPUTNetworkParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PUT-network",
		Method:             "PUT",
		PathPattern:        "/networks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PUTNetworkReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PUTNetworkOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
