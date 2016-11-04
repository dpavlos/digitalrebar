package eventselectors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new eventselectors API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for eventselectors API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DELETEEventselector deletes event selector
*/
func (a *Client) DELETEEventselector(params *DELETEEventselectorParams) (*DELETEEventselectorNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDELETEEventselectorParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DELETE-eventselector",
		Method:             "DELETE",
		PathPattern:        "/eventselectors/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DELETEEventselectorReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DELETEEventselectorNoContent), nil

}

/*
GETEventselector gets event selector
*/
func (a *Client) GETEventselector(params *GETEventselectorParams) (*GETEventselectorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETEventselectorParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GET-eventselector",
		Method:             "GET",
		PathPattern:        "/eventselectors/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GETEventselectorReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GETEventselectorOK), nil

}

/*
LISTEventselectors lists eventselectors
*/
func (a *Client) LISTEventselectors(params *LISTEventselectorsParams) (*LISTEventselectorsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLISTEventselectorsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "LIST-eventselectors",
		Method:             "GET",
		PathPattern:        "/eventselectors",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LISTEventselectorsReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*LISTEventselectorsOK), nil

}

/*
POSTEventselector creates event selector
*/
func (a *Client) POSTEventselector(params *POSTEventselectorParams) (*POSTEventselectorCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPOSTEventselectorParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "POST-eventselector",
		Method:             "POST",
		PathPattern:        "/eventselectors",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &POSTEventselectorReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*POSTEventselectorCreated), nil

}

/*
PUTEventselector updates event selector
*/
func (a *Client) PUTEventselector(params *PUTEventselectorParams) (*PUTEventselectorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPUTEventselectorParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PUT-eventselector",
		Method:             "PUT",
		PathPattern:        "/eventselectors/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PUTEventselectorReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PUTEventselectorOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
