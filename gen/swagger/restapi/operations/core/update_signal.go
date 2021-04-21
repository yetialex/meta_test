// Code generated by go-swagger; DO NOT EDIT.

package core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateSignalHandlerFunc turns a function with the right signature into a update signal handler
type UpdateSignalHandlerFunc func(UpdateSignalParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateSignalHandlerFunc) Handle(params UpdateSignalParams) middleware.Responder {
	return fn(params)
}

// UpdateSignalHandler interface for that can handle valid update signal params
type UpdateSignalHandler interface {
	Handle(UpdateSignalParams) middleware.Responder
}

// NewUpdateSignal creates a new http.Handler for the update signal operation
func NewUpdateSignal(ctx *middleware.Context, handler UpdateSignalHandler) *UpdateSignal {
	return &UpdateSignal{Context: ctx, Handler: handler}
}

/* UpdateSignal swagger:route PATCH /core/signals/{signal_id} core signals updateSignal

Update signal

*/
type UpdateSignal struct {
	Context *middleware.Context
	Handler UpdateSignalHandler
}

func (o *UpdateSignal) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateSignalParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// UpdateSignalBadRequestBody update signal bad request body
//
// swagger:model UpdateSignalBadRequestBody
type UpdateSignalBadRequestBody struct {

	// code
	// Example: 300
	Code int64 `json:"code,omitempty"`

	// message
	// Example: Something bad happens.
	Message string `json:"message,omitempty"`
}

// Validate validates this update signal bad request body
func (o *UpdateSignalBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update signal bad request body based on context it is used
func (o *UpdateSignalBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateSignalBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateSignalBadRequestBody) UnmarshalBinary(b []byte) error {
	var res UpdateSignalBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// UpdateSignalBody update signal body
//
// swagger:model UpdateSignalBody
type UpdateSignalBody struct {

	// acl
	ACL interface{} `json:"acl,omitempty"`

	// class
	// Example: analog
	// Enum: [analog discrete virtual]
	Class *string `json:"class,omitempty"`

	// description
	// Example: Description
	// Pattern: ^.{2,200}$
	Description *string `json:"description,omitempty"`

	// external id
	// Example: 1
	ExternalID *int64 `json:"external_id,omitempty"`

	// name
	// Example: node1
	// Pattern: ^[A-Za-z0-9_.]{2,50}$
	Name *string `json:"name,omitempty"`

	// prev id
	// Example: 1
	PrevID *int64 `json:"prev_id,omitempty"`

	// Measurement unit
	// Example: C
	Unit *string `json:"unit,omitempty"`

	// Value data type
	// Example: real
	// Enum: [undefined integer real boolean]
	ValueType *string `json:"value_type,omitempty"`
}

// Validate validates this update signal body
func (o *UpdateSignalBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateClass(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateValueType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var updateSignalBodyTypeClassPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["analog","discrete","virtual"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateSignalBodyTypeClassPropEnum = append(updateSignalBodyTypeClassPropEnum, v)
	}
}

const (

	// UpdateSignalBodyClassAnalog captures enum value "analog"
	UpdateSignalBodyClassAnalog string = "analog"

	// UpdateSignalBodyClassDiscrete captures enum value "discrete"
	UpdateSignalBodyClassDiscrete string = "discrete"

	// UpdateSignalBodyClassVirtual captures enum value "virtual"
	UpdateSignalBodyClassVirtual string = "virtual"
)

// prop value enum
func (o *UpdateSignalBody) validateClassEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateSignalBodyTypeClassPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateSignalBody) validateClass(formats strfmt.Registry) error {
	if swag.IsZero(o.Class) { // not required
		return nil
	}

	// value enum
	if err := o.validateClassEnum("body"+"."+"class", "body", *o.Class); err != nil {
		return err
	}

	return nil
}

func (o *UpdateSignalBody) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(o.Description) { // not required
		return nil
	}

	if err := validate.Pattern("body"+"."+"description", "body", *o.Description, `^.{2,200}$`); err != nil {
		return err
	}

	return nil
}

func (o *UpdateSignalBody) validateName(formats strfmt.Registry) error {
	if swag.IsZero(o.Name) { // not required
		return nil
	}

	if err := validate.Pattern("body"+"."+"name", "body", *o.Name, `^[A-Za-z0-9_.]{2,50}$`); err != nil {
		return err
	}

	return nil
}

var updateSignalBodyTypeValueTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["undefined","integer","real","boolean"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateSignalBodyTypeValueTypePropEnum = append(updateSignalBodyTypeValueTypePropEnum, v)
	}
}

const (

	// UpdateSignalBodyValueTypeUndefined captures enum value "undefined"
	UpdateSignalBodyValueTypeUndefined string = "undefined"

	// UpdateSignalBodyValueTypeInteger captures enum value "integer"
	UpdateSignalBodyValueTypeInteger string = "integer"

	// UpdateSignalBodyValueTypeReal captures enum value "real"
	UpdateSignalBodyValueTypeReal string = "real"

	// UpdateSignalBodyValueTypeBoolean captures enum value "boolean"
	UpdateSignalBodyValueTypeBoolean string = "boolean"
)

