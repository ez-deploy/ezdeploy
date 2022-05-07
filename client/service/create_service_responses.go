// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ez-deploy/ezdeploy/models"
)

// CreateServiceReader is a Reader for the CreateService structure.
type CreateServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateServiceCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCreateServiceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateServiceConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateServiceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateServiceCreated creates a CreateServiceCreated with default headers values
func NewCreateServiceCreated() *CreateServiceCreated {
	return &CreateServiceCreated{}
}

/* CreateServiceCreated describes a response with status code 201, with default header values.

Create Service Success, return service info.
*/
type CreateServiceCreated struct {
	Payload *models.ServiceInfo
}

func (o *CreateServiceCreated) Error() string {
	return fmt.Sprintf("[POST /service/create][%d] createServiceCreated  %+v", 201, o.Payload)
}
func (o *CreateServiceCreated) GetPayload() *models.ServiceInfo {
	return o.Payload
}

func (o *CreateServiceCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateServiceForbidden creates a CreateServiceForbidden with default headers values
func NewCreateServiceForbidden() *CreateServiceForbidden {
	return &CreateServiceForbidden{}
}

/* CreateServiceForbidden describes a response with status code 403, with default header values.

Create Service Failed, cause do not have permisssion
*/
type CreateServiceForbidden struct {
	Payload *models.Error
}

func (o *CreateServiceForbidden) Error() string {
	return fmt.Sprintf("[POST /service/create][%d] createServiceForbidden  %+v", 403, o.Payload)
}
func (o *CreateServiceForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateServiceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateServiceConflict creates a CreateServiceConflict with default headers values
func NewCreateServiceConflict() *CreateServiceConflict {
	return &CreateServiceConflict{}
}

/* CreateServiceConflict describes a response with status code 409, with default header values.

Create Service Failed, cause service exist
*/
type CreateServiceConflict struct {
	Payload *models.Error
}

func (o *CreateServiceConflict) Error() string {
	return fmt.Sprintf("[POST /service/create][%d] createServiceConflict  %+v", 409, o.Payload)
}
func (o *CreateServiceConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateServiceConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateServiceInternalServerError creates a CreateServiceInternalServerError with default headers values
func NewCreateServiceInternalServerError() *CreateServiceInternalServerError {
	return &CreateServiceInternalServerError{}
}

/* CreateServiceInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CreateServiceInternalServerError struct {
	Payload *models.Error
}

func (o *CreateServiceInternalServerError) Error() string {
	return fmt.Sprintf("[POST /service/create][%d] createServiceInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateServiceInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateServiceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
