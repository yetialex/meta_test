// Code generated by go-swagger; DO NOT EDIT.

package iba

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateOrUpdateGateSourceByGateNameAndMntHandlerFunc turns a function with the right signature into a create or update gate source by gate name and mnt handler
type CreateOrUpdateGateSourceByGateNameAndMntHandlerFunc func(CreateOrUpdateGateSourceByGateNameAndMntParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateOrUpdateGateSourceByGateNameAndMntHandlerFunc) Handle(params CreateOrUpdateGateSourceByGateNameAndMntParams) middleware.Responder {
	return fn(params)
}

// CreateOrUpdateGateSourceByGateNameAndMntHandler interface for that can handle valid create or update gate source by gate name and mnt params
type CreateOrUpdateGateSourceByGateNameAndMntHandler interface {
	Handle(CreateOrUpdateGateSourceByGateNameAndMntParams) middleware.Responder
}

// NewCreateOrUpdateGateSourceByGateNameAndMnt creates a new http.Handler for the create or update gate source by gate name and mnt operation
func NewCreateOrUpdateGateSourceByGateNameAndMnt(ctx *middleware.Context, handler CreateOrUpdateGateSourceByGateNameAndMntHandler) *CreateOrUpdateGateSourceByGateNameAndMnt {
	return &CreateOrUpdateGateSourceByGateNameAndMnt{Context: ctx, Handler: handler}
}

/* CreateOrUpdateGateSourceByGateNameAndMnt swagger:route PUT /ibas/gates/{gate_name}/mnts/{mnt} iba createOrUpdateGateSourceByGateNameAndMnt

Create or Update Gate Source by IBA Gate name and mnt

*/
type CreateOrUpdateGateSourceByGateNameAndMnt struct {
	Context *middleware.Context
	Handler CreateOrUpdateGateSourceByGateNameAndMntHandler
}

func (o *CreateOrUpdateGateSourceByGateNameAndMnt) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateOrUpdateGateSourceByGateNameAndMntParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// CreateOrUpdateGateSourceByGateNameAndMntBody create or update gate source by gate name and mnt body
//
// swagger:model CreateOrUpdateGateSourceByGateNameAndMntBody
type CreateOrUpdateGateSourceByGateNameAndMntBody struct {

	// comment
	Comment string `json:"comment,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this create or update gate source by gate name and mnt body
func (o *CreateOrUpdateGateSourceByGateNameAndMntBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create or update gate source by gate name and mnt body based on context it is used
func (o *CreateOrUpdateGateSourceByGateNameAndMntBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateOrUpdateGateSourceByGateNameAndMntBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateOrUpdateGateSourceByGateNameAndMntBody) UnmarshalBinary(b []byte) error {
	var res CreateOrUpdateGateSourceByGateNameAndMntBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
