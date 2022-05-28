// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateTagsRequest update tags request
//
// swagger:model updateTags_request
type UpdateTagsRequest struct {

	// A set of Tags
	// Required: true
	Tags *string `json:"Tags"`

	// If 'true', creates a new version of the file or folder. Defaults to 'false'.
	CreateVersion *bool `json:"createVersion,omitempty"`

	// URI of the resource. This is a URI value obtained from a fullview or metadata response
	// Required: true
	URI *string `json:"uri"`
}

// Validate validates this update tags request
func (m *UpdateTagsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateTagsRequest) validateTags(formats strfmt.Registry) error {

	if err := validate.Required("Tags", "body", m.Tags); err != nil {
		return err
	}

	return nil
}

func (m *UpdateTagsRequest) validateURI(formats strfmt.Registry) error {

	if err := validate.Required("uri", "body", m.URI); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update tags request based on context it is used
func (m *UpdateTagsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateTagsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateTagsRequest) UnmarshalBinary(b []byte) error {
	var res UpdateTagsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}