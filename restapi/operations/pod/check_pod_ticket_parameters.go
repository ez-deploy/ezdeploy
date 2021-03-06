// Code generated by go-swagger; DO NOT EDIT.

package pod

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewCheckPodTicketParams creates a new CheckPodTicketParams object
//
// There are no default values defined in the spec.
func NewCheckPodTicketParams() CheckPodTicketParams {

	return CheckPodTicketParams{}
}

// CheckPodTicketParams contains all the bound params for the check pod ticket operation
// typically these are obtained from a http.Request
//
// swagger:parameters CheckPodTicket
type CheckPodTicketParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: query
	*/
	TicketValue string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCheckPodTicketParams() beforehand.
func (o *CheckPodTicketParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qTicketValue, qhkTicketValue, _ := qs.GetOK("ticket_value")
	if err := o.bindTicketValue(qTicketValue, qhkTicketValue, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindTicketValue binds and validates parameter TicketValue from query.
func (o *CheckPodTicketParams) bindTicketValue(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("ticket_value", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("ticket_value", "query", raw); err != nil {
		return err
	}
	o.TicketValue = raw

	return nil
}
