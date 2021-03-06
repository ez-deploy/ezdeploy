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

// ListServiceHandlerFunc turns a function with the right signature into a list service handler
type ListServiceHandlerFunc func(ListServiceParams, *models.AuthInfo) middleware.Responder

// Handle executing the request and returning a response
func (fn ListServiceHandlerFunc) Handle(params ListServiceParams, principal *models.AuthInfo) middleware.Responder {
	return fn(params, principal)
}

// ListServiceHandler interface for that can handle valid list service params
type ListServiceHandler interface {
	Handle(ListServiceParams, *models.AuthInfo) middleware.Responder
}

// NewListService creates a new http.Handler for the list service operation
func NewListService(ctx *middleware.Context, handler ListServiceHandler) *ListService {
	return &ListService{Context: ctx, Handler: handler}
}

/* ListService swagger:route GET /service/list Service listService

List Service by project ID, service ID, service name.

*/
type ListService struct {
	Context *middleware.Context
	Handler ListServiceHandler
}

func (o *ListService) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewListServiceParams()
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

// ListServiceOKBody list service o k body
//
// swagger:model ListServiceOKBody
type ListServiceOKBody struct {

	// services
	Services []*models.ServiceInfo `json:"services"`
}

// Validate validates this list service o k body
func (o *ListServiceOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateServices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListServiceOKBody) validateServices(formats strfmt.Registry) error {
	if swag.IsZero(o.Services) { // not required
		return nil
	}

	for i := 0; i < len(o.Services); i++ {
		if swag.IsZero(o.Services[i]) { // not required
			continue
		}

		if o.Services[i] != nil {
			if err := o.Services[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listServiceOK" + "." + "services" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listServiceOK" + "." + "services" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list service o k body based on the context it is used
func (o *ListServiceOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateServices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListServiceOKBody) contextValidateServices(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Services); i++ {

		if o.Services[i] != nil {
			if err := o.Services[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listServiceOK" + "." + "services" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listServiceOK" + "." + "services" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListServiceOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListServiceOKBody) UnmarshalBinary(b []byte) error {
	var res ListServiceOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
