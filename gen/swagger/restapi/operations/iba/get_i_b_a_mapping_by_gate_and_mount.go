// Code generated by go-swagger; DO NOT EDIT.

package iba

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetIBAMappingByGateAndMountHandlerFunc turns a function with the right signature into a get i b a mapping by gate and mount handler
type GetIBAMappingByGateAndMountHandlerFunc func(GetIBAMappingByGateAndMountParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetIBAMappingByGateAndMountHandlerFunc) Handle(params GetIBAMappingByGateAndMountParams) middleware.Responder {
	return fn(params)
}

// GetIBAMappingByGateAndMountHandler interface for that can handle valid get i b a mapping by gate and mount params
type GetIBAMappingByGateAndMountHandler interface {
	Handle(GetIBAMappingByGateAndMountParams) middleware.Responder
}

// NewGetIBAMappingByGateAndMount creates a new http.Handler for the get i b a mapping by gate and mount operation
func NewGetIBAMappingByGateAndMount(ctx *middleware.Context, handler GetIBAMappingByGateAndMountHandler) *GetIBAMappingByGateAndMount {
	return &GetIBAMappingByGateAndMount{Context: ctx, Handler: handler}
}

/* GetIBAMappingByGateAndMount swagger:route GET /ibas/gates/${gate_name}/mnts/${mnt}/signals/mapping iba getIBAMappingByGateAndMount

Return IBA signal mapping by gate and mountpoint

*/
type GetIBAMappingByGateAndMount struct {
	Context *middleware.Context
	Handler GetIBAMappingByGateAndMountHandler
}

func (o *GetIBAMappingByGateAndMount) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetIBAMappingByGateAndMountParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetIBAMappingByGateAndMountBadRequestBody get i b a mapping by gate and mount bad request body
//
// swagger:model GetIBAMappingByGateAndMountBadRequestBody
type GetIBAMappingByGateAndMountBadRequestBody struct {

	// code
	// Example: 300
	Code int64 `json:"code,omitempty"`

	// message
	// Example: Something bad happens.
	Message string `json:"message,omitempty"`
}

