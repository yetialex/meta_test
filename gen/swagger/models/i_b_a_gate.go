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

// IBAGate i b a gate
//
// swagger:model iBAGate
type IBAGate struct {

	// Комментарий
	// Example: Some notes
	Comment *string `json:"comment,omitempty"`

	// IBA Gate ID
	// Example: 1
	ID *int64 `json:"id,omitempty"`

	// Имя IBA Gate
	// Example: p3apr3-pda-vc.p3.ia.nlmk
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this i b a gate
func (m *IBAGate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IBAGate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this i b a gate based on context it is used
func (m *IBAGate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IBAGate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IBAGate) UnmarshalBinary(b []byte) error {
	var res IBAGate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
