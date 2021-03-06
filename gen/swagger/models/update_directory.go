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

// UpdateDirectory update directory
//
// swagger:model updateDirectory
type UpdateDirectory struct {

	// acl
	ACL interface{} `json:"acl,omitempty"`

	// description
	// Example: Description
	// Pattern: ^.{2,200}$
	Description *string `json:"description,omitempty"`

	// name
	// Example: directory1
	// Pattern: ^[A-Za-z0-9_.]{2,50}$
	Name *string `json:"name,omitempty"`
}

// Validate validates this update directory
func (m *UpdateDirectory) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateDirectory) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.Pattern("description", "body", *m.Description, `^.{2,200}$`); err != nil {
		return err
	}

	return nil
}

func (m *UpdateDirectory) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", *m.Name, `^[A-Za-z0-9_.]{2,50}$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update directory based on context it is used
func (m *UpdateDirectory) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateDirectory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateDirectory) UnmarshalBinary(b []byte) error {
	var res UpdateDirectory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
