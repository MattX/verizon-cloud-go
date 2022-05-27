// Code generated by go-swagger; DO NOT EDIT.

package shares

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

// NewListSharesParams creates a new ListSharesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListSharesParams() *ListSharesParams {
	return &ListSharesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListSharesParamsWithTimeout creates a new ListSharesParams object
// with the ability to set a timeout on a request.
func NewListSharesParamsWithTimeout(timeout time.Duration) *ListSharesParams {
	return &ListSharesParams{
		timeout: timeout,
	}
}

// NewListSharesParamsWithContext creates a new ListSharesParams object
// with the ability to set a context for a request.
func NewListSharesParamsWithContext(ctx context.Context) *ListSharesParams {
	return &ListSharesParams{
		Context: ctx,
	}
}

// NewListSharesParamsWithHTTPClient creates a new ListSharesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListSharesParamsWithHTTPClient(client *http.Client) *ListSharesParams {
	return &ListSharesParams{
		HTTPClient: client,
	}
}

/* ListSharesParams contains all the parameters to send to the API endpoint
   for the list shares operation.

   Typically these are written to a http.Request.
*/
type ListSharesParams struct {

	/* Count.

	   Maximum items to include in a paginated response, for share requests.Value must be between 1 and 200. Default is 20
	*/
	Count *int64

	/* Cursor.

	   A cursor used in paginating the response. Cursors are returned in 'next' and 'prev' links in the response body. When cursor parameter is present other parameter values are ignored
	*/
	Cursor *string

	/* Filter.

	   Filters the returned shares. Currently the value can only be 'outbound' (shared by the user).

	   Default: "outbound"
	*/
	Filter *string

	/* Since.

	   The date and time, expressed in the W3C date and time format, after which messages should be returned.
	*/
	Since *string

	/* Until.

	   The date and time, expressed in the W3C date and time format, up to which messages should be returned.
	*/
	Until *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list shares params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListSharesParams) WithDefaults() *ListSharesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list shares params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListSharesParams) SetDefaults() {
	var (
		filterDefault = string("outbound")
	)

	val := ListSharesParams{
		Filter: &filterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the list shares params
func (o *ListSharesParams) WithTimeout(timeout time.Duration) *ListSharesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list shares params
func (o *ListSharesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list shares params
func (o *ListSharesParams) WithContext(ctx context.Context) *ListSharesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list shares params
func (o *ListSharesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list shares params
func (o *ListSharesParams) WithHTTPClient(client *http.Client) *ListSharesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list shares params
func (o *ListSharesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the list shares params
func (o *ListSharesParams) WithCount(count *int64) *ListSharesParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the list shares params
func (o *ListSharesParams) SetCount(count *int64) {
	o.Count = count
}

// WithCursor adds the cursor to the list shares params
func (o *ListSharesParams) WithCursor(cursor *string) *ListSharesParams {
	o.SetCursor(cursor)
	return o
}

// SetCursor adds the cursor to the list shares params
func (o *ListSharesParams) SetCursor(cursor *string) {
	o.Cursor = cursor
}

// WithFilter adds the filter to the list shares params
func (o *ListSharesParams) WithFilter(filter *string) *ListSharesParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the list shares params
func (o *ListSharesParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithSince adds the since to the list shares params
func (o *ListSharesParams) WithSince(since *string) *ListSharesParams {
	o.SetSince(since)
	return o
}

// SetSince adds the since to the list shares params
func (o *ListSharesParams) SetSince(since *string) {
	o.Since = since
}

// WithUntil adds the until to the list shares params
func (o *ListSharesParams) WithUntil(until *string) *ListSharesParams {
	o.SetUntil(until)
	return o
}

// SetUntil adds the until to the list shares params
func (o *ListSharesParams) SetUntil(until *string) {
	o.Until = until
}

// WriteToRequest writes these params to a swagger request
func (o *ListSharesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Cursor != nil {

		// query param cursor
		var qrCursor string

		if o.Cursor != nil {
			qrCursor = *o.Cursor
		}
		qCursor := qrCursor
		if qCursor != "" {

			if err := r.SetQueryParam("cursor", qCursor); err != nil {
				return err
			}
		}
	}

	if o.Filter != nil {

		// query param filter
		var qrFilter string

		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {

			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}
	}

	if o.Since != nil {

		// query param since
		var qrSince string

		if o.Since != nil {
			qrSince = *o.Since
		}
		qSince := qrSince
		if qSince != "" {

			if err := r.SetQueryParam("since", qSince); err != nil {
				return err
			}
		}
	}

	if o.Until != nil {

		// query param until
		var qrUntil string

		if o.Until != nil {
			qrUntil = *o.Until
		}
		qUntil := qrUntil
		if qUntil != "" {

			if err := r.SetQueryParam("until", qUntil); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
