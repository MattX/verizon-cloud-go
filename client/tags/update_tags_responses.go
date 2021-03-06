// Code generated by go-swagger; DO NOT EDIT.

package tags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mattx/verizon-cloud-go/models"
)

// UpdateTagsReader is a Reader for the UpdateTags structure.
type UpdateTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateTagsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateTagsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateTagsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateTagsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewUpdateTagsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateTagsOK creates a UpdateTagsOK with default headers values
func NewUpdateTagsOK() *UpdateTagsOK {
	return &UpdateTagsOK{}
}

/* UpdateTagsOK describes a response with status code 200, with default header values.

A tags response object
*/
type UpdateTagsOK struct {
	Payload *models.TagsResponse
}

func (o *UpdateTagsOK) Error() string {
	return fmt.Sprintf("[PUT /tags][%d] updateTagsOK  %+v", 200, o.Payload)
}
func (o *UpdateTagsOK) GetPayload() *models.TagsResponse {
	return o.Payload
}

func (o *UpdateTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TagsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTagsBadRequest creates a UpdateTagsBadRequest with default headers values
func NewUpdateTagsBadRequest() *UpdateTagsBadRequest {
	return &UpdateTagsBadRequest{}
}

/* UpdateTagsBadRequest describes a response with status code 400, with default header values.

[Bad Request] Query parameters missing or invalid.
*/
type UpdateTagsBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *UpdateTagsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /tags][%d] updateTagsBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateTagsBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateTagsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTagsUnauthorized creates a UpdateTagsUnauthorized with default headers values
func NewUpdateTagsUnauthorized() *UpdateTagsUnauthorized {
	return &UpdateTagsUnauthorized{}
}

/* UpdateTagsUnauthorized describes a response with status code 401, with default header values.

[Unauthorized] Bearer token is missing, expired, or invalid.
*/
type UpdateTagsUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *UpdateTagsUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /tags][%d] updateTagsUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateTagsUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateTagsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTagsForbidden creates a UpdateTagsForbidden with default headers values
func NewUpdateTagsForbidden() *UpdateTagsForbidden {
	return &UpdateTagsForbidden{}
}

/* UpdateTagsForbidden describes a response with status code 403, with default header values.

[Forbidden] User is not authorized to access storage APIs.
*/
type UpdateTagsForbidden struct {
	Payload *models.ErrorResponse
}

func (o *UpdateTagsForbidden) Error() string {
	return fmt.Sprintf("[PUT /tags][%d] updateTagsForbidden  %+v", 403, o.Payload)
}
func (o *UpdateTagsForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateTagsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTagsNotFound creates a UpdateTagsNotFound with default headers values
func NewUpdateTagsNotFound() *UpdateTagsNotFound {
	return &UpdateTagsNotFound{}
}

/* UpdateTagsNotFound describes a response with status code 404, with default header values.

[Not Found] File or folder was not found.
*/
type UpdateTagsNotFound struct {
	Payload *models.ErrorResponse
}

func (o *UpdateTagsNotFound) Error() string {
	return fmt.Sprintf("[PUT /tags][%d] updateTagsNotFound  %+v", 404, o.Payload)
}
func (o *UpdateTagsNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateTagsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTagsServiceUnavailable creates a UpdateTagsServiceUnavailable with default headers values
func NewUpdateTagsServiceUnavailable() *UpdateTagsServiceUnavailable {
	return &UpdateTagsServiceUnavailable{}
}

/* UpdateTagsServiceUnavailable describes a response with status code 503, with default header values.

[Service Unavailable] See response body for more detail.
*/
type UpdateTagsServiceUnavailable struct {
	Payload *models.ErrorResponse
}

func (o *UpdateTagsServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /tags][%d] updateTagsServiceUnavailable  %+v", 503, o.Payload)
}
func (o *UpdateTagsServiceUnavailable) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateTagsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
