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
)

// NewDeletePlaylistItemParams creates a new DeletePlaylistItemParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeletePlaylistItemParams() *DeletePlaylistItemParams {
	return &DeletePlaylistItemParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePlaylistItemParamsWithTimeout creates a new DeletePlaylistItemParams object
// with the ability to set a timeout on a request.
func NewDeletePlaylistItemParamsWithTimeout(timeout time.Duration) *DeletePlaylistItemParams {
	return &DeletePlaylistItemParams{
		timeout: timeout,
	}
}

// NewDeletePlaylistItemParamsWithContext creates a new DeletePlaylistItemParams object
// with the ability to set a context for a request.
func NewDeletePlaylistItemParamsWithContext(ctx context.Context) *DeletePlaylistItemParams {
	return &DeletePlaylistItemParams{
		Context: ctx,
	}
}

// NewDeletePlaylistItemParamsWithHTTPClient creates a new DeletePlaylistItemParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeletePlaylistItemParamsWithHTTPClient(client *http.Client) *DeletePlaylistItemParams {
	return &DeletePlaylistItemParams{
		HTTPClient: client,
	}
}

/* DeletePlaylistItemParams contains all the parameters to send to the API endpoint
   for the delete playlist item operation.

   Typically these are written to a http.Request.
*/
type DeletePlaylistItemParams struct {

	/* ItemUID.

	   Unique id related to a specific item in a playlist.
	*/
	ItemUID string

	/* PlaylistUID.

	   Unique id related to a specific playlist.
	*/
	PlaylistUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete playlist item params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePlaylistItemParams) WithDefaults() *DeletePlaylistItemParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete playlist item params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePlaylistItemParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete playlist item params
func (o *DeletePlaylistItemParams) WithTimeout(timeout time.Duration) *DeletePlaylistItemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete playlist item params
func (o *DeletePlaylistItemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete playlist item params
func (o *DeletePlaylistItemParams) WithContext(ctx context.Context) *DeletePlaylistItemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete playlist item params
func (o *DeletePlaylistItemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete playlist item params
func (o *DeletePlaylistItemParams) WithHTTPClient(client *http.Client) *DeletePlaylistItemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete playlist item params
func (o *DeletePlaylistItemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithItemUID adds the itemUID to the delete playlist item params
func (o *DeletePlaylistItemParams) WithItemUID(itemUID string) *DeletePlaylistItemParams {
	o.SetItemUID(itemUID)
	return o
}

// SetItemUID adds the itemUid to the delete playlist item params
func (o *DeletePlaylistItemParams) SetItemUID(itemUID string) {
	o.ItemUID = itemUID
}

// WithPlaylistUID adds the playlistUID to the delete playlist item params
func (o *DeletePlaylistItemParams) WithPlaylistUID(playlistUID string) *DeletePlaylistItemParams {
	o.SetPlaylistUID(playlistUID)
	return o
}

// SetPlaylistUID adds the playlistUid to the delete playlist item params
func (o *DeletePlaylistItemParams) SetPlaylistUID(playlistUID string) {
	o.PlaylistUID = playlistUID
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePlaylistItemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param itemUid
	if err := r.SetPathParam("itemUid", o.ItemUID); err != nil {
		return err
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
