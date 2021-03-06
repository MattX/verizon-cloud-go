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

	"github.com/mattx/verizon-cloud-go/models"
)

// NewRenameVirtualFolderParams creates a new RenameVirtualFolderParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRenameVirtualFolderParams() *RenameVirtualFolderParams {
	return &RenameVirtualFolderParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRenameVirtualFolderParamsWithTimeout creates a new RenameVirtualFolderParams object
// with the ability to set a timeout on a request.
func NewRenameVirtualFolderParamsWithTimeout(timeout time.Duration) *RenameVirtualFolderParams {
	return &RenameVirtualFolderParams{
		timeout: timeout,
	}
}

// NewRenameVirtualFolderParamsWithContext creates a new RenameVirtualFolderParams object
// with the ability to set a context for a request.
func NewRenameVirtualFolderParamsWithContext(ctx context.Context) *RenameVirtualFolderParams {
	return &RenameVirtualFolderParams{
		Context: ctx,
	}
}

// NewRenameVirtualFolderParamsWithHTTPClient creates a new RenameVirtualFolderParams object
// with the ability to set a custom HTTPClient for a request.
func NewRenameVirtualFolderParamsWithHTTPClient(client *http.Client) *RenameVirtualFolderParams {
	return &RenameVirtualFolderParams{
		HTTPClient: client,
	}
}

/* RenameVirtualFolderParams contains all the parameters to send to the API endpoint
   for the rename virtual folder operation.

   Typically these are written to a http.Request.
*/
type RenameVirtualFolderParams struct {

	/* RenameVirtualFolderRequest.

	   Allows a user to rename a virtual folder
	*/
	RenameVirtualFolderRequest *models.RenameVirtualFolderRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the rename virtual folder params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RenameVirtualFolderParams) WithDefaults() *RenameVirtualFolderParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the rename virtual folder params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RenameVirtualFolderParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the rename virtual folder params
func (o *RenameVirtualFolderParams) WithTimeout(timeout time.Duration) *RenameVirtualFolderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the rename virtual folder params
func (o *RenameVirtualFolderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the rename virtual folder params
func (o *RenameVirtualFolderParams) WithContext(ctx context.Context) *RenameVirtualFolderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the rename virtual folder params
func (o *RenameVirtualFolderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the rename virtual folder params
func (o *RenameVirtualFolderParams) WithHTTPClient(client *http.Client) *RenameVirtualFolderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the rename virtual folder params
func (o *RenameVirtualFolderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRenameVirtualFolderRequest adds the renameVirtualFolderRequest to the rename virtual folder params
func (o *RenameVirtualFolderParams) WithRenameVirtualFolderRequest(renameVirtualFolderRequest *models.RenameVirtualFolderRequest) *RenameVirtualFolderParams {
	o.SetRenameVirtualFolderRequest(renameVirtualFolderRequest)
	return o
}

// SetRenameVirtualFolderRequest adds the renameVirtualFolderRequest to the rename virtual folder params
func (o *RenameVirtualFolderParams) SetRenameVirtualFolderRequest(renameVirtualFolderRequest *models.RenameVirtualFolderRequest) {
	o.RenameVirtualFolderRequest = renameVirtualFolderRequest
}

// WriteToRequest writes these params to a swagger request
func (o *RenameVirtualFolderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.RenameVirtualFolderRequest != nil {
		if err := r.SetBodyParam(o.RenameVirtualFolderRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
