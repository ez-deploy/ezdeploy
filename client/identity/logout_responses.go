// Code generated by go-swagger; DO NOT EDIT.

package identity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ez-deploy/ezdeploy/models"
)

// LogoutReader is a Reader for the Logout structure.
type LogoutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LogoutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLogoutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewLogoutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewLogoutInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLogoutOK creates a LogoutOK with default headers values
func NewLogoutOK() *LogoutOK {
	return &LogoutOK{}
}

/* LogoutOK describes a response with status code 200, with default header values.

Logout Success
*/
type LogoutOK struct {
}

func (o *LogoutOK) Error() string {
	return fmt.Sprintf("[GET /user/logout][%d] logoutOK ", 200)
}

func (o *LogoutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLogoutUnauthorized creates a LogoutUnauthorized with default headers values
func NewLogoutUnauthorized() *LogoutUnauthorized {
	return &LogoutUnauthorized{}
}

/* LogoutUnauthorized describes a response with status code 401, with default header values.

Logout Failed, no login
*/
type LogoutUnauthorized struct {
	Payload *models.Error
}

func (o *LogoutUnauthorized) Error() string {
	return fmt.Sprintf("[GET /user/logout][%d] logoutUnauthorized  %+v", 401, o.Payload)
}
func (o *LogoutUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *LogoutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLogoutInternalServerError creates a LogoutInternalServerError with default headers values
func NewLogoutInternalServerError() *LogoutInternalServerError {
	return &LogoutInternalServerError{}
}

/* LogoutInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type LogoutInternalServerError struct {
	Payload *models.Error
}

func (o *LogoutInternalServerError) Error() string {
	return fmt.Sprintf("[GET /user/logout][%d] logoutInternalServerError  %+v", 500, o.Payload)
}
func (o *LogoutInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *LogoutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
