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
	"github.com/go-openapi/swag"
)

// NewGetPlaylistsParams creates a new GetPlaylistsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPlaylistsParams() *GetPlaylistsParams {
	return &GetPlaylistsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPlaylistsParamsWithTimeout creates a new GetPlaylistsParams object
// with the ability to set a timeout on a request.
func NewGetPlaylistsParamsWithTimeout(timeout time.Duration) *GetPlaylistsParams {
	return &GetPlaylistsParams{
		timeout: timeout,
	}
}

// NewGetPlaylistsParamsWithContext creates a new GetPlaylistsParams object
// with the ability to set a context for a request.
func NewGetPlaylistsParamsWithContext(ctx context.Context) *GetPlaylistsParams {
	return &GetPlaylistsParams{
		Context: ctx,
	}
}

// NewGetPlaylistsParamsWithHTTPClient creates a new GetPlaylistsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPlaylistsParamsWithHTTPClient(client *http.Client) *GetPlaylistsParams {
	return &GetPlaylistsParams{
		HTTPClient: client,
	}
}

/* GetPlaylistsParams contains all the parameters to send to the API endpoint
   for the get playlists operation.

   Typically these are written to a http.Request.
*/
type GetPlaylistsParams struct {

	/* Count.

	   Maximum children to include in a paginated response.  Defaulted to 20 if page is specified.
	*/
	Count *int64

	/* Page.

	   Page number to return, for paginated responses. Defaulted to 1 if count is specified.
	*/
	Page *int64

	/* Sort.

	   Specify sort order for response. Syntax is '{field}+{asc|desc}'. Valid values for 'field' are 'name' and 'creationDate'.
	*/
	Sort *string

	/* Type.

	   The type of the playlist. Can be one of 'image', 'music', 'video' or 'image-video'.
	*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get playlists params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPlaylistsParams) WithDefaults() *GetPlaylistsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get playlists params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPlaylistsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get playlists params
func (o *GetPlaylistsParams) WithTimeout(timeout time.Duration) *GetPlaylistsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get playlists params
func (o *GetPlaylistsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get playlists params
func (o *GetPlaylistsParams) WithContext(ctx context.Context) *GetPlaylistsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get playlists params
func (o *GetPlaylistsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get playlists params
func (o *GetPlaylistsParams) WithHTTPClient(client *http.Client) *GetPlaylistsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get playlists params
func (o *GetPlaylistsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the get playlists params
func (o *GetPlaylistsParams) WithCount(count *int64) *GetPlaylistsParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get playlists params
func (o *GetPlaylistsParams) SetCount(count *int64) {
	o.Count = count
}

// WithPage adds the page to the get playlists params
func (o *GetPlaylistsParams) WithPage(page *int64) *GetPlaylistsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get playlists params
func (o *GetPlaylistsParams) SetPage(page *int64) {
	o.Page = page
}

// WithSort adds the sort to the get playlists params
func (o *GetPlaylistsParams) WithSort(sort *string) *GetPlaylistsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get playlists params
func (o *GetPlaylistsParams) SetSort(sort *string) {
	o.Sort = sort
}

// WithType adds the typeVar to the get playlists params
func (o *GetPlaylistsParams) WithType(typeVar *string) *GetPlaylistsParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get playlists params
func (o *GetPlaylistsParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetPlaylistsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Count != nil {

		// query param count
		var qrCount int64

		if o.Count != nil {
			qrCount = *o.Count
		}
		qCount := swag.FormatInt64(qrCount)
		if qCount != "" {

			if err := r.SetQueryParam("count", qCount); err != nil {
				return err
			}
		}
	}

	if o.Page != nil {

		// query param page
		var qrPage int64

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.Sort != nil {

		// query param sort
		var qrSort string

		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {

			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}
	}

	if o.Type != nil {

		// query param type
		var qrType string

		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {

			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
