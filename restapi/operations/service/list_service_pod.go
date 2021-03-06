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

// ListServicePodHandlerFunc turns a function with the right signature into a list service pod handler
type ListServicePodHandlerFunc func(ListServicePodParams, *models.AuthInfo) middleware.Responder

// Handle executing the request and returning a response
func (fn ListServicePodHandlerFunc) Handle(params ListServicePodParams, principal *models.AuthInfo) middleware.Responder {
	return fn(params, principal)
}

// ListServicePodHandler interface for that can handle valid list service pod params
type ListServicePodHandler interface {
	Handle(ListServicePodParams, *models.AuthInfo) middleware.Responder
}

// NewListServicePod creates a new http.Handler for the list service pod operation
func NewListServicePod(ctx *middleware.Context, handler ListServicePodHandler) *ListServicePod {
	return &ListServicePod{Context: ctx, Handler: handler}
}

/* ListServicePod swagger:route GET /service/pod/list Service listServicePod

List Service Pod by service ID.

*/
type ListServicePod struct {
	Context *middleware.Context
	Handler ListServicePodHandler
}

func (o *ListServicePod) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewListServicePodParams()
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

// ListServicePodOKBody list service pod o k body
//
// swagger:model ListServicePodOKBody
type ListServicePodOKBody struct {

	// pods
	Pods []*models.PodInfo `json:"pods"`
}

// Validate validates this list service pod o k body
func (o *ListServicePodOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePods(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListServicePodOKBody) validatePods(formats strfmt.Registry) error {
	if swag.IsZero(o.Pods) { // not required
		return nil
	}

	for i := 0; i < len(o.Pods); i++ {
		if swag.IsZero(o.Pods[i]) { // not required
			continue
		}

		if o.Pods[i] != nil {
			if err := o.Pods[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listServicePodOK" + "." + "pods" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listServicePodOK" + "." + "pods" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list service pod o k body based on the context it is used
func (o *ListServicePodOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePods(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListServicePodOKBody) contextValidatePods(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Pods); i++ {

		if o.Pods[i] != nil {
			if err := o.Pods[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listServicePodOK" + "." + "pods" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listServicePodOK" + "." + "pods" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListServicePodOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListServicePodOKBody) UnmarshalBinary(b []byte) error {
	var res ListServicePodOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
