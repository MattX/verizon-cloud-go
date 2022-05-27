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

// LinkResponse link response
//
// swagger:model link_response
type LinkResponse struct {

	// The URI of the link.
	// Required: true
	Link *string `json:"link"`

	// The link identifier.
	// Required: true
	Rel *string `json:"rel"`
}

// Validate validates this link response
func (m *LinkResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLink(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LinkResponse) validateLink(formats strfmt.Registry) error {

	if err := validate.Required("link", "body", m.Link); err != nil {
		return err
	}

	return nil
}

func (m *LinkResponse) validateRel(formats strfmt.Registry) error {

	if err := validate.Required("rel", "body", m.Rel); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this link response based on context it is used
func (m *LinkResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LinkResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LinkResponse) UnmarshalBinary(b []byte) error {
	var res LinkResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}