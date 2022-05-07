// Code generated by go-swagger; DO NOT EDIT.

package service

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
	"github.com/go-openapi/swag"
)

// NewDeleteServiceParams creates a new DeleteServiceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteServiceParams() *DeleteServiceParams {
	return &DeleteServiceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteServiceParamsWithTimeout creates a new DeleteServiceParams object
// with the ability to set a timeout on a request.
func NewDeleteServiceParamsWithTimeout(timeout time.Duration) *DeleteServiceParams {
	return &DeleteServiceParams{
		timeout: timeout,
	}
}

// NewDeleteServiceParamsWithContext creates a new DeleteServiceParams object
// with the ability to set a context for a request.
func NewDeleteServiceParamsWithContext(ctx context.Context) *DeleteServiceParams {
	return &DeleteServiceParams{
		Context: ctx,
	}
}

// NewDeleteServiceParamsWithHTTPClient creates a new DeleteServiceParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteServiceParamsWithHTTPClient(client *http.Client) *DeleteServiceParams {
	return &DeleteServiceParams{
		HTTPClient: client,
	}
}

/* DeleteServiceParams contains all the parameters to send to the API endpoint
   for the delete service operation.

   Typically these are written to a http.Request.
*/
type DeleteServiceParams struct {

	// ID.
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete service params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteServiceParams) WithDefaults() *DeleteServiceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete service params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteServiceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete service params
func (o *DeleteServiceParams) WithTimeout(timeout time.Duration) *DeleteServiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete service params
func (o *DeleteServiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete service params
func (o *DeleteServiceParams) WithContext(ctx context.Context) *DeleteServiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete service params
func (o *DeleteServiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete service params
func (o *DeleteServiceParams) WithHTTPClient(client *http.Client) *DeleteServiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete service params
func (o *DeleteServiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete service params
func (o *DeleteServiceParams) WithID(id int64) *DeleteServiceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete service params
func (o *DeleteServiceParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteServiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param id
	qrID := o.ID
	qID := swag.FormatInt64(qrID)
	if qID != "" {

		if err := r.SetQueryParam("id", qID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
