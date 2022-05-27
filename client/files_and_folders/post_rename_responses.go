// Code generated by go-swagger; DO NOT EDIT.

package files_and_folders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/MattX/verizon-cloud-go/models"
)

// PostRenameReader is a Reader for the PostRename structure.
type PostRenameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRenameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostRenameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostRenameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostRenameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostRenameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostRenameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostRenameConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostRenameServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostRenameOK creates a PostRenameOK with default headers values
func NewPostRenameOK() *PostRenameOK {
	return &PostRenameOK{}
}

/* PostRenameOK describes a response with status code 200, with default header values.

OK
*/
type PostRenameOK struct {
	Payload *models.MetadataResponse
}

func (o *PostRenameOK) Error() string {
	return fmt.Sprintf("[POST /fops/rename][%d] postRenameOK  %+v", 200, o.Payload)
}
func (o *PostRenameOK) GetPayload() *models.MetadataResponse {
	return o.Payload
}

func (o *PostRenameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetadataResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRenameBadRequest creates a PostRenameBadRequest with default headers values
func NewPostRenameBadRequest() *PostRenameBadRequest {
	return &PostRenameBadRequest{}
}

/* PostRenameBadRequest describes a response with status code 400, with default header values.

[Bad Request] Query parameters missing or invalid.
*/
type PostRenameBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *PostRenameBadRequest) Error() string {
	return fmt.Sprintf("[POST /fops/rename][%d] postRenameBadRequest  %+v", 400, o.Payload)
}
func (o *PostRenameBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PostRenameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRenameUnauthorized creates a PostRenameUnauthorized with default headers values
func NewPostRenameUnauthorized() *PostRenameUnauthorized {
	return &PostRenameUnauthorized{}
}

/* PostRenameUnauthorized describes a response with status code 401, with default header values.

[Unauthorized] Bearer token is missing, expired, or invalid.
*/
type PostRenameUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *PostRenameUnauthorized) Error() string {
	return fmt.Sprintf("[POST /fops/rename][%d] postRenameUnauthorized  %+v", 401, o.Payload)
}
func (o *PostRenameUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PostRenameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRenameForbidden creates a PostRenameForbidden with default headers values
func NewPostRenameForbidden() *PostRenameForbidden {
	return &PostRenameForbidden{}
}

/* PostRenameForbidden describes a response with status code 403, with default header values.

[Forbidden] User is not authorized to access storage APIs.
*/
type PostRenameForbidden struct {
	Payload *models.ErrorResponse
}

func (o *PostRenameForbidden) Error() string {
	return fmt.Sprintf("[POST /fops/rename][%d] postRenameForbidden  %+v", 403, o.Payload)
}
func (o *PostRenameForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PostRenameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRenameNotFound creates a PostRenameNotFound with default headers values
func NewPostRenameNotFound() *PostRenameNotFound {
	return &PostRenameNotFound{}
}

/* PostRenameNotFound describes a response with status code 404, with default header values.

[Not Found] File was not found.
*/
type PostRenameNotFound struct {
	Payload *models.ErrorResponse
}

func (o *PostRenameNotFound) Error() string {
	return fmt.Sprintf("[POST /fops/rename][%d] postRenameNotFound  %+v", 404, o.Payload)
}
func (o *PostRenameNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PostRenameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRenameConflict creates a PostRenameConflict with default headers values
func NewPostRenameConflict() *PostRenameConflict {
	return &PostRenameConflict{}
}

/* PostRenameConflict describes a response with status code 409, with default header values.

Conflict.
*/
type PostRenameConflict struct {
	Payload *models.ErrorResponse
}

func (o *PostRenameConflict) Error() string {
	return fmt.Sprintf("[POST /fops/rename][%d] postRenameConflict  %+v", 409, o.Payload)
}
func (o *PostRenameConflict) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PostRenameConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRenameServiceUnavailable creates a PostRenameServiceUnavailable with default headers values
func NewPostRenameServiceUnavailable() *PostRenameServiceUnavailable {
	return &PostRenameServiceUnavailable{}
}

/* PostRenameServiceUnavailable describes a response with status code 503, with default header values.

[Service Unavailable] See response body for more detail.
*/
type PostRenameServiceUnavailable struct {
	Payload *models.ErrorResponse
}

func (o *PostRenameServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /fops/rename][%d] postRenameServiceUnavailable  %+v", 503, o.Payload)
}
func (o *PostRenameServiceUnavailable) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PostRenameServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
