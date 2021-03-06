// Code generated by go-swagger; DO NOT EDIT.

package r_b_a_c

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/ez-deploy/ezdeploy/models"
)

// GetUserRBACHandlerFunc turns a function with the right signature into a get user r b a c handler
type GetUserRBACHandlerFunc func(GetUserRBACParams, *models.AuthInfo) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUserRBACHandlerFunc) Handle(params GetUserRBACParams, principal *models.AuthInfo) middleware.Responder {
	return fn(params, principal)
}

// GetUserRBACHandler interface for that can handle valid get user r b a c params
type GetUserRBACHandler interface {
	Handle(GetUserRBACParams, *models.AuthInfo) middleware.Responder
}

// NewGetUserRBAC creates a new http.Handler for the get user r b a c operation
func NewGetUserRBAC(ctx *middleware.Context, handler GetUserRBACHandler) *GetUserRBAC {
	return &GetUserRBAC{Context: ctx, Handler: handler}
}

/* GetUserRBAC swagger:route GET /rbac/user/get RBAC getUserRBAC

Get User RBAC

*/
type GetUserRBAC struct {
	Context *middleware.Context
	Handler GetUserRBACHandler
}

func (o *GetUserRBAC) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetUserRBACParams()
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
