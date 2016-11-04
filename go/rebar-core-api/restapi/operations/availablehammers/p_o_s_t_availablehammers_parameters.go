package availablehammers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/models"
)

// NewPOSTAvailablehammersParams creates a new POSTAvailablehammersParams object
// with the default values initialized.
func NewPOSTAvailablehammersParams() POSTAvailablehammersParams {
	var ()
	return POSTAvailablehammersParams{}
}

// POSTAvailablehammersParams contains all the bound params for the p o s t availablehammers operation
// typically these are obtained from a http.Request
//
// swagger:parameters POST-availablehammers
type POSTAvailablehammersParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*
	  In: body
	*/
	Body *models.AvailablehammersInput
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *POSTAvailablehammersParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.AvailablehammersInput
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("body", "body", "", err))
		} else {
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