// Validate validates this get i b a mapping by gate and mount bad request body
func (o *GetIBAMappingByGateAndMountBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get i b a mapping by gate and mount bad request body based on context it is used
func (o *GetIBAMappingByGateAndMountBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetIBAMappingByGateAndMountBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetIBAMappingByGateAndMountBadRequestBody) UnmarshalBinary(b []byte) error {
	var res GetIBAMappingByGateAndMountBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetIBAMappingByGateAndMountInternalServerErrorBody get i b a mapping by gate and mount internal server error body
//
// swagger:model GetIBAMappingByGateAndMountInternalServerErrorBody
type GetIBAMappingByGateAndMountInternalServerErrorBody struct {

	// code
	// Example: 300
	Code int64 `json:"code,omitempty"`

	// message
	// Example: Something bad happens.
	Message string `json:"message,omitempty"`
}

// Validate validates this get i b a mapping by gate and mount internal server error body
func (o *GetIBAMappingByGateAndMountInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get i b a mapping by gate and mount internal server error body based on context it is used
func (o *GetIBAMappingByGateAndMountInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetIBAMappingByGateAndMountInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetIBAMappingByGateAndMountInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res GetIBAMappingByGateAndMountInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetIBAMappingByGateAndMountNotFoundBody get i b a mapping by gate and mount not found body
//
// swagger:model GetIBAMappingByGateAndMountNotFoundBody
type GetIBAMappingByGateAndMountNotFoundBody struct {

	// code
	// Example: 300
	Code int64 `json:"code,omitempty"`

	// message
	// Example: Something bad happens.
	Message string `json:"message,omitempty"`
}

// Validate validates this get i b a mapping by gate and mount not found body
func (o *GetIBAMappingByGateAndMountNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get i b a mapping by gate and mount not found body based on context it is used
func (o *GetIBAMappingByGateAndMountNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetIBAMappingByGateAndMountNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetIBAMappingByGateAndMountNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetIBAMappingByGateAndMountNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetIBAMappingByGateAndMountOKBody get i b a mapping by gate and mount o k body
//
// swagger:model GetIBAMappingByGateAndMountOKBody
type GetIBAMappingByGateAndMountOKBody struct {

	// items
	Items []*GetIBAMappingByGateAndMountOKBodyItemsItems0 `json:"items"`
}

// Validate validates this get i b a mapping by gate and mount o k body
func (o *GetIBAMappingByGateAndMountOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetIBAMappingByGateAndMountOKBody) validateItems(formats strfmt.Registry) error {
	if swag.IsZero(o.Items) { // not required
		return nil
	}

	for i := 0; i < len(o.Items); i++ {
		if swag.IsZero(o.Items[i]) { // not required
			continue
		}

		if o.Items[i] != nil {
			if err := o.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getIBAMappingByGateAndMountOK" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get i b a mapping by gate and mount o k body based on the context it is used
func (o *GetIBAMappingByGateAndMountOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetIBAMappingByGateAndMountOKBody) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Items); i++ {

		if o.Items[i] != nil {
			if err := o.Items[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getIBAMappingByGateAndMountOK" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetIBAMappingByGateAndMountOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetIBAMappingByGateAndMountOKBody) UnmarshalBinary(b []byte) error {
	var res GetIBAMappingByGateAndMountOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetIBAMappingByGateAndMountOKBodyItemsItems0 get i b a mapping by gate and mount o k body items items0
//
// swagger:model GetIBAMappingByGateAndMountOKBodyItemsItems0
type GetIBAMappingByGateAndMountOKBodyItemsItems0 struct {

	// is digital
	// Example: true
	// Required: true
	IsDigital *bool `json:"is_digital"`

	// module number
	// Example: 1
	// Required: true
	ModuleNumber *int64 `json:"module_number"`

	// number in module
	// Example: 102
	// Required: true
	NumberInModule *int64 `json:"number_in_module"`

	// storage id
	// Example: 1
	// Required: true
	StorageID *int64 `json:"storage_id"`
}

// Validate validates this get i b a mapping by gate and mount o k body items items0
func (o *GetIBAMappingByGateAndMountOKBodyItemsItems0) Validate(formats strfmt.Registry) error {
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

	if err := o.validateStorageID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetIBAMappingByGateAndMountOKBodyItemsItems0) validateIsDigital(formats strfmt.Registry) error {

	if err := validate.Required("is_digital", "body", o.IsDigital); err != nil {
		return err
	}

	return nil
}

func (o *GetIBAMappingByGateAndMountOKBodyItemsItems0) validateModuleNumber(formats strfmt.Registry) error {

	if err := validate.Required("module_number", "body", o.ModuleNumber); err != nil {
		return err
	}

	return nil
}

func (o *GetIBAMappingByGateAndMountOKBodyItemsItems0) validateNumberInModule(formats strfmt.Registry) error {

	if err := validate.Required("number_in_module", "body", o.NumberInModule); err != nil {
		return err
	}

	return nil
}

func (o *GetIBAMappingByGateAndMountOKBodyItemsItems0) validateStorageID(formats strfmt.Registry) error {

	if err := validate.Required("storage_id", "body", o.StorageID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get i b a mapping by gate and mount o k body items items0 based on context it is used
func (o *GetIBAMappingByGateAndMountOKBodyItemsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetIBAMappingByGateAndMountOKBodyItemsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetIBAMappingByGateAndMountOKBodyItemsItems0) UnmarshalBinary(b []byte) error {
	var res GetIBAMappingByGateAndMountOKBodyItemsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetIBAMappingByGateAndMountUnauthorizedBody get i b a mapping by gate and mount unauthorized body
//
// swagger:model GetIBAMappingByGateAndMountUnauthorizedBody
type GetIBAMappingByGateAndMountUnauthorizedBody struct {

	// code
	// Example: 300
	Code int64 `json:"code,omitempty"`

	// message
	// Example: Something bad happens.
	Message string `json:"message,omitempty"`
}

// Validate validates this get i b a mapping by gate and mount unauthorized body
func (o *GetIBAMappingByGateAndMountUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get i b a mapping by gate and mount unauthorized body based on context it is used
func (o *GetIBAMappingByGateAndMountUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetIBAMappingByGateAndMountUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetIBAMappingByGateAndMountUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res GetIBAMappingByGateAndMountUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}