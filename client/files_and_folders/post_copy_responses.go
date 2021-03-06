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

// PostCopyReader is a Reader for the PostCopy structure.
type PostCopyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCopyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostCopyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostCopyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostCopyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostCopyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostCopyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostCopyServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostCopyOK creates a PostCopyOK with default headers values
func NewPostCopyOK() *PostCopyOK {
	return &PostCopyOK{}
}

/* PostCopyOK describes a response with status code 200, with default header values.

OK
*/
type PostCopyOK struct {
	Payload *models.MetadataResponse
}

func (o *PostCopyOK) Error() string {
	return fmt.Sprintf("[POST /fops/copy][%d] postCopyOK  %+v", 200, o.Payload)
}
func (o *PostCopyOK) GetPayload() *models.MetadataResponse {
	return o.Payload
}

func (o *PostCopyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetadataResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCopyBadRequest creates a PostCopyBadRequest with default headers values
func NewPostCopyBadRequest() *PostCopyBadRequest {
	return &PostCopyBadRequest{}
}

/* PostCopyBadRequest describes a response with status code 400, with default header values.

[Bad Request] Query parameters missing or invalid.
*/
type PostCopyBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *PostCopyBadRequest) Error() string {
	return fmt.Sprintf("[POST /fops/copy][%d] postCopyBadRequest  %+v", 400, o.Payload)
}
func (o *PostCopyBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PostCopyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCopyUnauthorized creates a PostCopyUnauthorized with default headers values
func NewPostCopyUnauthorized() *PostCopyUnauthorized {
	return &PostCopyUnauthorized{}
}

/* PostCopyUnauthorized describes a response with status code 401, with default header values.

[Unauthorized] Bearer token is missing, expired, or invalid.
*/
type PostCopyUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *PostCopyUnauthorized) Error() string {
	return fmt.Sprintf("[POST /fops/copy][%d] postCopyUnauthorized  %+v", 401, o.Payload)
}
func (o *PostCopyUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PostCopyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCopyForbidden creates a PostCopyForbidden with default headers values
func NewPostCopyForbidden() *PostCopyForbidden {
	return &PostCopyForbidden{}
}

/* PostCopyForbidden describes a response with status code 403, with default header values.

[Forbidden] User is not authorized to access storage APIs.
*/
type PostCopyForbidden struct {
	Payload *models.ErrorResponse
}

func (o *PostCopyForbidden) Error() string {
	return fmt.Sprintf("[POST /fops/copy][%d] postCopyForbidden  %+v", 403, o.Payload)
}
func (o *PostCopyForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PostCopyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCopyNotFound creates a PostCopyNotFound with default headers values
func NewPostCopyNotFound() *PostCopyNotFound {
	return &PostCopyNotFound{}
}

/* PostCopyNotFound describes a response with status code 404, with default header values.

[Not Found] File was not found.
*/
type PostCopyNotFound struct {
	Payload *models.ErrorResponse
}

func (o *PostCopyNotFound) Error() string {
	return fmt.Sprintf("[POST /fops/copy][%d] postCopyNotFound  %+v", 404, o.Payload)
}
func (o *PostCopyNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PostCopyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCopyServiceUnavailable creates a PostCopyServiceUnavailable with default headers values
func NewPostCopyServiceUnavailable() *PostCopyServiceUnavailable {
	return &PostCopyServiceUnavailable{}
}

/* PostCopyServiceUnavailable describes a response with status code 503, with default header values.

[Service Unavailable] See response body for more detail.
*/
type PostCopyServiceUnavailable struct {
	Payload *models.ErrorResponse
}

func (o *PostCopyServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /fops/copy][%d] postCopyServiceUnavailable  %+v", 503, o.Payload)
}
func (o *PostCopyServiceUnavailable) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PostCopyServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
