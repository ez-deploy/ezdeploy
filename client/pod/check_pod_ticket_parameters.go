// Code generated by go-swagger; DO NOT EDIT.

package pod

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewCheckPodTicketParams creates a new CheckPodTicketParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCheckPodTicketParams() *CheckPodTicketParams {
	return &CheckPodTicketParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCheckPodTicketParamsWithTimeout creates a new CheckPodTicketParams object
// with the ability to set a timeout on a request.
func NewCheckPodTicketParamsWithTimeout(timeout time.Duration) *CheckPodTicketParams {
	return &CheckPodTicketParams{
		timeout: timeout,
	}
}

// NewCheckPodTicketParamsWithContext creates a new CheckPodTicketParams object
// with the ability to set a context for a request.
func NewCheckPodTicketParamsWithContext(ctx context.Context) *CheckPodTicketParams {
	return &CheckPodTicketParams{
		Context: ctx,
	}
}

// NewCheckPodTicketParamsWithHTTPClient creates a new CheckPodTicketParams object
// with the ability to set a custom HTTPClient for a request.
func NewCheckPodTicketParamsWithHTTPClient(client *http.Client) *CheckPodTicketParams {
	return &CheckPodTicketParams{
		HTTPClient: client,
	}
}

/* CheckPodTicketParams contains all the parameters to send to the API endpoint
   for the check pod ticket operation.

   Typically these are written to a http.Request.
*/
type CheckPodTicketParams struct {

	// TicketValue.
	TicketValue string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the check pod ticket params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CheckPodTicketParams) WithDefaults() *CheckPodTicketParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the check pod ticket params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CheckPodTicketParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the check pod ticket params
func (o *CheckPodTicketParams) WithTimeout(timeout time.Duration) *CheckPodTicketParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the check pod ticket params
func (o *CheckPodTicketParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the check pod ticket params
func (o *CheckPodTicketParams) WithContext(ctx context.Context) *CheckPodTicketParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the check pod ticket params
func (o *CheckPodTicketParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the check pod ticket params
func (o *CheckPodTicketParams) WithHTTPClient(client *http.Client) *CheckPodTicketParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the check pod ticket params
func (o *CheckPodTicketParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTicketValue adds the ticketValue to the check pod ticket params
func (o *CheckPodTicketParams) WithTicketValue(ticketValue string) *CheckPodTicketParams {
	o.SetTicketValue(ticketValue)
	return o
}

// SetTicketValue adds the ticketValue to the check pod ticket params
func (o *CheckPodTicketParams) SetTicketValue(ticketValue string) {
	o.TicketValue = ticketValue
}

// WriteToRequest writes these params to a swagger request
func (o *CheckPodTicketParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param ticket_value
	qrTicketValue := o.TicketValue
	qTicketValue := qrTicketValue
	if qTicketValue != "" {

		if err := r.SetQueryParam("ticket_value", qTicketValue); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
