// Code generated by go-swagger; DO NOT EDIT.

package pod

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/ez-deploy/ezdeploy/models"
)

// CreatePodTicketHandlerFunc turns a function with the right signature into a create pod ticket handler
type CreatePodTicketHandlerFunc func(CreatePodTicketParams, *models.AuthInfo) middleware.Responder

// Handle executing the request and returning a response
func (fn CreatePodTicketHandlerFunc) Handle(params CreatePodTicketParams, principal *models.AuthInfo) middleware.Responder {
	return fn(params, principal)
}

// CreatePodTicketHandler interface for that can handle valid create pod ticket params
type CreatePodTicketHandler interface {
	Handle(CreatePodTicketParams, *models.AuthInfo) middleware.Responder
}

// NewCreatePodTicket creates a new http.Handler for the create pod ticket operation
func NewCreatePodTicket(ctx *middleware.Context, handler CreatePodTicketHandler) *CreatePodTicket {
	return &CreatePodTicket{Context: ctx, Handler: handler}
}

/* CreatePodTicket swagger:route POST /visit/pod/ticket/create Pod createPodTicket

Create Visit Pod Once Ticket

*/
type CreatePodTicket struct {
	Context *middleware.Context
	Handler CreatePodTicketHandler
}

func (o *CreatePodTicket) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreatePodTicketParams()
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