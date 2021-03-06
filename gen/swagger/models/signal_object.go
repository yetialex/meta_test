// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SignalObject signal object
//
// swagger:model signalObject
type SignalObject struct {

	// class
	// Example: analog
	// Required: true
	// Enum: [analog discrete virtual]
	Class *string `json:"class"`

	// ISO 8601 format
	// Example: 2018-03-20T09:12:28.123Z
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"created_at"`

	// description
	// Example: Description
	// Pattern: ^.{2,200}$
	Description string `json:"description,omitempty"`

	// external id
	// Example: 1
	ExternalID string `json:"external_id,omitempty"`

	// has children
	// Example: true
	HasChildren *bool `json:"has_children,omitempty"`

	// id
	// Example: 1
	// Required: true
	ID *int64 `json:"id"`

	// name
	// Example: some_signal
	// Required: true
	Name *string `json:"name"`

	// Measurement unit
	// Example: C
	Unit *string `json:"unit,omitempty"`

	// ISO 8601 format
	// Example: 2018-03-20T09:12:28.123Z
	// Required: true
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"updated_at"`

	// Value data type
	// Example: real
	// Required: true
	// Enum: [undefined integer real double boolean]
	ValueType *string `json:"value_type"`
}

// Validate validates this signal object
func (m *SignalObject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClass(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValueType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var signalObjectTypeClassPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["analog","discrete","virtual"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		signalObjectTypeClassPropEnum = append(signalObjectTypeClassPropEnum, v)
	}
}

const (

	// SignalObjectClassAnalog captures enum value "analog"
	SignalObjectClassAnalog string = "analog"

	// SignalObjectClassDiscrete captures enum value "discrete"
	SignalObjectClassDiscrete string = "discrete"

	// SignalObjectClassVirtual captures enum value "virtual"
	SignalObjectClassVirtual string = "virtual"
)

// prop value enum
func (m *SignalObject) validateClassEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, signalObjectTypeClassPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SignalObject) validateClass(formats strfmt.Registry) error {

	if err := validate.Required("class", "body", m.Class); err != nil {
		return err
	}

	// value enum
	if err := m.validateClassEnum("class", "body", *m.Class); err != nil {
		return err
	}

	return nil
}

func (m *SignalObject) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SignalObject) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.Pattern("description", "body", m.Description, `^.{2,200}$`); err != nil {
		return err
	}

	return nil
}

func (m *SignalObject) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *SignalObject) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *SignalObject) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

var signalObjectTypeValueTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["undefined","integer","real","double","boolean"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		signalObjectTypeValueTypePropEnum = append(signalObjectTypeValueTypePropEnum, v)
	}
}

const (

	// SignalObjectValueTypeUndefined captures enum value "undefined"
	SignalObjectValueTypeUndefined string = "undefined"

	// SignalObjectValueTypeInteger captures enum value "integer"
	SignalObjectValueTypeInteger string = "integer"

	// SignalObjectValueTypeReal captures enum value "real"
	SignalObjectValueTypeReal string = "real"

	// SignalObjectValueTypeDouble captures enum value "double"
	SignalObjectValueTypeDouble string = "double"

	// SignalObjectValueTypeBoolean captures enum value "boolean"
	SignalObjectValueTypeBoolean string = "boolean"
)

// prop value enum
func (m *SignalObject) validateValueTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, signalObjectTypeValueTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SignalObject) validateValueType(formats strfmt.Registry) error {

	if err := validate.Required("value_type", "body", m.ValueType); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueTypeEnum("value_type", "body", *m.ValueType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this signal object based on context it is used
func (m *SignalObject) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SignalObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SignalObject) UnmarshalBinary(b []byte) error {
	var res SignalObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
