// Code generated by go-swagger; DO NOT EDIT.

package playlists

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mattx/verizon-cloud-go/client/models"
)

// PostPlaylistItemReader is a Reader for the PostPlaylistItem structure.
type PostPlaylistItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPlaylistItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostPlaylistItemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostPlaylistItemBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostPlaylistItemUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostPlaylistItemForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostPlaylistItemNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostPlaylistItemServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostPlaylistItemOK creates a PostPlaylistItemOK with default headers values
func NewPostPlaylistItemOK() *PostPlaylistItemOK {
	return &PostPlaylistItemOK{}
}

/* PostPlaylistItemOK describes a response with status code 200, with default header values.

OK
*/
type PostPlaylistItemOK struct {
	Payload *models.PlaylistAddResponse
}

func (o *PostPlaylistItemOK) Error() string {
	return fmt.Sprintf("[POST /playlists/{playlistUid}/items][%d] postPlaylistItemOK  %+v", 200, o.Payload)
}
func (o *PostPlaylistItemOK) GetPayload() *models.PlaylistAddResponse {
	return o.Payload
}

func (o *PostPlaylistItemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PlaylistAddResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPlaylistItemBadRequest creates a PostPlaylistItemBadRequest with default headers values
func NewPostPlaylistItemBadRequest() *PostPlaylistItemBadRequest {
	return &PostPlaylistItemBadRequest{}
}

/* PostPlaylistItemBadRequest describes a response with status code 400, with default header values.

[Bad Request] Parameters missing or invalid.
*/
type PostPlaylistItemBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *PostPlaylistItemBadRequest) Error() string {
	return fmt.Sprintf("[POST /playlists/{playlistUid}/items][%d] postPlaylistItemBadRequest  %+v", 400, o.Payload)
}
func (o *PostPlaylistItemBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PostPlaylistItemBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPlaylistItemUnauthorized creates a PostPlaylistItemUnauthorized with default headers values
func NewPostPlaylistItemUnauthorized() *PostPlaylistItemUnauthorized {
	return &PostPlaylistItemUnauthorized{}
}

/* PostPlaylistItemUnauthorized describes a response with status code 401, with default header values.

[Unauthorized] Bearer token is missing, expired, or invalid.
*/
type PostPlaylistItemUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *PostPlaylistItemUnauthorized) Error() string {
	return fmt.Sprintf("[POST /playlists/{playlistUid}/items][%d] postPlaylistItemUnauthorized  %+v", 401, o.Payload)
}
func (o *PostPlaylistItemUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PostPlaylistItemUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPlaylistItemForbidden creates a PostPlaylistItemForbidden with default headers values
func NewPostPlaylistItemForbidden() *PostPlaylistItemForbidden {
	return &PostPlaylistItemForbidden{}
}

/* PostPlaylistItemForbidden describes a response with status code 403, with default header values.

[Forbidden] User is not authorized to access storage APIs.
*/
type PostPlaylistItemForbidden struct {
	Payload *models.ErrorResponse
}

func (o *PostPlaylistItemForbidden) Error() string {
	return fmt.Sprintf("[POST /playlists/{playlistUid}/items][%d] postPlaylistItemForbidden  %+v", 403, o.Payload)
}
func (o *PostPlaylistItemForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PostPlaylistItemForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPlaylistItemNotFound creates a PostPlaylistItemNotFound with default headers values
func NewPostPlaylistItemNotFound() *PostPlaylistItemNotFound {
	return &PostPlaylistItemNotFound{}
}

/* PostPlaylistItemNotFound describes a response with status code 404, with default header values.

[Not Found] Resource not found.
*/
type PostPlaylistItemNotFound struct {
	Payload *models.ErrorResponse
}

func (o *PostPlaylistItemNotFound) Error() string {
	return fmt.Sprintf("[POST /playlists/{playlistUid}/items][%d] postPlaylistItemNotFound  %+v", 404, o.Payload)
}
func (o *PostPlaylistItemNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PostPlaylistItemNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPlaylistItemServiceUnavailable creates a PostPlaylistItemServiceUnavailable with default headers values
func NewPostPlaylistItemServiceUnavailable() *PostPlaylistItemServiceUnavailable {
	return &PostPlaylistItemServiceUnavailable{}
}

/* PostPlaylistItemServiceUnavailable describes a response with status code 503, with default header values.

[Service Unavailable] See response body for more detail.
*/
type PostPlaylistItemServiceUnavailable struct {
	Payload *models.ErrorResponse
}

func (o *PostPlaylistItemServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /playlists/{playlistUid}/items][%d] postPlaylistItemServiceUnavailable  %+v", 503, o.Payload)
}
func (o *PostPlaylistItemServiceUnavailable) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PostPlaylistItemServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}