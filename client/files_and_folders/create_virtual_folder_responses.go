// Code generated by go-swagger; DO NOT EDIT.

package files_and_folders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mattx/verizon-cloud-go/client/models"
)

// CreateVirtualFolderReader is a Reader for the CreateVirtualFolder structure.
type CreateVirtualFolderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateVirtualFolderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateVirtualFolderCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateVirtualFolderBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateVirtualFolderUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateVirtualFolderForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewCreateVirtualFolderServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateVirtualFolderCreated creates a CreateVirtualFolderCreated with default headers values
func NewCreateVirtualFolderCreated() *CreateVirtualFolderCreated {
	return &CreateVirtualFolderCreated{}
}

/* CreateVirtualFolderCreated describes a response with status code 201, with default header values.

A virtual folder response object
*/
type CreateVirtualFolderCreated struct {
	Payload *models.VirtualfolderResponse
}

func (o *CreateVirtualFolderCreated) Error() string {
	return fmt.Sprintf("[POST /virtualfolder/{name}][%d] createVirtualFolderCreated  %+v", 201, o.Payload)
}
func (o *CreateVirtualFolderCreated) GetPayload() *models.VirtualfolderResponse {
	return o.Payload
}

func (o *CreateVirtualFolderCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualfolderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVirtualFolderBadRequest creates a CreateVirtualFolderBadRequest with default headers values
func NewCreateVirtualFolderBadRequest() *CreateVirtualFolderBadRequest {
	return &CreateVirtualFolderBadRequest{}
}

/* CreateVirtualFolderBadRequest describes a response with status code 400, with default header values.

[Bad Request] Query parameters missing or invalid.
*/
type CreateVirtualFolderBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *CreateVirtualFolderBadRequest) Error() string {
	return fmt.Sprintf("[POST /virtualfolder/{name}][%d] createVirtualFolderBadRequest  %+v", 400, o.Payload)
}
func (o *CreateVirtualFolderBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateVirtualFolderBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVirtualFolderUnauthorized creates a CreateVirtualFolderUnauthorized with default headers values
func NewCreateVirtualFolderUnauthorized() *CreateVirtualFolderUnauthorized {
	return &CreateVirtualFolderUnauthorized{}
}

/* CreateVirtualFolderUnauthorized describes a response with status code 401, with default header values.

[Unauthorized] Bearer token is missing, expired, or invalid.
*/
type CreateVirtualFolderUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CreateVirtualFolderUnauthorized) Error() string {
	return fmt.Sprintf("[POST /virtualfolder/{name}][%d] createVirtualFolderUnauthorized  %+v", 401, o.Payload)
}
func (o *CreateVirtualFolderUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateVirtualFolderUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVirtualFolderForbidden creates a CreateVirtualFolderForbidden with default headers values
func NewCreateVirtualFolderForbidden() *CreateVirtualFolderForbidden {
	return &CreateVirtualFolderForbidden{}
}

/* CreateVirtualFolderForbidden describes a response with status code 403, with default header values.

[Forbidden] User is not authorized to access storage APIs.
*/
type CreateVirtualFolderForbidden struct {
	Payload *models.ErrorResponse
}

func (o *CreateVirtualFolderForbidden) Error() string {
	return fmt.Sprintf("[POST /virtualfolder/{name}][%d] createVirtualFolderForbidden  %+v", 403, o.Payload)
}
func (o *CreateVirtualFolderForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateVirtualFolderForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVirtualFolderServiceUnavailable creates a CreateVirtualFolderServiceUnavailable with default headers values
func NewCreateVirtualFolderServiceUnavailable() *CreateVirtualFolderServiceUnavailable {
	return &CreateVirtualFolderServiceUnavailable{}
}

/* CreateVirtualFolderServiceUnavailable describes a response with status code 503, with default header values.

[Service Unavailable] See response body for more detail.
*/
type CreateVirtualFolderServiceUnavailable struct {
	Payload *models.ErrorResponse
}

func (o *CreateVirtualFolderServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /virtualfolder/{name}][%d] createVirtualFolderServiceUnavailable  %+v", 503, o.Payload)
}
func (o *CreateVirtualFolderServiceUnavailable) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateVirtualFolderServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
