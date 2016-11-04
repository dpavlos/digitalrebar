package deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDELETEDeploymentsParams creates a new DELETEDeploymentsParams object
// with the default values initialized.
func NewDELETEDeploymentsParams() DELETEDeploymentsParams {
	var ()
	return DELETEDeploymentsParams{}
}

// DELETEDeploymentsParams contains all the bound params for the d e l e t e deployments operation
// typically these are obtained from a http.Request
//
// swagger:parameters DELETE-deployments
type DELETEDeploymentsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*
	  Required: true
	  In: path
	*/
	ID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *DELETEDeploymentsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DELETEDeploymentsParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.ID = raw

	return nil
}
