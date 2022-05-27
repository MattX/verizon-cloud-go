// Code generated by go-swagger; DO NOT EDIT.

package files_and_folders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/MattX/verizon-cloud-go/client/models"
)

// RenameVirtualFolderReader is a Reader for the RenameVirtualFolder structure.
type RenameVirtualFolderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RenameVirtualFolderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRenameVirtualFolderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRenameVirtualFolderBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRenameVirtualFolderUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRenameVirtualFolderForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewRenameVirtualFolderServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRenameVirtualFolderOK creates a RenameVirtualFolderOK with default headers values
func NewRenameVirtualFolderOK() *RenameVirtualFolderOK {
	return &RenameVirtualFolderOK{}
}

/* RenameVirtualFolderOK describes a response with status code 200, with default header values.

A virtual folder response object
*/
type RenameVirtualFolderOK struct {
	Payload *models.VirtualfolderResponse
}

func (o *RenameVirtualFolderOK) Error() string {
	return fmt.Sprintf("[PUT /virtualfolder][%d] renameVirtualFolderOK  %+v", 200, o.Payload)
}
func (o *RenameVirtualFolderOK) GetPayload() *models.VirtualfolderResponse {
	return o.Payload
}

func (o *RenameVirtualFolderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualfolderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRenameVirtualFolderBadRequest creates a RenameVirtualFolderBadRequest with default headers values
func NewRenameVirtualFolderBadRequest() *RenameVirtualFolderBadRequest {
	return &RenameVirtualFolderBadRequest{}
}

/* RenameVirtualFolderBadRequest describes a response with status code 400, with default header values.

[Bad Request] Query parameters missing or invalid.
*/
type RenameVirtualFolderBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *RenameVirtualFolderBadRequest) Error() string {
	return fmt.Sprintf("[PUT /virtualfolder][%d] renameVirtualFolderBadRequest  %+v", 400, o.Payload)
}
func (o *RenameVirtualFolderBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *RenameVirtualFolderBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRenameVirtualFolderUnauthorized creates a RenameVirtualFolderUnauthorized with default headers values
func NewRenameVirtualFolderUnauthorized() *RenameVirtualFolderUnauthorized {
	return &RenameVirtualFolderUnauthorized{}
}

/* RenameVirtualFolderUnauthorized describes a response with status code 401, with default header values.

[Unauthorized] Bearer token is missing, expired, or invalid.
*/
type RenameVirtualFolderUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *RenameVirtualFolderUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /virtualfolder][%d] renameVirtualFolderUnauthorized  %+v", 401, o.Payload)
}
func (o *RenameVirtualFolderUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *RenameVirtualFolderUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRenameVirtualFolderForbidden creates a RenameVirtualFolderForbidden with default headers values
func NewRenameVirtualFolderForbidden() *RenameVirtualFolderForbidden {
	return &RenameVirtualFolderForbidden{}
}

/* RenameVirtualFolderForbidden describes a response with status code 403, with default header values.

[Forbidden] User is not authorized to access storage APIs.
*/
type RenameVirtualFolderForbidden struct {
	Payload *models.ErrorResponse
}

func (o *RenameVirtualFolderForbidden) Error() string {
	return fmt.Sprintf("[PUT /virtualfolder][%d] renameVirtualFolderForbidden  %+v", 403, o.Payload)
}
func (o *RenameVirtualFolderForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *RenameVirtualFolderForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRenameVirtualFolderServiceUnavailable creates a RenameVirtualFolderServiceUnavailable with default headers values
func NewRenameVirtualFolderServiceUnavailable() *RenameVirtualFolderServiceUnavailable {
	return &RenameVirtualFolderServiceUnavailable{}
}

/* RenameVirtualFolderServiceUnavailable describes a response with status code 503, with default header values.

[Service Unavailable] See response body for more detail.
*/
type RenameVirtualFolderServiceUnavailable struct {
	Payload *models.ErrorResponse
}

func (o *RenameVirtualFolderServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /virtualfolder][%d] renameVirtualFolderServiceUnavailable  %+v", 503, o.Payload)
}
func (o *RenameVirtualFolderServiceUnavailable) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *RenameVirtualFolderServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
