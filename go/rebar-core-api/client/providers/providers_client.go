package providers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new providers API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for providers API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DELETEProvider deletes provider
*/
func (a *Client) DELETEProvider(params *DELETEProviderParams) (*DELETEProviderNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDELETEProviderParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DELETE-provider",
		Method:             "DELETE",
		PathPattern:        "/providers/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DELETEProviderReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DELETEProviderNoContent), nil

}

/*
GETProvider gets provider
*/
func (a *Client) GETProvider(params *GETProviderParams) (*GETProviderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETProviderParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GET-provider",
		Method:             "GET",
		PathPattern:        "/providers/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GETProviderReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GETProviderOK), nil

}

/*
LISTProviders lists providers
*/
func (a *Client) LISTProviders(params *LISTProvidersParams) (*LISTProvidersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLISTProvidersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "LIST-providers",
		Method:             "GET",
		PathPattern:        "/providers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LISTProvidersReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*LISTProvidersOK), nil

}

/*
POSTProvider creates provider
*/
func (a *Client) POSTProvider(params *POSTProviderParams) (*POSTProviderCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPOSTProviderParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "POST-provider",
		Method:             "POST",
		PathPattern:        "/providers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &POSTProviderReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*POSTProviderCreated), nil

}

/*
PUTProvider updates provider
*/
func (a *Client) PUTProvider(params *PUTProviderParams) (*PUTProviderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPUTProviderParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PUT-provider",
		Method:             "PUT",
		PathPattern:        "/providers/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PUTProviderReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PUTProviderOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
