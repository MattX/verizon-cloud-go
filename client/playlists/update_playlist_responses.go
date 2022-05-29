// Code generated by go-swagger; DO NOT EDIT.

package playlists

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mattx/verizon-cloud-go/models"
)

// UpdatePlaylistReader is a Reader for the UpdatePlaylist structure.
type UpdatePlaylistReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdatePlaylistReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdatePlaylistOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdatePlaylistBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdatePlaylistUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdatePlaylistForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdatePlaylistNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewUpdatePlaylistServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdatePlaylistOK creates a UpdatePlaylistOK with default headers values
func NewUpdatePlaylistOK() *UpdatePlaylistOK {
	return &UpdatePlaylistOK{}
}

/* UpdatePlaylistOK describes a response with status code 200, with default header values.

OK
*/
type UpdatePlaylistOK struct {
	Payload *models.Playlist
}

func (o *UpdatePlaylistOK) Error() string {
	return fmt.Sprintf("[PUT /playlists/{playlistUid}/items][%d] updatePlaylistOK  %+v", 200, o.Payload)
}
func (o *UpdatePlaylistOK) GetPayload() *models.Playlist {
	return o.Payload
}

func (o *UpdatePlaylistOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Playlist)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePlaylistBadRequest creates a UpdatePlaylistBadRequest with default headers values
func NewUpdatePlaylistBadRequest() *UpdatePlaylistBadRequest {
	return &UpdatePlaylistBadRequest{}
}

/* UpdatePlaylistBadRequest describes a response with status code 400, with default header values.

[Bad Request] Parameters missing or invalid.
*/
type UpdatePlaylistBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *UpdatePlaylistBadRequest) Error() string {
	return fmt.Sprintf("[PUT /playlists/{playlistUid}/items][%d] updatePlaylistBadRequest  %+v", 400, o.Payload)
}
func (o *UpdatePlaylistBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdatePlaylistBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePlaylistUnauthorized creates a UpdatePlaylistUnauthorized with default headers values
func NewUpdatePlaylistUnauthorized() *UpdatePlaylistUnauthorized {
	return &UpdatePlaylistUnauthorized{}
}

/* UpdatePlaylistUnauthorized describes a response with status code 401, with default header values.

[Unauthorized] Bearer token is missing, expired, or invalid.
*/
type UpdatePlaylistUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *UpdatePlaylistUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /playlists/{playlistUid}/items][%d] updatePlaylistUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdatePlaylistUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdatePlaylistUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePlaylistForbidden creates a UpdatePlaylistForbidden with default headers values
func NewUpdatePlaylistForbidden() *UpdatePlaylistForbidden {
	return &UpdatePlaylistForbidden{}
}

/* UpdatePlaylistForbidden describes a response with status code 403, with default header values.

[Forbidden] User is not authorized to access storage APIs.
*/
type UpdatePlaylistForbidden struct {
	Payload *models.ErrorResponse
}

func (o *UpdatePlaylistForbidden) Error() string {
	return fmt.Sprintf("[PUT /playlists/{playlistUid}/items][%d] updatePlaylistForbidden  %+v", 403, o.Payload)
}
func (o *UpdatePlaylistForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdatePlaylistForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePlaylistNotFound creates a UpdatePlaylistNotFound with default headers values
func NewUpdatePlaylistNotFound() *UpdatePlaylistNotFound {
	return &UpdatePlaylistNotFound{}
}

/* UpdatePlaylistNotFound describes a response with status code 404, with default header values.

[Not Found] Resource not found.
*/
type UpdatePlaylistNotFound struct {
	Payload *models.ErrorResponse
}

func (o *UpdatePlaylistNotFound) Error() string {
	return fmt.Sprintf("[PUT /playlists/{playlistUid}/items][%d] updatePlaylistNotFound  %+v", 404, o.Payload)
}
func (o *UpdatePlaylistNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdatePlaylistNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePlaylistServiceUnavailable creates a UpdatePlaylistServiceUnavailable with default headers values
func NewUpdatePlaylistServiceUnavailable() *UpdatePlaylistServiceUnavailable {
	return &UpdatePlaylistServiceUnavailable{}
}

/* UpdatePlaylistServiceUnavailable describes a response with status code 503, with default header values.

[Service Unavailable] See response body for more detail.
*/
type UpdatePlaylistServiceUnavailable struct {
	Payload *models.ErrorResponse
}

func (o *UpdatePlaylistServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /playlists/{playlistUid}/items][%d] updatePlaylistServiceUnavailable  %+v", 503, o.Payload)
}
func (o *UpdatePlaylistServiceUnavailable) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdatePlaylistServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}