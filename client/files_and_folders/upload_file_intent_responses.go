// Code generated by go-swagger; DO NOT EDIT.

package files_and_folders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mattx/verizon-cloud-go/models"
)

// UploadFileIntentReader is a Reader for the UploadFileIntent structure.
type UploadFileIntentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadFileIntentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUploadFileIntentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUploadFileIntentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUploadFileIntentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUploadFileIntentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewUploadFileIntentRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewUploadFileIntentServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUploadFileIntentOK creates a UploadFileIntentOK with default headers values
func NewUploadFileIntentOK() *UploadFileIntentOK {
	return &UploadFileIntentOK{}
}

/* UploadFileIntentOK describes a response with status code 200, with default header values.

OK
*/
type UploadFileIntentOK struct {
	Payload *models.Uploadurls
}

func (o *UploadFileIntentOK) Error() string {
	return fmt.Sprintf("[GET /fileupload/intent][%d] uploadFileIntentOK  %+v", 200, o.Payload)
}
func (o *UploadFileIntentOK) GetPayload() *models.Uploadurls {
	return o.Payload
}

func (o *UploadFileIntentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Uploadurls)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadFileIntentBadRequest creates a UploadFileIntentBadRequest with default headers values
func NewUploadFileIntentBadRequest() *UploadFileIntentBadRequest {
	return &UploadFileIntentBadRequest{}
}

/* UploadFileIntentBadRequest describes a response with status code 400, with default header values.

[Bad Request] Query parameters missing or invalid.
*/
type UploadFileIntentBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *UploadFileIntentBadRequest) Error() string {
	return fmt.Sprintf("[GET /fileupload/intent][%d] uploadFileIntentBadRequest  %+v", 400, o.Payload)
}
func (o *UploadFileIntentBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UploadFileIntentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadFileIntentUnauthorized creates a UploadFileIntentUnauthorized with default headers values
func NewUploadFileIntentUnauthorized() *UploadFileIntentUnauthorized {
	return &UploadFileIntentUnauthorized{}
}

/* UploadFileIntentUnauthorized describes a response with status code 401, with default header values.

[Unauthorized] Bearer token is missing, expired, or invalid.
*/
type UploadFileIntentUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *UploadFileIntentUnauthorized) Error() string {
	return fmt.Sprintf("[GET /fileupload/intent][%d] uploadFileIntentUnauthorized  %+v", 401, o.Payload)
}
func (o *UploadFileIntentUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UploadFileIntentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadFileIntentForbidden creates a UploadFileIntentForbidden with default headers values
func NewUploadFileIntentForbidden() *UploadFileIntentForbidden {
	return &UploadFileIntentForbidden{}
}

/* UploadFileIntentForbidden describes a response with status code 403, with default header values.

[Forbidden] User is not authorized to access storage APIs.
*/
type UploadFileIntentForbidden struct {
	Payload *models.ErrorResponse
}

func (o *UploadFileIntentForbidden) Error() string {
	return fmt.Sprintf("[GET /fileupload/intent][%d] uploadFileIntentForbidden  %+v", 403, o.Payload)
}
func (o *UploadFileIntentForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UploadFileIntentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadFileIntentRequestEntityTooLarge creates a UploadFileIntentRequestEntityTooLarge with default headers values
func NewUploadFileIntentRequestEntityTooLarge() *UploadFileIntentRequestEntityTooLarge {
	return &UploadFileIntentRequestEntityTooLarge{}
}

/* UploadFileIntentRequestEntityTooLarge describes a response with status code 413, with default header values.

[Request Too Long] User quota exceeded.
*/
type UploadFileIntentRequestEntityTooLarge struct {
	Payload *models.ErrorResponse
}

func (o *UploadFileIntentRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /fileupload/intent][%d] uploadFileIntentRequestEntityTooLarge  %+v", 413, o.Payload)
}
func (o *UploadFileIntentRequestEntityTooLarge) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UploadFileIntentRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadFileIntentServiceUnavailable creates a UploadFileIntentServiceUnavailable with default headers values
func NewUploadFileIntentServiceUnavailable() *UploadFileIntentServiceUnavailable {
	return &UploadFileIntentServiceUnavailable{}
}

/* UploadFileIntentServiceUnavailable describes a response with status code 503, with default header values.

[Service Unavailable] See response body for more detail.
*/
type UploadFileIntentServiceUnavailable struct {
	Payload *models.ErrorResponse
}

func (o *UploadFileIntentServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /fileupload/intent][%d] uploadFileIntentServiceUnavailable  %+v", 503, o.Payload)
}
func (o *UploadFileIntentServiceUnavailable) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UploadFileIntentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
