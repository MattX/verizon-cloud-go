// Code generated by go-swagger; DO NOT EDIT.

package tags

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

	"github.com/MattX/verizon-cloud-go/models"
)

// NewUpdateTagsParams creates a new UpdateTagsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateTagsParams() *UpdateTagsParams {
	return &UpdateTagsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateTagsParamsWithTimeout creates a new UpdateTagsParams object
// with the ability to set a timeout on a request.
func NewUpdateTagsParamsWithTimeout(timeout time.Duration) *UpdateTagsParams {
	return &UpdateTagsParams{
		timeout: timeout,
	}
}

// NewUpdateTagsParamsWithContext creates a new UpdateTagsParams object
// with the ability to set a context for a request.
func NewUpdateTagsParamsWithContext(ctx context.Context) *UpdateTagsParams {
	return &UpdateTagsParams{
		Context: ctx,
	}
}

// NewUpdateTagsParamsWithHTTPClient creates a new UpdateTagsParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateTagsParamsWithHTTPClient(client *http.Client) *UpdateTagsParams {
	return &UpdateTagsParams{
		HTTPClient: client,
	}
}

/* UpdateTagsParams contains all the parameters to send to the API endpoint
   for the update tags operation.

   Typically these are written to a http.Request.
*/
type UpdateTagsParams struct {

	/* UpdateTags.

	   Allows a user to update Tags on a file or folder
	*/
	UpdateTags *models.UpdateTagsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateTagsParams) WithDefaults() *UpdateTagsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateTagsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update tags params
func (o *UpdateTagsParams) WithTimeout(timeout time.Duration) *UpdateTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update tags params
func (o *UpdateTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update tags params
func (o *UpdateTagsParams) WithContext(ctx context.Context) *UpdateTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update tags params
func (o *UpdateTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update tags params
func (o *UpdateTagsParams) WithHTTPClient(client *http.Client) *UpdateTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update tags params
func (o *UpdateTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUpdateTags adds the updateTags to the update tags params
func (o *UpdateTagsParams) WithUpdateTags(updateTags *models.UpdateTagsRequest) *UpdateTagsParams {
	o.SetUpdateTags(updateTags)
	return o
}

// SetUpdateTags adds the updateTags to the update tags params
func (o *UpdateTagsParams) SetUpdateTags(updateTags *models.UpdateTagsRequest) {
	o.UpdateTags = updateTags
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.UpdateTags != nil {
		if err := r.SetBodyParam(o.UpdateTags); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
