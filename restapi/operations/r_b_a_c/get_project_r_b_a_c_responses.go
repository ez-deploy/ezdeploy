// Code generated by go-swagger; DO NOT EDIT.

package r_b_a_c

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/ez-deploy/ezdeploy/models"
)

// GetProjectRBACOKCode is the HTTP code returned for type GetProjectRBACOK
const GetProjectRBACOKCode int = 200

/*GetProjectRBACOK List All Projects RBAC Success, return project RBAC info.

swagger:response getProjectRBACOK
*/
type GetProjectRBACOK struct {

	/*
	  In: Body
	*/
	Payload *models.ProjectRole `json:"body,omitempty"`
}

// NewGetProjectRBACOK creates GetProjectRBACOK with default headers values
func NewGetProjectRBACOK() *GetProjectRBACOK {

	return &GetProjectRBACOK{}
}

// WithPayload adds the payload to the get project r b a c o k response
func (o *GetProjectRBACOK) WithPayload(payload *models.ProjectRole) *GetProjectRBACOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get project r b a c o k response
func (o *GetProjectRBACOK) SetPayload(payload *models.ProjectRole) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProjectRBACOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetProjectRBACForbiddenCode is the HTTP code returned for type GetProjectRBACForbidden
const GetProjectRBACForbiddenCode int = 403

/*GetProjectRBACForbidden Get Project RBAC Failed, cause project not exist

swagger:response getProjectRBACForbidden
*/
type GetProjectRBACForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetProjectRBACForbidden creates GetProjectRBACForbidden with default headers values
func NewGetProjectRBACForbidden() *GetProjectRBACForbidden {

	return &GetProjectRBACForbidden{}
}

// WithPayload adds the payload to the get project r b a c forbidden response
func (o *GetProjectRBACForbidden) WithPayload(payload *models.Error) *GetProjectRBACForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get project r b a c forbidden response
func (o *GetProjectRBACForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProjectRBACForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetProjectRBACInternalServerErrorCode is the HTTP code returned for type GetProjectRBACInternalServerError
const GetProjectRBACInternalServerErrorCode int = 500

/*GetProjectRBACInternalServerError Server Error

swagger:response getProjectRBACInternalServerError
*/
type GetProjectRBACInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetProjectRBACInternalServerError creates GetProjectRBACInternalServerError with default headers values
func NewGetProjectRBACInternalServerError() *GetProjectRBACInternalServerError {

	return &GetProjectRBACInternalServerError{}
}

// WithPayload adds the payload to the get project r b a c internal server error response
func (o *GetProjectRBACInternalServerError) WithPayload(payload *models.Error) *GetProjectRBACInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get project r b a c internal server error response
func (o *GetProjectRBACInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProjectRBACInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}