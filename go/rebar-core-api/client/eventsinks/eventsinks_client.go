package eventsinks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new eventsinks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for eventsinks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DELETEEventsink deletes event sink
*/
func (a *Client) DELETEEventsink(params *DELETEEventsinkParams) (*DELETEEventsinkNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDELETEEventsinkParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DELETE-eventsink",
		Method:             "DELETE",
		PathPattern:        "/eventsinks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DELETEEventsinkReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DELETEEventsinkNoContent), nil

}

/*
GETEventsink gets event sink
*/
func (a *Client) GETEventsink(params *GETEventsinkParams) (*GETEventsinkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETEventsinkParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GET-eventsink",
		Method:             "GET",
		PathPattern:        "/eventsinks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GETEventsinkReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GETEventsinkOK), nil

}

/*
LISTEventsinks lists eventsinks
*/
func (a *Client) LISTEventsinks(params *LISTEventsinksParams) (*LISTEventsinksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLISTEventsinksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "LIST-eventsinks",
		Method:             "GET",
		PathPattern:        "/eventsinks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LISTEventsinksReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*LISTEventsinksOK), nil

}

/*
POSTEventsink creates event sink
*/
func (a *Client) POSTEventsink(params *POSTEventsinkParams) (*POSTEventsinkCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPOSTEventsinkParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "POST-eventsink",
		Method:             "POST",
		PathPattern:        "/eventsinks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &POSTEventsinkReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*POSTEventsinkCreated), nil

}

/*
PUTEventsink updates event sink
*/
func (a *Client) PUTEventsink(params *PUTEventsinkParams) (*PUTEventsinkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPUTEventsinkParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PUT-eventsink",
		Method:             "PUT",
		PathPattern:        "/eventsinks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PUTEventsinkReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PUTEventsinkOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
