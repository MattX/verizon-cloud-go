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
)

// NewCreateVirtualFolderParams creates a new CreateVirtualFolderParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateVirtualFolderParams() *CreateVirtualFolderParams {
	return &CreateVirtualFolderParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateVirtualFolderParamsWithTimeout creates a new CreateVirtualFolderParams object
// with the ability to set a timeout on a request.
func NewCreateVirtualFolderParamsWithTimeout(timeout time.Duration) *CreateVirtualFolderParams {
	return &CreateVirtualFolderParams{
		timeout: timeout,
	}
}

// NewCreateVirtualFolderParamsWithContext creates a new CreateVirtualFolderParams object
// with the ability to set a context for a request.
func NewCreateVirtualFolderParamsWithContext(ctx context.Context) *CreateVirtualFolderParams {
	return &CreateVirtualFolderParams{
		Context: ctx,
	}
}

// NewCreateVirtualFolderParamsWithHTTPClient creates a new CreateVirtualFolderParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateVirtualFolderParamsWithHTTPClient(client *http.Client) *CreateVirtualFolderParams {
	return &CreateVirtualFolderParams{
		HTTPClient: client,
	}
}

/* CreateVirtualFolderParams contains all the parameters to send to the API endpoint
   for the create virtual folder operation.

   Typically these are written to a http.Request.
*/
type CreateVirtualFolderParams struct {

	/* Name.

	   Name of the virtual folder
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create virtual folder params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateVirtualFolderParams) WithDefaults() *CreateVirtualFolderParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create virtual folder params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateVirtualFolderParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create virtual folder params
func (o *CreateVirtualFolderParams) WithTimeout(timeout time.Duration) *CreateVirtualFolderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create virtual folder params
func (o *CreateVirtualFolderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create virtual folder params
func (o *CreateVirtualFolderParams) WithContext(ctx context.Context) *CreateVirtualFolderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create virtual folder params
func (o *CreateVirtualFolderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create virtual folder params
func (o *CreateVirtualFolderParams) WithHTTPClient(client *http.Client) *CreateVirtualFolderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create virtual folder params
func (o *CreateVirtualFolderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the create virtual folder params
func (o *CreateVirtualFolderParams) WithName(name string) *CreateVirtualFolderParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the create virtual folder params
func (o *CreateVirtualFolderParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *CreateVirtualFolderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}