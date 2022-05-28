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

// GetFilesReader is a Reader for the GetFiles structure.
type GetFilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetFilesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetFilesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFilesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFilesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewGetFilesGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetFilesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFilesOK creates a GetFilesOK with default headers values
func NewGetFilesOK() *GetFilesOK {
	return &GetFilesOK{}
}

/* GetFilesOK describes a response with status code 200, with default header values.

OK
*/
type GetFilesOK struct {
}

func (o *GetFilesOK) Error() string {
	return fmt.Sprintf("[GET /files/{path}][%d] getFilesOK ", 200)
}

func (o *GetFilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFilesBadRequest creates a GetFilesBadRequest with default headers values
func NewGetFilesBadRequest() *GetFilesBadRequest {
	return &GetFilesBadRequest{}
}

/* GetFilesBadRequest describes a response with status code 400, with default header values.

[Bad Request] Query parameters missing or invalid.
*/
type GetFilesBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *GetFilesBadRequest) Error() string {
	return fmt.Sprintf("[GET /files/{path}][%d] getFilesBadRequest  %+v", 400, o.Payload)
}
func (o *GetFilesBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetFilesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFilesUnauthorized creates a GetFilesUnauthorized with default headers values
func NewGetFilesUnauthorized() *GetFilesUnauthorized {
	return &GetFilesUnauthorized{}
}

/* GetFilesUnauthorized describes a response with status code 401, with default header values.

[Unauthorized] Bearer token is missing, expired, or invalid.
*/
type GetFilesUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *GetFilesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /files/{path}][%d] getFilesUnauthorized  %+v", 401, o.Payload)
}
func (o *GetFilesUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetFilesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFilesForbidden creates a GetFilesForbidden with default headers values
func NewGetFilesForbidden() *GetFilesForbidden {
	return &GetFilesForbidden{}
}

/* GetFilesForbidden describes a response with status code 403, with default header values.

[Forbidden] User is not authorized to access storage APIs.
*/
type GetFilesForbidden struct {
	Payload *models.ErrorResponse
}

func (o *GetFilesForbidden) Error() string {
	return fmt.Sprintf("[GET /files/{path}][%d] getFilesForbidden  %+v", 403, o.Payload)
}
func (o *GetFilesForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetFilesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFilesNotFound creates a GetFilesNotFound with default headers values
func NewGetFilesNotFound() *GetFilesNotFound {
	return &GetFilesNotFound{}
}

/* GetFilesNotFound describes a response with status code 404, with default header values.

[Not Found] File was not found.
*/
type GetFilesNotFound struct {
	Payload *models.ErrorResponse
}

func (o *GetFilesNotFound) Error() string {
	return fmt.Sprintf("[GET /files/{path}][%d] getFilesNotFound  %+v", 404, o.Payload)
}
func (o *GetFilesNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetFilesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFilesGone creates a GetFilesGone with default headers values
func NewGetFilesGone() *GetFilesGone {
	return &GetFilesGone{}
}

/* GetFilesGone describes a response with status code 410, with default header values.

[Gone] Content was removed.
*/
type GetFilesGone struct {
	Payload *models.ErrorResponse
}

func (o *GetFilesGone) Error() string {
	return fmt.Sprintf("[GET /files/{path}][%d] getFilesGone  %+v", 410, o.Payload)
}
func (o *GetFilesGone) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetFilesGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFilesServiceUnavailable creates a GetFilesServiceUnavailable with default headers values
func NewGetFilesServiceUnavailable() *GetFilesServiceUnavailable {
	return &GetFilesServiceUnavailable{}
}

/* GetFilesServiceUnavailable describes a response with status code 503, with default header values.

[Service Unavailable] See response body for more detail.
*/
type GetFilesServiceUnavailable struct {
	Payload *models.ErrorResponse
}

func (o *GetFilesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /files/{path}][%d] getFilesServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetFilesServiceUnavailable) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetFilesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}