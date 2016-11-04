package profiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new profiles API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for profiles API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DELETEProfile deletes profile
*/
func (a *Client) DELETEProfile(params *DELETEProfileParams) (*DELETEProfileNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDELETEProfileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DELETE-profile",
		Method:             "DELETE",
		PathPattern:        "/profiles/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DELETEProfileReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DELETEProfileNoContent), nil

}

/*
GETProfile gets profile
*/
func (a *Client) GETProfile(params *GETProfileParams) (*GETProfileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETProfileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GET-profile",
		Method:             "GET",
		PathPattern:        "/profiles/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GETProfileReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GETProfileOK), nil

}

/*
LISTProfiles lists profiles
*/
func (a *Client) LISTProfiles(params *LISTProfilesParams) (*LISTProfilesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLISTProfilesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "LIST-profiles",
		Method:             "GET",
		PathPattern:        "/profiles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LISTProfilesReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*LISTProfilesOK), nil

}

/*
POSTProfile creates profile
*/
func (a *Client) POSTProfile(params *POSTProfileParams) (*POSTProfileCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPOSTProfileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "POST-profile",
		Method:             "POST",
		PathPattern:        "/profiles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &POSTProfileReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*POSTProfileCreated), nil

}

/*
PUTProfile updates profile
*/
func (a *Client) PUTProfile(params *PUTProfileParams) (*PUTProfileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPUTProfileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PUT-profile",
		Method:             "PUT",
		PathPattern:        "/profiles/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PUTProfileReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PUTProfileOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
