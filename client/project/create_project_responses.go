// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ez-deploy/ezdeploy/models"
)

// CreateProjectReader is a Reader for the CreateProject structure.
type CreateProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateProjectCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewCreateProjectConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateProjectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateProjectCreated creates a CreateProjectCreated with default headers values
func NewCreateProjectCreated() *CreateProjectCreated {
	return &CreateProjectCreated{}
}

/* CreateProjectCreated describes a response with status code 201, with default header values.

Create Project Success, return project info.
*/
type CreateProjectCreated struct {
	Payload *models.ProjectInfo
}

func (o *CreateProjectCreated) Error() string {
	return fmt.Sprintf("[POST /project/create][%d] createProjectCreated  %+v", 201, o.Payload)
}
func (o *CreateProjectCreated) GetPayload() *models.ProjectInfo {
	return o.Payload
}

func (o *CreateProjectCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProjectInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProjectConflict creates a CreateProjectConflict with default headers values
func NewCreateProjectConflict() *CreateProjectConflict {
	return &CreateProjectConflict{}
}

/* CreateProjectConflict describes a response with status code 409, with default header values.

Create Project Failed, cause project exist
*/
type CreateProjectConflict struct {
	Payload *models.Error
}

func (o *CreateProjectConflict) Error() string {
	return fmt.Sprintf("[POST /project/create][%d] createProjectConflict  %+v", 409, o.Payload)
}
func (o *CreateProjectConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateProjectConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProjectInternalServerError creates a CreateProjectInternalServerError with default headers values
func NewCreateProjectInternalServerError() *CreateProjectInternalServerError {
	return &CreateProjectInternalServerError{}
}

/* CreateProjectInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CreateProjectInternalServerError struct {
	Payload *models.Error
}

func (o *CreateProjectInternalServerError) Error() string {
	return fmt.Sprintf("[POST /project/create][%d] createProjectInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateProjectInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateProjectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
