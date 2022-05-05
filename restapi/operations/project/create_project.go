// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/ez-deploy/ezdeploy/models"
)

// CreateProjectHandlerFunc turns a function with the right signature into a create project handler
type CreateProjectHandlerFunc func(CreateProjectParams, *models.AuthInfo) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateProjectHandlerFunc) Handle(params CreateProjectParams, principal *models.AuthInfo) middleware.Responder {
	return fn(params, principal)
}

// CreateProjectHandler interface for that can handle valid create project params
type CreateProjectHandler interface {
	Handle(CreateProjectParams, *models.AuthInfo) middleware.Responder
}

// NewCreateProject creates a new http.Handler for the create project operation
func NewCreateProject(ctx *middleware.Context, handler CreateProjectHandler) *CreateProject {
	return &CreateProject{Context: ctx, Handler: handler}
}

/* CreateProject swagger:route POST /project/create Project createProject

Create Project

*/
type CreateProject struct {
	Context *middleware.Context
	Handler CreateProjectHandler
}

func (o *CreateProject) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateProjectParams()
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
