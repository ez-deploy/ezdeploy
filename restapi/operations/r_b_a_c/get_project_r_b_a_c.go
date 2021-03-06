// Code generated by go-swagger; DO NOT EDIT.

package r_b_a_c

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/ez-deploy/ezdeploy/models"
)

// GetProjectRBACHandlerFunc turns a function with the right signature into a get project r b a c handler
type GetProjectRBACHandlerFunc func(GetProjectRBACParams, *models.AuthInfo) middleware.Responder

// Handle executing the request and returning a response
func (fn GetProjectRBACHandlerFunc) Handle(params GetProjectRBACParams, principal *models.AuthInfo) middleware.Responder {
	return fn(params, principal)
}

// GetProjectRBACHandler interface for that can handle valid get project r b a c params
type GetProjectRBACHandler interface {
	Handle(GetProjectRBACParams, *models.AuthInfo) middleware.Responder
}

// NewGetProjectRBAC creates a new http.Handler for the get project r b a c operation
func NewGetProjectRBAC(ctx *middleware.Context, handler GetProjectRBACHandler) *GetProjectRBAC {
	return &GetProjectRBAC{Context: ctx, Handler: handler}
}

/* GetProjectRBAC swagger:route GET /rbac/project/get RBAC getProjectRBAC

Get Project RBAC

*/
type GetProjectRBAC struct {
	Context *middleware.Context
	Handler GetProjectRBACHandler
}

func (o *GetProjectRBAC) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetProjectRBACParams()
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
