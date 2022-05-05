// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/ez-deploy/ezdeploy/models"
)

// ListServiceVersionOKCode is the HTTP code returned for type ListServiceVersionOK
const ListServiceVersionOKCode int = 200

/*ListServiceVersionOK List Service Version Success, return service version info.

swagger:response listServiceVersionOK
*/
type ListServiceVersionOK struct {

	/*
	  In: Body
	*/
	Payload []*models.ServiceVersion `json:"body,omitempty"`
}

// NewListServiceVersionOK creates ListServiceVersionOK with default headers values
func NewListServiceVersionOK() *ListServiceVersionOK {

	return &ListServiceVersionOK{}
}

// WithPayload adds the payload to the list service version o k response
func (o *ListServiceVersionOK) WithPayload(payload []*models.ServiceVersion) *ListServiceVersionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list service version o k response
func (o *ListServiceVersionOK) SetPayload(payload []*models.ServiceVersion) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListServiceVersionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.ServiceVersion, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListServiceVersionForbiddenCode is the HTTP code returned for type ListServiceVersionForbidden
const ListServiceVersionForbiddenCode int = 403

/*ListServiceVersionForbidden List Service Version Failed, cause do not have permisssion

swagger:response listServiceVersionForbidden
*/
type ListServiceVersionForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListServiceVersionForbidden creates ListServiceVersionForbidden with default headers values
func NewListServiceVersionForbidden() *ListServiceVersionForbidden {

	return &ListServiceVersionForbidden{}
}

// WithPayload adds the payload to the list service version forbidden response
func (o *ListServiceVersionForbidden) WithPayload(payload *models.Error) *ListServiceVersionForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list service version forbidden response
func (o *ListServiceVersionForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListServiceVersionForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListServiceVersionInternalServerErrorCode is the HTTP code returned for type ListServiceVersionInternalServerError
const ListServiceVersionInternalServerErrorCode int = 500

/*ListServiceVersionInternalServerError Server Error

swagger:response listServiceVersionInternalServerError
*/
type ListServiceVersionInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListServiceVersionInternalServerError creates ListServiceVersionInternalServerError with default headers values
func NewListServiceVersionInternalServerError() *ListServiceVersionInternalServerError {

	return &ListServiceVersionInternalServerError{}
}

// WithPayload adds the payload to the list service version internal server error response
func (o *ListServiceVersionInternalServerError) WithPayload(payload *models.Error) *ListServiceVersionInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list service version internal server error response
func (o *ListServiceVersionInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListServiceVersionInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}