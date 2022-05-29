// Code generated by go-swagger; DO NOT EDIT.

package playlists

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/mattx/verizon-cloud-go/models"
)

// NewPostPlaylistItemParams creates a new PostPlaylistItemParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostPlaylistItemParams() *PostPlaylistItemParams {
	return &PostPlaylistItemParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostPlaylistItemParamsWithTimeout creates a new PostPlaylistItemParams object
// with the ability to set a timeout on a request.
func NewPostPlaylistItemParamsWithTimeout(timeout time.Duration) *PostPlaylistItemParams {
	return &PostPlaylistItemParams{
		timeout: timeout,
	}
}

// NewPostPlaylistItemParamsWithContext creates a new PostPlaylistItemParams object
// with the ability to set a context for a request.
func NewPostPlaylistItemParamsWithContext(ctx context.Context) *PostPlaylistItemParams {
	return &PostPlaylistItemParams{
		Context: ctx,
	}
}

// NewPostPlaylistItemParamsWithHTTPClient creates a new PostPlaylistItemParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostPlaylistItemParamsWithHTTPClient(client *http.Client) *PostPlaylistItemParams {
	return &PostPlaylistItemParams{
		HTTPClient: client,
	}
}

/* PostPlaylistItemParams contains all the parameters to send to the API endpoint
   for the post playlist item operation.

   Typically these are written to a http.Request.
*/
type PostPlaylistItemParams struct {

	/* PlaylistItems.

	   Request object to add items in playlist.
	*/
	PlaylistItems *models.PlaylistAddRequest

	/* PlaylistUID.

	   Unique id related to a specific playlist.
	*/
	PlaylistUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post playlist item params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostPlaylistItemParams) WithDefaults() *PostPlaylistItemParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post playlist item params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostPlaylistItemParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post playlist item params
func (o *PostPlaylistItemParams) WithTimeout(timeout time.Duration) *PostPlaylistItemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post playlist item params
func (o *PostPlaylistItemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post playlist item params
func (o *PostPlaylistItemParams) WithContext(ctx context.Context) *PostPlaylistItemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post playlist item params
func (o *PostPlaylistItemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post playlist item params
func (o *PostPlaylistItemParams) WithHTTPClient(client *http.Client) *PostPlaylistItemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post playlist item params
func (o *PostPlaylistItemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPlaylistItems adds the playlistItems to the post playlist item params
func (o *PostPlaylistItemParams) WithPlaylistItems(playlistItems *models.PlaylistAddRequest) *PostPlaylistItemParams {
	o.SetPlaylistItems(playlistItems)
	return o
}

// SetPlaylistItems adds the playlistItems to the post playlist item params
func (o *PostPlaylistItemParams) SetPlaylistItems(playlistItems *models.PlaylistAddRequest) {
	o.PlaylistItems = playlistItems
}

// WithPlaylistUID adds the playlistUID to the post playlist item params
func (o *PostPlaylistItemParams) WithPlaylistUID(playlistUID string) *PostPlaylistItemParams {
	o.SetPlaylistUID(playlistUID)
	return o
}

// SetPlaylistUID adds the playlistUid to the post playlist item params
func (o *PostPlaylistItemParams) SetPlaylistUID(playlistUID string) {
	o.PlaylistUID = playlistUID
}

// WriteToRequest writes these params to a swagger request
func (o *PostPlaylistItemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.PlaylistItems != nil {
		if err := r.SetBodyParam(o.PlaylistItems); err != nil {
			return err
		}
	}

	// path param playlistUid
	if err := r.SetPathParam("playlistUid", o.PlaylistUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}