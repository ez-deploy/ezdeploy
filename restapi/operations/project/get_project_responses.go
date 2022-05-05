// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/ez-deploy/ezdeploy/models"
)

// GetProjectOKCode is the HTTP code returned for type GetProjectOK
const GetProjectOKCode int = 200

/*GetProjectOK List All Projects Success, return project info. (project name, project id, project description)

swagger:response getProjectOK
*/
type GetProjectOK struct {

	/*
	  In: Body
	*/
	Payload *models.ProjectInfo `json:"body,omitempty"`
}

// NewGetProjectOK creates GetProjectOK with default headers values
func NewGetProjectOK() *GetProjectOK {

	return &GetProjectOK{}
}

// WithPayload adds the payload to the get project o k response
func (o *GetProjectOK) WithPayload(payload *models.ProjectInfo) *GetProjectOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get project o k response
func (o *GetProjectOK) SetPayload(payload *models.ProjectInfo) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProjectOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetProjectNotFoundCode is the HTTP code returned for type GetProjectNotFound
const GetProjectNotFoundCode int = 404

/*GetProjectNotFound Get Project Failed, cause project not exist

swagger:response getProjectNotFound
*/
type GetProjectNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetProjectNotFound creates GetProjectNotFound with default headers values
func NewGetProjectNotFound() *GetProjectNotFound {

	return &GetProjectNotFound{}
}

// WithPayload adds the payload to the get project not found response
func (o *GetProjectNotFound) WithPayload(payload *models.Error) *GetProjectNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get project not found response
func (o *GetProjectNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProjectNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetProjectInternalServerErrorCode is the HTTP code returned for type GetProjectInternalServerError
const GetProjectInternalServerErrorCode int = 500

/*GetProjectInternalServerError Server Error

swagger:response getProjectInternalServerError
*/
type GetProjectInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetProjectInternalServerError creates GetProjectInternalServerError with default headers values
func NewGetProjectInternalServerError() *GetProjectInternalServerError {

	return &GetProjectInternalServerError{}
}

// WithPayload adds the payload to the get project internal server error response
func (o *GetProjectInternalServerError) WithPayload(payload *models.Error) *GetProjectInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get project internal server error response
func (o *GetProjectInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProjectInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
