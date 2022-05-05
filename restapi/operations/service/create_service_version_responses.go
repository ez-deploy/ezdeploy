// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/ez-deploy/ezdeploy/models"
)

// CreateServiceVersionCreatedCode is the HTTP code returned for type CreateServiceVersionCreated
const CreateServiceVersionCreatedCode int = 201

/*CreateServiceVersionCreated Create Service Version Success, return service version info.

swagger:response createServiceVersionCreated
*/
type CreateServiceVersionCreated struct {

	/*
	  In: Body
	*/
	Payload *models.ServiceVersion `json:"body,omitempty"`
}

// NewCreateServiceVersionCreated creates CreateServiceVersionCreated with default headers values
func NewCreateServiceVersionCreated() *CreateServiceVersionCreated {

	return &CreateServiceVersionCreated{}
}

// WithPayload adds the payload to the create service version created response
func (o *CreateServiceVersionCreated) WithPayload(payload *models.ServiceVersion) *CreateServiceVersionCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create service version created response
func (o *CreateServiceVersionCreated) SetPayload(payload *models.ServiceVersion) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateServiceVersionCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateServiceVersionForbiddenCode is the HTTP code returned for type CreateServiceVersionForbidden
const CreateServiceVersionForbiddenCode int = 403

/*CreateServiceVersionForbidden Create Service Version Failed, cause do not have permisssion

swagger:response createServiceVersionForbidden
*/
type CreateServiceVersionForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateServiceVersionForbidden creates CreateServiceVersionForbidden with default headers values
func NewCreateServiceVersionForbidden() *CreateServiceVersionForbidden {

	return &CreateServiceVersionForbidden{}
}

// WithPayload adds the payload to the create service version forbidden response
func (o *CreateServiceVersionForbidden) WithPayload(payload *models.Error) *CreateServiceVersionForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create service version forbidden response
func (o *CreateServiceVersionForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateServiceVersionForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateServiceVersionConflictCode is the HTTP code returned for type CreateServiceVersionConflict
const CreateServiceVersionConflictCode int = 409

/*CreateServiceVersionConflict Create Service Version Failed, cause service version exist

swagger:response createServiceVersionConflict
*/
type CreateServiceVersionConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateServiceVersionConflict creates CreateServiceVersionConflict with default headers values
func NewCreateServiceVersionConflict() *CreateServiceVersionConflict {

	return &CreateServiceVersionConflict{}
}

// WithPayload adds the payload to the create service version conflict response
func (o *CreateServiceVersionConflict) WithPayload(payload *models.Error) *CreateServiceVersionConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create service version conflict response
func (o *CreateServiceVersionConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateServiceVersionConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateServiceVersionInternalServerErrorCode is the HTTP code returned for type CreateServiceVersionInternalServerError
const CreateServiceVersionInternalServerErrorCode int = 500

/*CreateServiceVersionInternalServerError Server Error

swagger:response createServiceVersionInternalServerError
*/
type CreateServiceVersionInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateServiceVersionInternalServerError creates CreateServiceVersionInternalServerError with default headers values
func NewCreateServiceVersionInternalServerError() *CreateServiceVersionInternalServerError {

	return &CreateServiceVersionInternalServerError{}
}

// WithPayload adds the payload to the create service version internal server error response
func (o *CreateServiceVersionInternalServerError) WithPayload(payload *models.Error) *CreateServiceVersionInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create service version internal server error response
func (o *CreateServiceVersionInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateServiceVersionInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}