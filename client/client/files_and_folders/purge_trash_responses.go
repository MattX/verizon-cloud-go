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

// PurgeTrashReader is a Reader for the PurgeTrash structure.
type PurgeTrashReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PurgeTrashReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPurgeTrashOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPurgeTrashBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPurgeTrashUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPurgeTrashForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPurgeTrashNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPurgeTrashServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPurgeTrashOK creates a PurgeTrashOK with default headers values
func NewPurgeTrashOK() *PurgeTrashOK {
	return &PurgeTrashOK{}
}

/* PurgeTrashOK describes a response with status code 200, with default header values.

OK
*/
type PurgeTrashOK struct {
}

func (o *PurgeTrashOK) Error() string {
	return fmt.Sprintf("[DELETE /trash][%d] purgeTrashOK ", 200)
}

func (o *PurgeTrashOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPurgeTrashBadRequest creates a PurgeTrashBadRequest with default headers values
func NewPurgeTrashBadRequest() *PurgeTrashBadRequest {
	return &PurgeTrashBadRequest{}
}

/* PurgeTrashBadRequest describes a response with status code 400, with default header values.

[Bad Request]
*/
type PurgeTrashBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *PurgeTrashBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /trash][%d] purgeTrashBadRequest  %+v", 400, o.Payload)
}
func (o *PurgeTrashBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PurgeTrashBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPurgeTrashUnauthorized creates a PurgeTrashUnauthorized with default headers values
func NewPurgeTrashUnauthorized() *PurgeTrashUnauthorized {
	return &PurgeTrashUnauthorized{}
}

/* PurgeTrashUnauthorized describes a response with status code 401, with default header values.

[Unauthorized] Bearer token is missing, expired, or invalid.
*/
type PurgeTrashUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *PurgeTrashUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /trash][%d] purgeTrashUnauthorized  %+v", 401, o.Payload)
}
func (o *PurgeTrashUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PurgeTrashUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPurgeTrashForbidden creates a PurgeTrashForbidden with default headers values
func NewPurgeTrashForbidden() *PurgeTrashForbidden {
	return &PurgeTrashForbidden{}
}

/* PurgeTrashForbidden describes a response with status code 403, with default header values.

[Forbidden] User is not authorized to access storage APIs.
*/
type PurgeTrashForbidden struct {
	Payload *models.ErrorResponse
}

func (o *PurgeTrashForbidden) Error() string {
	return fmt.Sprintf("[DELETE /trash][%d] purgeTrashForbidden  %+v", 403, o.Payload)
}
func (o *PurgeTrashForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PurgeTrashForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPurgeTrashNotFound creates a PurgeTrashNotFound with default headers values
func NewPurgeTrashNotFound() *PurgeTrashNotFound {
	return &PurgeTrashNotFound{}
}

/* PurgeTrashNotFound describes a response with status code 404, with default header values.

[Not Found] Resource not found.
*/
type PurgeTrashNotFound struct {
	Payload *models.ErrorResponse
}

func (o *PurgeTrashNotFound) Error() string {
	return fmt.Sprintf("[DELETE /trash][%d] purgeTrashNotFound  %+v", 404, o.Payload)
}
func (o *PurgeTrashNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PurgeTrashNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPurgeTrashServiceUnavailable creates a PurgeTrashServiceUnavailable with default headers values
func NewPurgeTrashServiceUnavailable() *PurgeTrashServiceUnavailable {
	return &PurgeTrashServiceUnavailable{}
}

/* PurgeTrashServiceUnavailable describes a response with status code 503, with default header values.

[Service Unavailable] See response body for more detail.
*/
type PurgeTrashServiceUnavailable struct {
	Payload *models.ErrorResponse
}

func (o *PurgeTrashServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /trash][%d] purgeTrashServiceUnavailable  %+v", 503, o.Payload)
}
func (o *PurgeTrashServiceUnavailable) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PurgeTrashServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}