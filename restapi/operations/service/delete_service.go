// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/ez-deploy/ezdeploy/models"
)

// DeleteServiceHandlerFunc turns a function with the right signature into a delete service handler
type DeleteServiceHandlerFunc func(DeleteServiceParams, *models.AuthInfo) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteServiceHandlerFunc) Handle(params DeleteServiceParams, principal *models.AuthInfo) middleware.Responder {
	return fn(params, principal)
}

// DeleteServiceHandler interface for that can handle valid delete service params
type DeleteServiceHandler interface {
	Handle(DeleteServiceParams, *models.AuthInfo) middleware.Responder
}

// NewDeleteService creates a new http.Handler for the delete service operation
func NewDeleteService(ctx *middleware.Context, handler DeleteServiceHandler) *DeleteService {
	return &DeleteService{Context: ctx, Handler: handler}
}

/* DeleteService swagger:route DELETE /service/delete Service deleteService

Delete Service

*/
type DeleteService struct {
	Context *middleware.Context
	Handler DeleteServiceHandler
}

func (o *DeleteService) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteServiceParams()
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
