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

// CreateIBAGateHandlerFunc turns a function with the right signature into a create i b a gate handler
type CreateIBAGateHandlerFunc func(CreateIBAGateParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateIBAGateHandlerFunc) Handle(params CreateIBAGateParams) middleware.Responder {
	return fn(params)
}

// CreateIBAGateHandler interface for that can handle valid create i b a gate params
type CreateIBAGateHandler interface {
	Handle(CreateIBAGateParams) middleware.Responder
}

// NewCreateIBAGate creates a new http.Handler for the create i b a gate operation
func NewCreateIBAGate(ctx *middleware.Context, handler CreateIBAGateHandler) *CreateIBAGate {
	return &CreateIBAGate{Context: ctx, Handler: handler}
}

/* CreateIBAGate swagger:route PUT /ibas/gates/${gate_name} iba createIBAGate

Create or update Gate

*/
type CreateIBAGate struct {
	Context *middleware.Context
	Handler CreateIBAGateHandler
}

func (o *CreateIBAGate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateIBAGateParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// CreateIBAGateBody create i b a gate body
//
// swagger:model CreateIBAGateBody
type CreateIBAGateBody struct {

	// comment
	Comment string `json:"comment,omitempty"`
}

// Validate validates this create i b a gate body
func (o *CreateIBAGateBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create i b a gate body based on context it is used
func (o *CreateIBAGateBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateIBAGateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateIBAGateBody) UnmarshalBinary(b []byte) error {
	var res CreateIBAGateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
