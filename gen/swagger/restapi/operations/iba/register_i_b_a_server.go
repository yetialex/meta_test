// Code generated by go-swagger; DO NOT EDIT.

package iba

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// RegisterIBAServerHandlerFunc turns a function with the right signature into a register i b a server handler
type RegisterIBAServerHandlerFunc func(RegisterIBAServerParams) middleware.Responder

// Handle executing the request and returning a response
func (fn RegisterIBAServerHandlerFunc) Handle(params RegisterIBAServerParams) middleware.Responder {
	return fn(params)
}

// RegisterIBAServerHandler interface for that can handle valid register i b a server params
type RegisterIBAServerHandler interface {
	Handle(RegisterIBAServerParams) middleware.Responder
}

// NewRegisterIBAServer creates a new http.Handler for the register i b a server operation
func NewRegisterIBAServer(ctx *middleware.Context, handler RegisterIBAServerHandler) *RegisterIBAServer {
	return &RegisterIBAServer{Context: ctx, Handler: handler}
}

/* RegisterIBAServer swagger:route PUT /ibas/servers/ iba registerIBAServer

Register new Server

*/
type RegisterIBAServer struct {
	Context *middleware.Context
	Handler RegisterIBAServerHandler
}

func (o *RegisterIBAServer) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewRegisterIBAServerParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
