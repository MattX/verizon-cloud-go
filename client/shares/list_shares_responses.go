// Code generated by go-swagger; DO NOT EDIT.

package shares

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mattx/verizon-cloud-go/models"
)

// ListSharesReader is a Reader for the ListShares structure.
type ListSharesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSharesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListSharesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListSharesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListSharesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListSharesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewListSharesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListSharesOK creates a ListSharesOK with default headers values
func NewListSharesOK() *ListSharesOK {
	return &ListSharesOK{}
}

/* ListSharesOK describes a response with status code 200, with default header values.

OK
*/
type ListSharesOK struct {
	Payload *models.ShareResponseList
}

func (o *ListSharesOK) Error() string {
	return fmt.Sprintf("[GET /shares][%d] listSharesOK  %+v", 200, o.Payload)
}
func (o *ListSharesOK) GetPayload() *models.ShareResponseList {
	return o.Payload
}

func (o *ListSharesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ShareResponseList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSharesBadRequest creates a ListSharesBadRequest with default headers values
func NewListSharesBadRequest() *ListSharesBadRequest {
	return &ListSharesBadRequest{}
}

/* ListSharesBadRequest describes a response with status code 400, with default header values.

[Bad Request] Parameters missing or any invalid parameter passed.
*/
type ListSharesBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *ListSharesBadRequest) Error() string {
	return fmt.Sprintf("[GET /shares][%d] listSharesBadRequest  %+v", 400, o.Payload)
}
func (o *ListSharesBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListSharesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSharesUnauthorized creates a ListSharesUnauthorized with default headers values
func NewListSharesUnauthorized() *ListSharesUnauthorized {
	return &ListSharesUnauthorized{}
}

/* ListSharesUnauthorized describes a response with status code 401, with default header values.

[Unauthorized] Bearer token is missing, expired, or invalid.
*/
type ListSharesUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *ListSharesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /shares][%d] listSharesUnauthorized  %+v", 401, o.Payload)
}
func (o *ListSharesUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListSharesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSharesForbidden creates a ListSharesForbidden with default headers values
func NewListSharesForbidden() *ListSharesForbidden {
	return &ListSharesForbidden{}
}

/* ListSharesForbidden describes a response with status code 403, with default header values.

[Forbidden] User is not authorized to access storage APIs.
*/
type ListSharesForbidden struct {
	Payload *models.ErrorResponse
}

func (o *ListSharesForbidden) Error() string {
	return fmt.Sprintf("[GET /shares][%d] listSharesForbidden  %+v", 403, o.Payload)
}
func (o *ListSharesForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListSharesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSharesServiceUnavailable creates a ListSharesServiceUnavailable with default headers values
func NewListSharesServiceUnavailable() *ListSharesServiceUnavailable {
	return &ListSharesServiceUnavailable{}
}

/* ListSharesServiceUnavailable describes a response with status code 503, with default header values.

[Service Unavailable] See response body for more detail.
*/
type ListSharesServiceUnavailable struct {
	Payload *models.ErrorResponse
}

func (o *ListSharesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /shares][%d] listSharesServiceUnavailable  %+v", 503, o.Payload)
}
func (o *ListSharesServiceUnavailable) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListSharesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
