// Code generated by go-swagger; DO NOT EDIT.

package iba

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateOrUpdateIBAGateHandlerFunc turns a function with the right signature into a create or update i b a gate handler
type CreateOrUpdateIBAGateHandlerFunc func(CreateOrUpdateIBAGateParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateOrUpdateIBAGateHandlerFunc) Handle(params CreateOrUpdateIBAGateParams) middleware.Responder {
	return fn(params)
}

// CreateOrUpdateIBAGateHandler interface for that can handle valid create or update i b a gate params
type CreateOrUpdateIBAGateHandler interface {
	Handle(CreateOrUpdateIBAGateParams) middleware.Responder
}

// NewCreateOrUpdateIBAGate creates a new http.Handler for the create or update i b a gate operation
func NewCreateOrUpdateIBAGate(ctx *middleware.Context, handler CreateOrUpdateIBAGateHandler) *CreateOrUpdateIBAGate {
	return &CreateOrUpdateIBAGate{Context: ctx, Handler: handler}
}

/* CreateOrUpdateIBAGate swagger:route PUT /ibas/gates iba createOrUpdateIBAGate

Create or update Gate

*/
type CreateOrUpdateIBAGate struct {
	Context *middleware.Context
	Handler CreateOrUpdateIBAGateHandler
}

func (o *CreateOrUpdateIBAGate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateOrUpdateIBAGateParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
