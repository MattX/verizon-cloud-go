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

// FolderRequest folder request
//
// swagger:model folder_request
type FolderRequest struct {

	// Folder name to be created.
	// Required: true
	Name *string `json:"name"`

	// Specifies what happens if a folder of the same name exists at the target path. Set to 'overwrite' to overwrite the existing folder. Set to 'modify' to treat the new folder as a modification of the old folder. If the 'override' parameter is not set, the following algorithm is applied if the existing folder has a 'deleted' attribute set to 'true', the folder is overwritten; if the old folder's 'deleted' attribute is false or not set, the folder is modified.
	// Required: true
	Override *string `json:"override"`

	// Path where the folder has to be created.
	// Required: true
	Path *string `json:"path"`
}

// Validate validates this folder request
func (m *FolderRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOverride(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePath(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FolderRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *FolderRequest) validateOverride(formats strfmt.Registry) error {

	if err := validate.Required("override", "body", m.Override); err != nil {
		return err
	}

	return nil
}

func (m *FolderRequest) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("path", "body", m.Path); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this folder request based on context it is used
func (m *FolderRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FolderRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FolderRequest) UnmarshalBinary(b []byte) error {
	var res FolderRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
