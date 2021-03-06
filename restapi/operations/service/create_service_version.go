// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/ez-deploy/ezdeploy/models"
)

// CreateServiceVersionHandlerFunc turns a function with the right signature into a create service version handler
type CreateServiceVersionHandlerFunc func(CreateServiceVersionParams, *models.AuthInfo) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateServiceVersionHandlerFunc) Handle(params CreateServiceVersionParams, principal *models.AuthInfo) middleware.Responder {
	return fn(params, principal)
}

// CreateServiceVersionHandler interface for that can handle valid create service version params
type CreateServiceVersionHandler interface {
	Handle(CreateServiceVersionParams, *models.AuthInfo) middleware.Responder
}

// NewCreateServiceVersion creates a new http.Handler for the create service version operation
func NewCreateServiceVersion(ctx *middleware.Context, handler CreateServiceVersionHandler) *CreateServiceVersion {
	return &CreateServiceVersion{Context: ctx, Handler: handler}
}

/* CreateServiceVersion swagger:route POST /service/version/create Service createServiceVersion

Create Service Version

*/
type CreateServiceVersion struct {
	Context *middleware.Context
	Handler CreateServiceVersionHandler
}

func (o *CreateServiceVersion) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateServiceVersionParams()
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
