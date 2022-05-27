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

// AccountResponse account response
//
// swagger:model account_response
type AccountResponse struct {

	// Username. Available only for MDN accounts.
	Alias string `json:"alias,omitempty"`

	// The user identifier.
	// Required: true
	ID *string `json:"id"`

	// usage
	// Required: true
	Usage *AccountResponseUsage `json:"usage"`
}

// Validate validates this account response
func (m *AccountResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountResponse) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *AccountResponse) validateUsage(formats strfmt.Registry) error {

	if err := validate.Required("usage", "body", m.Usage); err != nil {
		return err
	}

	if m.Usage != nil {
		if err := m.Usage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("usage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("usage")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this account response based on the context it is used
func (m *AccountResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUsage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountResponse) contextValidateUsage(ctx context.Context, formats strfmt.Registry) error {

	if m.Usage != nil {
		if err := m.Usage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("usage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("usage")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountResponse) UnmarshalBinary(b []byte) error {
	var res AccountResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AccountResponseUsage account response usage
//
// swagger:model AccountResponseUsage
type AccountResponseUsage struct {

	// Total quota space available to the user on the cloud in bytes
	// Required: true
	Quota *int64 `json:"quota"`

	// Total quota space used in bytes
	// Required: true
	QuotaUsed *int64 `json:"quotaUsed"`
}

// Validate validates this account response usage
func (m *AccountResponseUsage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQuota(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuotaUsed(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountResponseUsage) validateQuota(formats strfmt.Registry) error {

	if err := validate.Required("usage"+"."+"quota", "body", m.Quota); err != nil {
		return err
	}

	return nil
}

func (m *AccountResponseUsage) validateQuotaUsed(formats strfmt.Registry) error {

	if err := validate.Required("usage"+"."+"quotaUsed", "body", m.QuotaUsed); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this account response usage based on context it is used
func (m *AccountResponseUsage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccountResponseUsage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountResponseUsage) UnmarshalBinary(b []byte) error {
	var res AccountResponseUsage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}