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

// DeleteServiceReader is a Reader for the DeleteService structure.
type DeleteServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteServiceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteServiceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteServiceOK creates a DeleteServiceOK with default headers values
func NewDeleteServiceOK() *DeleteServiceOK {
	return &DeleteServiceOK{}
}

/* DeleteServiceOK describes a response with status code 200, with default header values.

Delete Service Success, return service info.
*/
type DeleteServiceOK struct {
	Payload *models.ServiceInfo
}

func (o *DeleteServiceOK) Error() string {
	return fmt.Sprintf("[DELETE /service/delete][%d] deleteServiceOK  %+v", 200, o.Payload)
}
func (o *DeleteServiceOK) GetPayload() *models.ServiceInfo {
	return o.Payload
}

func (o *DeleteServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteServiceForbidden creates a DeleteServiceForbidden with default headers values
func NewDeleteServiceForbidden() *DeleteServiceForbidden {
	return &DeleteServiceForbidden{}
}

/* DeleteServiceForbidden describes a response with status code 403, with default header values.

Delete Service Failed, cause do not have permisssion
*/
type DeleteServiceForbidden struct {
	Payload *models.Error
}

func (o *DeleteServiceForbidden) Error() string {
	return fmt.Sprintf("[DELETE /service/delete][%d] deleteServiceForbidden  %+v", 403, o.Payload)
}
func (o *DeleteServiceForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteServiceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteServiceInternalServerError creates a DeleteServiceInternalServerError with default headers values
func NewDeleteServiceInternalServerError() *DeleteServiceInternalServerError {
	return &DeleteServiceInternalServerError{}
}

/* DeleteServiceInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type DeleteServiceInternalServerError struct {
	Payload *models.Error
}

func (o *DeleteServiceInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /service/delete][%d] deleteServiceInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteServiceInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteServiceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}