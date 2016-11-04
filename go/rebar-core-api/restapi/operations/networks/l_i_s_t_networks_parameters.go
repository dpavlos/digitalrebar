package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
)

// NewLISTNetworksParams creates a new LISTNetworksParams object
// with the default values initialized.
func NewLISTNetworksParams() LISTNetworksParams {
	var ()
	return LISTNetworksParams{}
}

// LISTNetworksParams contains all the bound params for the l i s t networks operation
// typically these are obtained from a http.Request
//
// swagger:parameters LIST-networks
type LISTNetworksParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *LISTNetworksParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
