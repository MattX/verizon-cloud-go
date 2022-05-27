// Code generated by go-swagger; DO NOT EDIT.

package playlists

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/MattX/verizon-cloud-go/client/models"
)

// GetPlaylistContentReader is a Reader for the GetPlaylistContent structure.
type GetPlaylistContentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPlaylistContentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPlaylistContentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetPlaylistContentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPlaylistContentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPlaylistContentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetPlaylistContentServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPlaylistContentOK creates a GetPlaylistContentOK with default headers values
func NewGetPlaylistContentOK() *GetPlaylistContentOK {
	return &GetPlaylistContentOK{}
}

/* GetPlaylistContentOK describes a response with status code 200, with default header values.

OK
*/
type GetPlaylistContentOK struct {
}

func (o *GetPlaylistContentOK) Error() string {
	return fmt.Sprintf("[GET /playlists/{playlistUid}/items/{itemUid}][%d] getPlaylistContentOK ", 200)
}

func (o *GetPlaylistContentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPlaylistContentUnauthorized creates a GetPlaylistContentUnauthorized with default headers values
func NewGetPlaylistContentUnauthorized() *GetPlaylistContentUnauthorized {
	return &GetPlaylistContentUnauthorized{}
}

/* GetPlaylistContentUnauthorized describes a response with status code 401, with default header values.

[Unauthorized] Bearer token is missing, expired, or invalid.
*/
type GetPlaylistContentUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *GetPlaylistContentUnauthorized) Error() string {
	return fmt.Sprintf("[GET /playlists/{playlistUid}/items/{itemUid}][%d] getPlaylistContentUnauthorized  %+v", 401, o.Payload)
}
func (o *GetPlaylistContentUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetPlaylistContentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPlaylistContentForbidden creates a GetPlaylistContentForbidden with default headers values
func NewGetPlaylistContentForbidden() *GetPlaylistContentForbidden {
	return &GetPlaylistContentForbidden{}
}

/* GetPlaylistContentForbidden describes a response with status code 403, with default header values.

[Forbidden] User is not authorized to access storage APIs.
*/
type GetPlaylistContentForbidden struct {
	Payload *models.ErrorResponse
}

func (o *GetPlaylistContentForbidden) Error() string {
	return fmt.Sprintf("[GET /playlists/{playlistUid}/items/{itemUid}][%d] getPlaylistContentForbidden  %+v", 403, o.Payload)
}
func (o *GetPlaylistContentForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetPlaylistContentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPlaylistContentNotFound creates a GetPlaylistContentNotFound with default headers values
func NewGetPlaylistContentNotFound() *GetPlaylistContentNotFound {
	return &GetPlaylistContentNotFound{}
}

/* GetPlaylistContentNotFound describes a response with status code 404, with default header values.

[Not Found] Resource not found.
*/
type GetPlaylistContentNotFound struct {
	Payload *models.ErrorResponse
}

func (o *GetPlaylistContentNotFound) Error() string {
	return fmt.Sprintf("[GET /playlists/{playlistUid}/items/{itemUid}][%d] getPlaylistContentNotFound  %+v", 404, o.Payload)
}
func (o *GetPlaylistContentNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetPlaylistContentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPlaylistContentServiceUnavailable creates a GetPlaylistContentServiceUnavailable with default headers values
func NewGetPlaylistContentServiceUnavailable() *GetPlaylistContentServiceUnavailable {
	return &GetPlaylistContentServiceUnavailable{}
}

/* GetPlaylistContentServiceUnavailable describes a response with status code 503, with default header values.

[Service Unavailable] See response body for more detail.
*/
type GetPlaylistContentServiceUnavailable struct {
	Payload *models.ErrorResponse
}

func (o *GetPlaylistContentServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /playlists/{playlistUid}/items/{itemUid}][%d] getPlaylistContentServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetPlaylistContentServiceUnavailable) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetPlaylistContentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
