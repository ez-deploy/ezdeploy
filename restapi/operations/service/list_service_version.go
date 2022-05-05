// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/ez-deploy/ezdeploy/models"
)

// ListServiceVersionHandlerFunc turns a function with the right signature into a list service version handler
type ListServiceVersionHandlerFunc func(ListServiceVersionParams, *models.AuthInfo) middleware.Responder

// Handle executing the request and returning a response
func (fn ListServiceVersionHandlerFunc) Handle(params ListServiceVersionParams, principal *models.AuthInfo) middleware.Responder {
	return fn(params, principal)
}

// ListServiceVersionHandler interface for that can handle valid list service version params
type ListServiceVersionHandler interface {
	Handle(ListServiceVersionParams, *models.AuthInfo) middleware.Responder
}

// NewListServiceVersion creates a new http.Handler for the list service version operation
func NewListServiceVersion(ctx *middleware.Context, handler ListServiceVersionHandler) *ListServiceVersion {
	return &ListServiceVersion{Context: ctx, Handler: handler}
}

/* ListServiceVersion swagger:route GET /service/version/list Service listServiceVersion

List Service Version by service ID.

*/
type ListServiceVersion struct {
	Context *middleware.Context
	Handler ListServiceVersionHandler
}

func (o *ListServiceVersion) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewListServiceVersionParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.AuthInfo
	if uprinc != nil {
		principal = uprinc.(*models.AuthInfo) // this is really a models.AuthInfo, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// ListServiceVersionOKBody list service version o k body
//
// swagger:model ListServiceVersionOKBody
type ListServiceVersionOKBody struct {

	// versions
	Versions []*models.ServiceVersion `json:"versions"`
}

// Validate validates this list service version o k body
func (o *ListServiceVersionOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateVersions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListServiceVersionOKBody) validateVersions(formats strfmt.Registry) error {
	if swag.IsZero(o.Versions) { // not required
		return nil
	}

	for i := 0; i < len(o.Versions); i++ {
		if swag.IsZero(o.Versions[i]) { // not required
			continue
		}

		if o.Versions[i] != nil {
			if err := o.Versions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listServiceVersionOK" + "." + "versions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listServiceVersionOK" + "." + "versions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list service version o k body based on the context it is used
func (o *ListServiceVersionOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateVersions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListServiceVersionOKBody) contextValidateVersions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Versions); i++ {

		if o.Versions[i] != nil {
			if err := o.Versions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listServiceVersionOK" + "." + "versions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listServiceVersionOK" + "." + "versions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListServiceVersionOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListServiceVersionOKBody) UnmarshalBinary(b []byte) error {
	var res ListServiceVersionOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
