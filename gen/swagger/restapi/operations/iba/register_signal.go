// Code generated by go-swagger; DO NOT EDIT.

package iba

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

// RegisterSignalHandlerFunc turns a function with the right signature into a register signal handler
type RegisterSignalHandlerFunc func(RegisterSignalParams) middleware.Responder

// Handle executing the request and returning a response
func (fn RegisterSignalHandlerFunc) Handle(params RegisterSignalParams) middleware.Responder {
	return fn(params)
}

// RegisterSignalHandler interface for that can handle valid register signal params
type RegisterSignalHandler interface {
	Handle(RegisterSignalParams) middleware.Responder
}

// NewRegisterSignal creates a new http.Handler for the register signal operation
func NewRegisterSignal(ctx *middleware.Context, handler RegisterSignalHandler) *RegisterSignal {
	return &RegisterSignal{Context: ctx, Handler: handler}
}

/* RegisterSignal swagger:route PUT /ibas/servers/{iba_server_id}/signals iba registerSignal

Register new signal for IBA Server(batch mode)

*/
type RegisterSignal struct {
	Context *middleware.Context
	Handler RegisterSignalHandler
}

func (o *RegisterSignal) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewRegisterSignalParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// RegisterSignalBody register signal body
//
// swagger:model RegisterSignalBody
type RegisterSignalBody struct {

	// is digital
	// Required: true
	IsDigital *bool `json:"is_digital"`

	// module number
	// Required: true
	ModuleNumber *int64 `json:"module_number"`

	// number in module
	// Required: true
	NumberInModule *int64 `json:"number_in_module"`

	// type
	// Required: true
	// Enum: [bool real double int]
	Type *string `json:"type"`
}

// Validate validates this register signal body
func (o *RegisterSignalBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateIsDigital(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateModuleNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNumberInModule(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RegisterSignalBody) validateIsDigital(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"is_digital", "body", o.IsDigital); err != nil {
		return err
	}

	return nil
}

func (o *RegisterSignalBody) validateModuleNumber(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"module_number", "body", o.ModuleNumber); err != nil {
		return err
	}

	return nil
}

func (o *RegisterSignalBody) validateNumberInModule(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"number_in_module", "body", o.NumberInModule); err != nil {
		return err
	}

	return nil
}

var registerSignalBodyTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["bool","real","double","int"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		registerSignalBodyTypeTypePropEnum = append(registerSignalBodyTypeTypePropEnum, v)
	}
}

const (

	// RegisterSignalBodyTypeBool captures enum value "bool"
	RegisterSignalBodyTypeBool string = "bool"

	// RegisterSignalBodyTypeReal captures enum value "real"
	RegisterSignalBodyTypeReal string = "real"

	// RegisterSignalBodyTypeDouble captures enum value "double"
	RegisterSignalBodyTypeDouble string = "double"

	// RegisterSignalBodyTypeInt captures enum value "int"
	RegisterSignalBodyTypeInt string = "int"
)

// prop value enum
func (o *RegisterSignalBody) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, registerSignalBodyTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *RegisterSignalBody) validateType(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"type", "body", o.Type); err != nil {
		return err
	}

	// value enum
	if err := o.validateTypeEnum("body"+"."+"type", "body", *o.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this register signal body based on context it is used
func (o *RegisterSignalBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RegisterSignalBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RegisterSignalBody) UnmarshalBinary(b []byte) error {
	var res RegisterSignalBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
