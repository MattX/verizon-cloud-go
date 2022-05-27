// Code generated by go-swagger; DO NOT EDIT.

package files_and_folders

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

// NewDeleteDeleteParams creates a new DeleteDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteDeleteParams() *DeleteDeleteParams {
	return &DeleteDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDeleteParamsWithTimeout creates a new DeleteDeleteParams object
// with the ability to set a timeout on a request.
func NewDeleteDeleteParamsWithTimeout(timeout time.Duration) *DeleteDeleteParams {
	return &DeleteDeleteParams{
		timeout: timeout,
	}
}

// NewDeleteDeleteParamsWithContext creates a new DeleteDeleteParams object
// with the ability to set a context for a request.
func NewDeleteDeleteParamsWithContext(ctx context.Context) *DeleteDeleteParams {
	return &DeleteDeleteParams{
		Context: ctx,
	}
}

// NewDeleteDeleteParamsWithHTTPClient creates a new DeleteDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteDeleteParamsWithHTTPClient(client *http.Client) *DeleteDeleteParams {
	return &DeleteDeleteParams{
		HTTPClient: client,
	}
}

/* DeleteDeleteParams contains all the parameters to send to the API endpoint
   for the delete delete operation.

   Typically these are written to a http.Request.
*/
type DeleteDeleteParams struct {

	/* Path.

	   Full path of the file/folder to be deleted.
	*/
	Path string

	/* Purge.

	   If 'true', permanently deletes the file/folder.
	*/
	Purge *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDeleteParams) WithDefaults() *DeleteDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDeleteParams) SetDefaults() {
	var (
		purgeDefault = bool(false)
	)

	val := DeleteDeleteParams{
		Purge: &purgeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete delete params
func (o *DeleteDeleteParams) WithTimeout(timeout time.Duration) *DeleteDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete delete params
func (o *DeleteDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete delete params
func (o *DeleteDeleteParams) WithContext(ctx context.Context) *DeleteDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete delete params
func (o *DeleteDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete delete params
func (o *DeleteDeleteParams) WithHTTPClient(client *http.Client) *DeleteDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete delete params
func (o *DeleteDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPath adds the path to the delete delete params
func (o *DeleteDeleteParams) WithPath(path string) *DeleteDeleteParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the delete delete params
func (o *DeleteDeleteParams) SetPath(path string) {
	o.Path = path
}

// WithPurge adds the purge to the delete delete params
func (o *DeleteDeleteParams) WithPurge(purge *bool) *DeleteDeleteParams {
	o.SetPurge(purge)
	return o
}

// SetPurge adds the purge to the delete delete params
func (o *DeleteDeleteParams) SetPurge(purge *bool) {
	o.Purge = purge
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param path
	qrPath := o.Path
	qPath := qrPath
	if qPath != "" {

		if err := r.SetQueryParam("path", qPath); err != nil {
			return err
		}
	}

	if o.Purge != nil {

		// query param purge
		var qrPurge bool

		if o.Purge != nil {
			qrPurge = *o.Purge
		}
		qPurge := swag.FormatBool(qrPurge)
		if qPurge != "" {

			if err := r.SetQueryParam("purge", qPurge); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}