// Code generated by go-swagger; DO NOT EDIT.

package flashbacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/MattX/verizon-cloud-go/client/models"
)

// GetFlashbacksReader is a Reader for the GetFlashbacks structure.
type GetFlashbacksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFlashbacksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFlashbacksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetFlashbacksNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetFlashbacksBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetFlashbacksUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFlashbacksForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFlashbacksNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetFlashbacksServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFlashbacksOK creates a GetFlashbacksOK with default headers values
func NewGetFlashbacksOK() *GetFlashbacksOK {
	return &GetFlashbacksOK{}
}

/* GetFlashbacksOK describes a response with status code 200, with default header values.

A flashback response object
*/
type GetFlashbacksOK struct {
	Payload *models.FlashbackResponse
}

func (o *GetFlashbacksOK) Error() string {
	return fmt.Sprintf("[GET /flashbacks][%d] getFlashbacksOK  %+v", 200, o.Payload)
}
func (o *GetFlashbacksOK) GetPayload() *models.FlashbackResponse {
	return o.Payload
}

func (o *GetFlashbacksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FlashbackResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlashbacksNoContent creates a GetFlashbacksNoContent with default headers values
func NewGetFlashbacksNoContent() *GetFlashbacksNoContent {
	return &GetFlashbacksNoContent{}
}

/* GetFlashbacksNoContent describes a response with status code 204, with default header values.

OK. [The virtual folder is having no image files with image score.]
*/
type GetFlashbacksNoContent struct {
}

func (o *GetFlashbacksNoContent) Error() string {
	return fmt.Sprintf("[GET /flashbacks][%d] getFlashbacksNoContent ", 204)
}

func (o *GetFlashbacksNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFlashbacksBadRequest creates a GetFlashbacksBadRequest with default headers values
func NewGetFlashbacksBadRequest() *GetFlashbacksBadRequest {
	return &GetFlashbacksBadRequest{}
}

/* GetFlashbacksBadRequest describes a response with status code 400, with default header values.

[Bad Request] Query parameters missing or invalid.
*/
type GetFlashbacksBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *GetFlashbacksBadRequest) Error() string {
	return fmt.Sprintf("[GET /flashbacks][%d] getFlashbacksBadRequest  %+v", 400, o.Payload)
}
func (o *GetFlashbacksBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetFlashbacksBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlashbacksUnauthorized creates a GetFlashbacksUnauthorized with default headers values
func NewGetFlashbacksUnauthorized() *GetFlashbacksUnauthorized {
	return &GetFlashbacksUnauthorized{}
}

/* GetFlashbacksUnauthorized describes a response with status code 401, with default header values.

[Unauthorized] Bearer token is missing, expired, or invalid.
*/
type GetFlashbacksUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *GetFlashbacksUnauthorized) Error() string {
	return fmt.Sprintf("[GET /flashbacks][%d] getFlashbacksUnauthorized  %+v", 401, o.Payload)
}
func (o *GetFlashbacksUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetFlashbacksUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlashbacksForbidden creates a GetFlashbacksForbidden with default headers values
func NewGetFlashbacksForbidden() *GetFlashbacksForbidden {
	return &GetFlashbacksForbidden{}
}

/* GetFlashbacksForbidden describes a response with status code 403, with default header values.

[Forbidden] User is not authorized to access storage APIs.
*/
type GetFlashbacksForbidden struct {
	Payload *models.ErrorResponse
}

func (o *GetFlashbacksForbidden) Error() string {
	return fmt.Sprintf("[GET /flashbacks][%d] getFlashbacksForbidden  %+v", 403, o.Payload)
}
func (o *GetFlashbacksForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetFlashbacksForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlashbacksNotFound creates a GetFlashbacksNotFound with default headers values
func NewGetFlashbacksNotFound() *GetFlashbacksNotFound {
	return &GetFlashbacksNotFound{}
}

/* GetFlashbacksNotFound describes a response with status code 404, with default header values.

[Not Found] File was not found.
*/
type GetFlashbacksNotFound struct {
	Payload *models.ErrorResponse
}

func (o *GetFlashbacksNotFound) Error() string {
	return fmt.Sprintf("[GET /flashbacks][%d] getFlashbacksNotFound  %+v", 404, o.Payload)
}
func (o *GetFlashbacksNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetFlashbacksNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlashbacksServiceUnavailable creates a GetFlashbacksServiceUnavailable with default headers values
func NewGetFlashbacksServiceUnavailable() *GetFlashbacksServiceUnavailable {
	return &GetFlashbacksServiceUnavailable{}
}

/* GetFlashbacksServiceUnavailable describes a response with status code 503, with default header values.

[Service Unavailable] See response body for more detail.
*/
type GetFlashbacksServiceUnavailable struct {
	Payload *models.ErrorResponse
}

func (o *GetFlashbacksServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /flashbacks][%d] getFlashbacksServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetFlashbacksServiceUnavailable) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetFlashbacksServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
