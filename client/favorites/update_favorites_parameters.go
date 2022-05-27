// Code generated by go-swagger; DO NOT EDIT.

package favorites

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

// NewUpdateFavoritesParams creates a new UpdateFavoritesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateFavoritesParams() *UpdateFavoritesParams {
	return &UpdateFavoritesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateFavoritesParamsWithTimeout creates a new UpdateFavoritesParams object
// with the ability to set a timeout on a request.
func NewUpdateFavoritesParamsWithTimeout(timeout time.Duration) *UpdateFavoritesParams {
	return &UpdateFavoritesParams{
		timeout: timeout,
	}
}

// NewUpdateFavoritesParamsWithContext creates a new UpdateFavoritesParams object
// with the ability to set a context for a request.
func NewUpdateFavoritesParamsWithContext(ctx context.Context) *UpdateFavoritesParams {
	return &UpdateFavoritesParams{
		Context: ctx,
	}
}

// NewUpdateFavoritesParamsWithHTTPClient creates a new UpdateFavoritesParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateFavoritesParamsWithHTTPClient(client *http.Client) *UpdateFavoritesParams {
	return &UpdateFavoritesParams{
		HTTPClient: client,
	}
}

/* UpdateFavoritesParams contains all the parameters to send to the API endpoint
   for the update favorites operation.

   Typically these are written to a http.Request.
*/
type UpdateFavoritesParams struct {

	/* UpdateFavorites.

	   Allows a mark the list of file or folder as favorites
	*/
	UpdateFavorites *models.UpdateFavoriteRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update favorites params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateFavoritesParams) WithDefaults() *UpdateFavoritesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update favorites params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateFavoritesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update favorites params
func (o *UpdateFavoritesParams) WithTimeout(timeout time.Duration) *UpdateFavoritesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update favorites params
func (o *UpdateFavoritesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update favorites params
func (o *UpdateFavoritesParams) WithContext(ctx context.Context) *UpdateFavoritesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update favorites params
func (o *UpdateFavoritesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update favorites params
func (o *UpdateFavoritesParams) WithHTTPClient(client *http.Client) *UpdateFavoritesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update favorites params
func (o *UpdateFavoritesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUpdateFavorites adds the updateFavorites to the update favorites params
func (o *UpdateFavoritesParams) WithUpdateFavorites(updateFavorites *models.UpdateFavoriteRequest) *UpdateFavoritesParams {
	o.SetUpdateFavorites(updateFavorites)
	return o
}

// SetUpdateFavorites adds the updateFavorites to the update favorites params
func (o *UpdateFavoritesParams) SetUpdateFavorites(updateFavorites *models.UpdateFavoriteRequest) {
	o.UpdateFavorites = updateFavorites
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateFavoritesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.UpdateFavorites != nil {
		if err := r.SetBodyParam(o.UpdateFavorites); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
