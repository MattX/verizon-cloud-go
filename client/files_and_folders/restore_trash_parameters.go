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

// NewRestoreTrashParams creates a new RestoreTrashParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRestoreTrashParams() *RestoreTrashParams {
	return &RestoreTrashParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRestoreTrashParamsWithTimeout creates a new RestoreTrashParams object
// with the ability to set a timeout on a request.
func NewRestoreTrashParamsWithTimeout(timeout time.Duration) *RestoreTrashParams {
	return &RestoreTrashParams{
		timeout: timeout,
	}
}

// NewRestoreTrashParamsWithContext creates a new RestoreTrashParams object
// with the ability to set a context for a request.
func NewRestoreTrashParamsWithContext(ctx context.Context) *RestoreTrashParams {
	return &RestoreTrashParams{
		Context: ctx,
	}
}

// NewRestoreTrashParamsWithHTTPClient creates a new RestoreTrashParams object
// with the ability to set a custom HTTPClient for a request.
func NewRestoreTrashParamsWithHTTPClient(client *http.Client) *RestoreTrashParams {
	return &RestoreTrashParams{
		HTTPClient: client,
	}
}

/* RestoreTrashParams contains all the parameters to send to the API endpoint
   for the restore trash operation.

   Typically these are written to a http.Request.
*/
type RestoreTrashParams struct {

	/* TrashcanItems.

	   Request object to restore files or folders from trash.
	*/
	TrashcanItems *models.RestoreTrashRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the restore trash params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestoreTrashParams) WithDefaults() *RestoreTrashParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the restore trash params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestoreTrashParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the restore trash params
func (o *RestoreTrashParams) WithTimeout(timeout time.Duration) *RestoreTrashParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the restore trash params
func (o *RestoreTrashParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the restore trash params
func (o *RestoreTrashParams) WithContext(ctx context.Context) *RestoreTrashParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the restore trash params
func (o *RestoreTrashParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the restore trash params
func (o *RestoreTrashParams) WithHTTPClient(client *http.Client) *RestoreTrashParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the restore trash params
func (o *RestoreTrashParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTrashcanItems adds the trashcanItems to the restore trash params
func (o *RestoreTrashParams) WithTrashcanItems(trashcanItems *models.RestoreTrashRequest) *RestoreTrashParams {
	o.SetTrashcanItems(trashcanItems)
	return o
}

// SetTrashcanItems adds the trashcanItems to the restore trash params
func (o *RestoreTrashParams) SetTrashcanItems(trashcanItems *models.RestoreTrashRequest) {
	o.TrashcanItems = trashcanItems
}

// WriteToRequest writes these params to a swagger request
func (o *RestoreTrashParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.TrashcanItems != nil {
		if err := r.SetBodyParam(o.TrashcanItems); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
