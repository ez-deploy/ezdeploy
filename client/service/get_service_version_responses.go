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

// GetServiceVersionReader is a Reader for the GetServiceVersion structure.
type GetServiceVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServiceVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServiceVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetServiceVersionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetServiceVersionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetServiceVersionOK creates a GetServiceVersionOK with default headers values
func NewGetServiceVersionOK() *GetServiceVersionOK {
	return &GetServiceVersionOK{}
}

/* GetServiceVersionOK describes a response with status code 200, with default header values.

List Service Version Success, return service version info.
*/
type GetServiceVersionOK struct {
	Payload *models.ServiceVersion
}

func (o *GetServiceVersionOK) Error() string {
	return fmt.Sprintf("[GET /service/version/get][%d] getServiceVersionOK  %+v", 200, o.Payload)
}
func (o *GetServiceVersionOK) GetPayload() *models.ServiceVersion {
	return o.Payload
}

func (o *GetServiceVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceVersion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceVersionForbidden creates a GetServiceVersionForbidden with default headers values
func NewGetServiceVersionForbidden() *GetServiceVersionForbidden {
	return &GetServiceVersionForbidden{}
}

/* GetServiceVersionForbidden describes a response with status code 403, with default header values.

List Service Version Failed, cause do not have permisssion
*/
type GetServiceVersionForbidden struct {
	Payload *models.Error
}

func (o *GetServiceVersionForbidden) Error() string {
	return fmt.Sprintf("[GET /service/version/get][%d] getServiceVersionForbidden  %+v", 403, o.Payload)
}
func (o *GetServiceVersionForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetServiceVersionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceVersionInternalServerError creates a GetServiceVersionInternalServerError with default headers values
func NewGetServiceVersionInternalServerError() *GetServiceVersionInternalServerError {
	return &GetServiceVersionInternalServerError{}
}

/* GetServiceVersionInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GetServiceVersionInternalServerError struct {
	Payload *models.Error
}

func (o *GetServiceVersionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /service/version/get][%d] getServiceVersionInternalServerError  %+v", 500, o.Payload)
}
func (o *GetServiceVersionInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetServiceVersionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}