// prop value enum
func (o *UpdateSignalBody) validateValueTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateSignalBodyTypeValueTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateSignalBody) validateValueType(formats strfmt.Registry) error {
	if swag.IsZero(o.ValueType) { // not required
		return nil
	}

	// value enum
	if err := o.validateValueTypeEnum("body"+"."+"value_type", "body", *o.ValueType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update signal body based on context it is used
func (o *UpdateSignalBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateSignalBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateSignalBody) UnmarshalBinary(b []byte) error {
	var res UpdateSignalBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// UpdateSignalInternalServerErrorBody update signal internal server error body
//
// swagger:model UpdateSignalInternalServerErrorBody
type UpdateSignalInternalServerErrorBody struct {

	// code
	// Example: 300
	Code int64 `json:"code,omitempty"`

	// message
	// Example: Something bad happens.
	Message string `json:"message,omitempty"`
}

// Validate validates this update signal internal server error body
func (o *UpdateSignalInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update signal internal server error body based on context it is used
func (o *UpdateSignalInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateSignalInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateSignalInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res UpdateSignalInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// UpdateSignalNotFoundBody update signal not found body
//
// swagger:model UpdateSignalNotFoundBody
type UpdateSignalNotFoundBody struct {

	// code
	// Example: 300
	Code int64 `json:"code,omitempty"`

	// message
	// Example: Something bad happens.
	Message string `json:"message,omitempty"`
}

// Validate validates this update signal not found body
func (o *UpdateSignalNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update signal not found body based on context it is used
func (o *UpdateSignalNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateSignalNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateSignalNotFoundBody) UnmarshalBinary(b []byte) error {
	var res UpdateSignalNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// UpdateSignalOKBody update signal o k body
//
// swagger:model UpdateSignalOKBody
type UpdateSignalOKBody struct {

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
	ExternalID int64 `json:"external_id,omitempty"`

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
	// Enum: [undefined integer real boolean]
	ValueType *string `json:"value_type"`
}

// Validate validates this update signal o k body
func (o *UpdateSignalOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateClass(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateValueType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var updateSignalOKBodyTypeClassPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["analog","discrete","virtual"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateSignalOKBodyTypeClassPropEnum = append(updateSignalOKBodyTypeClassPropEnum, v)
	}
}

const (

	// UpdateSignalOKBodyClassAnalog captures enum value "analog"
	UpdateSignalOKBodyClassAnalog string = "analog"

	// UpdateSignalOKBodyClassDiscrete captures enum value "discrete"
	UpdateSignalOKBodyClassDiscrete string = "discrete"

	// UpdateSignalOKBodyClassVirtual captures enum value "virtual"
	UpdateSignalOKBodyClassVirtual string = "virtual"
)

// prop value enum
func (o *UpdateSignalOKBody) validateClassEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateSignalOKBodyTypeClassPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateSignalOKBody) validateClass(formats strfmt.Registry) error {

	if err := validate.Required("updateSignalOK"+"."+"class", "body", o.Class); err != nil {
		return err
	}

	// value enum
	if err := o.validateClassEnum("updateSignalOK"+"."+"class", "body", *o.Class); err != nil {
		return err
	}

	return nil
}

func (o *UpdateSignalOKBody) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updateSignalOK"+"."+"created_at", "body", o.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("updateSignalOK"+"."+"created_at", "body", "date-time", o.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *UpdateSignalOKBody) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(o.Description) { // not required
		return nil
	}

	if err := validate.Pattern("updateSignalOK"+"."+"description", "body", o.Description, `^.{2,200}$`); err != nil {
		return err
	}

	return nil
}

func (o *UpdateSignalOKBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("updateSignalOK"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *UpdateSignalOKBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("updateSignalOK"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *UpdateSignalOKBody) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updateSignalOK"+"."+"updated_at", "body", o.UpdatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("updateSignalOK"+"."+"updated_at", "body", "date-time", o.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

var updateSignalOKBodyTypeValueTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["undefined","integer","real","boolean"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateSignalOKBodyTypeValueTypePropEnum = append(updateSignalOKBodyTypeValueTypePropEnum, v)
	}
}

const (

	// UpdateSignalOKBodyValueTypeUndefined captures enum value "undefined"
	UpdateSignalOKBodyValueTypeUndefined string = "undefined"

	// UpdateSignalOKBodyValueTypeInteger captures enum value "integer"
	UpdateSignalOKBodyValueTypeInteger string = "integer"

	// UpdateSignalOKBodyValueTypeReal captures enum value "real"
	UpdateSignalOKBodyValueTypeReal string = "real"

	// UpdateSignalOKBodyValueTypeBoolean captures enum value "boolean"
	UpdateSignalOKBodyValueTypeBoolean string = "boolean"
)

// prop value enum
func (o *UpdateSignalOKBody) validateValueTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateSignalOKBodyTypeValueTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateSignalOKBody) validateValueType(formats strfmt.Registry) error {

	if err := validate.Required("updateSignalOK"+"."+"value_type", "body", o.ValueType); err != nil {
		return err
	}

	// value enum
	if err := o.validateValueTypeEnum("updateSignalOK"+"."+"value_type", "body", *o.ValueType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update signal o k body based on context it is used
func (o *UpdateSignalOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateSignalOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateSignalOKBody) UnmarshalBinary(b []byte) error {
	var res UpdateSignalOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// UpdateSignalUnauthorizedBody update signal unauthorized body
//
// swagger:model UpdateSignalUnauthorizedBody
type UpdateSignalUnauthorizedBody struct {

	// code
	// Example: 300
	Code int64 `json:"code,omitempty"`

	// message
	// Example: Something bad happens.
	Message string `json:"message,omitempty"`
}

// Validate validates this update signal unauthorized body
func (o *UpdateSignalUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update signal unauthorized body based on context it is used
func (o *UpdateSignalUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateSignalUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateSignalUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res UpdateSignalUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
