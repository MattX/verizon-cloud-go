// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ContactsDelta contacts delta
//
// swagger:model contacts_delta
type ContactsDelta struct {

	// Array of contact ID's
	Delete []string `json:"delete"`

	// Number of contacts
	Itemcount int64 `json:"itemcount,omitempty"`

	// Array of contacts
	Modify []*Contact `json:"modify"`

	// Array of contacts
	New []*Contact `json:"new"`
}

// Validate validates this contacts delta
func (m *ContactsDelta) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateModify(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNew(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContactsDelta) validateModify(formats strfmt.Registry) error {
	if swag.IsZero(m.Modify) { // not required
		return nil
	}

	for i := 0; i < len(m.Modify); i++ {
		if swag.IsZero(m.Modify[i]) { // not required
			continue
		}

		if m.Modify[i] != nil {
			if err := m.Modify[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("modify" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("modify" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ContactsDelta) validateNew(formats strfmt.Registry) error {
	if swag.IsZero(m.New) { // not required
		return nil
	}

	for i := 0; i < len(m.New); i++ {
		if swag.IsZero(m.New[i]) { // not required
			continue
		}

		if m.New[i] != nil {
			if err := m.New[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("new" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("new" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this contacts delta based on the context it is used
func (m *ContactsDelta) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateModify(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNew(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContactsDelta) contextValidateModify(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Modify); i++ {

		if m.Modify[i] != nil {
			if err := m.Modify[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("modify" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("modify" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ContactsDelta) contextValidateNew(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.New); i++ {

		if m.New[i] != nil {
			if err := m.New[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("new" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("new" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContactsDelta) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContactsDelta) UnmarshalBinary(b []byte) error {
	var res ContactsDelta
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}