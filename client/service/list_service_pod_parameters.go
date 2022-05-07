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

// NewListServicePodParams creates a new ListServicePodParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListServicePodParams() *ListServicePodParams {
	return &ListServicePodParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListServicePodParamsWithTimeout creates a new ListServicePodParams object
// with the ability to set a timeout on a request.
func NewListServicePodParamsWithTimeout(timeout time.Duration) *ListServicePodParams {
	return &ListServicePodParams{
		timeout: timeout,
	}
}

// NewListServicePodParamsWithContext creates a new ListServicePodParams object
// with the ability to set a context for a request.
func NewListServicePodParamsWithContext(ctx context.Context) *ListServicePodParams {
	return &ListServicePodParams{
		Context: ctx,
	}
}

// NewListServicePodParamsWithHTTPClient creates a new ListServicePodParams object
// with the ability to set a custom HTTPClient for a request.
func NewListServicePodParamsWithHTTPClient(client *http.Client) *ListServicePodParams {
	return &ListServicePodParams{
		HTTPClient: client,
	}
}

/* ListServicePodParams contains all the parameters to send to the API endpoint
   for the list service pod operation.

   Typically these are written to a http.Request.
*/
type ListServicePodParams struct {

	// ServiceID.
	ServiceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list service pod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListServicePodParams) WithDefaults() *ListServicePodParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list service pod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListServicePodParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list service pod params
func (o *ListServicePodParams) WithTimeout(timeout time.Duration) *ListServicePodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list service pod params
func (o *ListServicePodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list service pod params
func (o *ListServicePodParams) WithContext(ctx context.Context) *ListServicePodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list service pod params
func (o *ListServicePodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list service pod params
func (o *ListServicePodParams) WithHTTPClient(client *http.Client) *ListServicePodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list service pod params
func (o *ListServicePodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithServiceID adds the serviceID to the list service pod params
func (o *ListServicePodParams) WithServiceID(serviceID int64) *ListServicePodParams {
	o.SetServiceID(serviceID)
	return o
}

// SetServiceID adds the serviceId to the list service pod params
func (o *ListServicePodParams) SetServiceID(serviceID int64) {
	o.ServiceID = serviceID
}

// WriteToRequest writes these params to a swagger request
func (o *ListServicePodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param service_id
	qrServiceID := o.ServiceID
	qServiceID := swag.FormatInt64(qrServiceID)
	if qServiceID != "" {

		if err := r.SetQueryParam("service_id", qServiceID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
