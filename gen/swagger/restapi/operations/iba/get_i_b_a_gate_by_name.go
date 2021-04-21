// Code generated by go-swagger; DO NOT EDIT.

package iba

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetIBAGateByNameHandlerFunc turns a function with the right signature into a get i b a gate by name handler
type GetIBAGateByNameHandlerFunc func(GetIBAGateByNameParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetIBAGateByNameHandlerFunc) Handle(params GetIBAGateByNameParams) middleware.Responder {
	return fn(params)
}

// GetIBAGateByNameHandler interface for that can handle valid get i b a gate by name params
type GetIBAGateByNameHandler interface {
	Handle(GetIBAGateByNameParams) middleware.Responder
}

// NewGetIBAGateByName creates a new http.Handler for the get i b a gate by name operation
func NewGetIBAGateByName(ctx *middleware.Context, handler GetIBAGateByNameHandler) *GetIBAGateByName {
	return &GetIBAGateByName{Context: ctx, Handler: handler}
}

/* GetIBAGateByName swagger:route GET /ibas/gates/${gate_name} iba getIBAGateByName

Get Gate by name

*/
type GetIBAGateByName struct {
	Context *middleware.Context
	Handler GetIBAGateByNameHandler
}

func (o *GetIBAGateByName) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetIBAGateByNameParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetIBAGateByNameBadRequestBody get i b a gate by name bad request body
//
// swagger:model GetIBAGateByNameBadRequestBody
type GetIBAGateByNameBadRequestBody struct {

	// code
	// Example: 300
	Code int64 `json:"code,omitempty"`

	// message
	// Example: Something bad happens.
	Message string `json:"message,omitempty"`
}

// Validate validates this get i b a gate by name bad request body
func (o *GetIBAGateByNameBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get i b a gate by name bad request body based on context it is used
func (o *GetIBAGateByNameBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetIBAGateByNameBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetIBAGateByNameBadRequestBody) UnmarshalBinary(b []byte) error {
	var res GetIBAGateByNameBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetIBAGateByNameNotFoundBody get i b a gate by name not found body
//
// swagger:model GetIBAGateByNameNotFoundBody
type GetIBAGateByNameNotFoundBody struct {

	// code
	// Example: 300
	Code int64 `json:"code,omitempty"`

	// message
	// Example: Something bad happens.
	Message string `json:"message,omitempty"`
}

// Validate validates this get i b a gate by name not found body
func (o *GetIBAGateByNameNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get i b a gate by name not found body based on context it is used
func (o *GetIBAGateByNameNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetIBAGateByNameNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetIBAGateByNameNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetIBAGateByNameNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetIBAGateByNameOKBody get i b a gate by name o k body
//
// swagger:model GetIBAGateByNameOKBody
type GetIBAGateByNameOKBody struct {

	// items
	Items *GetIBAGateByNameOKBodyItems `json:"items,omitempty"`
}

// Validate validates this get i b a gate by name o k body
func (o *GetIBAGateByNameOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetIBAGateByNameOKBody) validateItems(formats strfmt.Registry) error {
	if swag.IsZero(o.Items) { // not required
		return nil
	}

	if o.Items != nil {
		if err := o.Items.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getIBAGateByNameOK" + "." + "items")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get i b a gate by name o k body based on the context it is used
func (o *GetIBAGateByNameOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetIBAGateByNameOKBody) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	if o.Items != nil {
		if err := o.Items.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getIBAGateByNameOK" + "." + "items")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetIBAGateByNameOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetIBAGateByNameOKBody) UnmarshalBinary(b []byte) error {
	var res GetIBAGateByNameOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetIBAGateByNameOKBodyItems get i b a gate by name o k body items
//
// swagger:model GetIBAGateByNameOKBodyItems
type GetIBAGateByNameOKBodyItems struct {

	// Комментарий
	// Example: Some notes
	Comment string `json:"comment,omitempty"`

	// Имя IBA Gate
	// Example: p3apr3-pda-vc.p3.ia.nlmk
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this get i b a gate by name o k body items
func (o *GetIBAGateByNameOKBodyItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetIBAGateByNameOKBodyItems) validateName(formats strfmt.Registry) error {

	if err := validate.Required("getIBAGateByNameOK"+"."+"items"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get i b a gate by name o k body items based on context it is used
func (o *GetIBAGateByNameOKBodyItems) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetIBAGateByNameOKBodyItems) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetIBAGateByNameOKBodyItems) UnmarshalBinary(b []byte) error {
	var res GetIBAGateByNameOKBodyItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetIBAGateByNameUnauthorizedBody get i b a gate by name unauthorized body
//
// swagger:model GetIBAGateByNameUnauthorizedBody
type GetIBAGateByNameUnauthorizedBody struct {

	// code
	// Example: 300
	Code int64 `json:"code,omitempty"`

	// message
	// Example: Something bad happens.
	Message string `json:"message,omitempty"`
}

// Validate validates this get i b a gate by name unauthorized body
func (o *GetIBAGateByNameUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get i b a gate by name unauthorized body based on context it is used
func (o *GetIBAGateByNameUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetIBAGateByNameUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetIBAGateByNameUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res GetIBAGateByNameUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
