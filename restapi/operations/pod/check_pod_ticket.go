// Code generated by go-swagger; DO NOT EDIT.

package pod

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CheckPodTicketHandlerFunc turns a function with the right signature into a check pod ticket handler
type CheckPodTicketHandlerFunc func(CheckPodTicketParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CheckPodTicketHandlerFunc) Handle(params CheckPodTicketParams) middleware.Responder {
	return fn(params)
}

// CheckPodTicketHandler interface for that can handle valid check pod ticket params
type CheckPodTicketHandler interface {
	Handle(CheckPodTicketParams) middleware.Responder
}

// NewCheckPodTicket creates a new http.Handler for the check pod ticket operation
func NewCheckPodTicket(ctx *middleware.Context, handler CheckPodTicketHandler) *CheckPodTicket {
	return &CheckPodTicket{Context: ctx, Handler: handler}
}

/* CheckPodTicket swagger:route GET /visit/pod/ticket/check Pod checkPodTicket

Check Pod Ticket, and delete it

*/
type CheckPodTicket struct {
	Context *middleware.Context
	Handler CheckPodTicketHandler
}

func (o *CheckPodTicket) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCheckPodTicketParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}