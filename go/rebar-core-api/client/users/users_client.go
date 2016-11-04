package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new users API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for users API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DELETEUser deletes user
*/
func (a *Client) DELETEUser(params *DELETEUserParams) (*DELETEUserNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDELETEUserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DELETE-user",
		Method:             "DELETE",
		PathPattern:        "/users/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DELETEUserReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DELETEUserNoContent), nil

}

/*
GETUser gets user
*/
func (a *Client) GETUser(params *GETUserParams) (*GETUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETUserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GET-user",
		Method:             "GET",
		PathPattern:        "/users/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GETUserReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GETUserOK), nil

}

/*
LISTUsers lists users
*/
func (a *Client) LISTUsers(params *LISTUsersParams) (*LISTUsersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLISTUsersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "LIST-users",
		Method:             "GET",
		PathPattern:        "/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LISTUsersReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*LISTUsersOK), nil

}

/*
POSTUser creates user
*/
func (a *Client) POSTUser(params *POSTUserParams) (*POSTUserCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPOSTUserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "POST-user",
		Method:             "POST",
		PathPattern:        "/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &POSTUserReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*POSTUserCreated), nil

}

/*
PUTUser updates user
*/
func (a *Client) PUTUser(params *PUTUserParams) (*PUTUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPUTUserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PUT-user",
		Method:             "PUT",
		PathPattern:        "/users/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PUTUserReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PUTUserOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
