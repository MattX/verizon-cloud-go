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

	"github.com/mattx/verizon-cloud-go/client/models"
)

// NewPostRenameParams creates a new PostRenameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostRenameParams() *PostRenameParams {
	return &PostRenameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostRenameParamsWithTimeout creates a new PostRenameParams object
// with the ability to set a timeout on a request.
func NewPostRenameParamsWithTimeout(timeout time.Duration) *PostRenameParams {
	return &PostRenameParams{
		timeout: timeout,
	}
}

// NewPostRenameParamsWithContext creates a new PostRenameParams object
// with the ability to set a context for a request.
func NewPostRenameParamsWithContext(ctx context.Context) *PostRenameParams {
	return &PostRenameParams{
		Context: ctx,
	}
}

// NewPostRenameParamsWithHTTPClient creates a new PostRenameParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostRenameParamsWithHTTPClient(client *http.Client) *PostRenameParams {
	return &PostRenameParams{
		HTTPClient: client,
	}
}

/* PostRenameParams contains all the parameters to send to the API endpoint
   for the post rename operation.

   Typically these are written to a http.Request.
*/
type PostRenameParams struct {

	/* FileRenameRequest.

	   'safe' and 'conflictsolve' are optional parameters in the request object to rename a file/folder.
	*/
	FileRenameRequest *models.FopsChangeRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post rename params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostRenameParams) WithDefaults() *PostRenameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post rename params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostRenameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post rename params
func (o *PostRenameParams) WithTimeout(timeout time.Duration) *PostRenameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post rename params
func (o *PostRenameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post rename params
func (o *PostRenameParams) WithContext(ctx context.Context) *PostRenameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post rename params
func (o *PostRenameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post rename params
func (o *PostRenameParams) WithHTTPClient(client *http.Client) *PostRenameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post rename params
func (o *PostRenameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFileRenameRequest adds the fileRenameRequest to the post rename params
func (o *PostRenameParams) WithFileRenameRequest(fileRenameRequest *models.FopsChangeRequest) *PostRenameParams {
	o.SetFileRenameRequest(fileRenameRequest)
	return o
}

// SetFileRenameRequest adds the fileRenameRequest to the post rename params
func (o *PostRenameParams) SetFileRenameRequest(fileRenameRequest *models.FopsChangeRequest) {
	o.FileRenameRequest = fileRenameRequest
}

// WriteToRequest writes these params to a swagger request
func (o *PostRenameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.FileRenameRequest != nil {
		if err := r.SetBodyParam(o.FileRenameRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}