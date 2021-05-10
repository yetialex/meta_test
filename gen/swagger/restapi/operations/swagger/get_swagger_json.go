// Code generated by go-swagger; DO NOT EDIT.

package swagger

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetSwaggerJSONHandlerFunc turns a function with the right signature into a get swagger JSON handler
type GetSwaggerJSONHandlerFunc func(GetSwaggerJSONParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetSwaggerJSONHandlerFunc) Handle(params GetSwaggerJSONParams) middleware.Responder {
	return fn(params)
}

// GetSwaggerJSONHandler interface for that can handle valid get swagger JSON params
type GetSwaggerJSONHandler interface {
	Handle(GetSwaggerJSONParams) middleware.Responder
}

// NewGetSwaggerJSON creates a new http.Handler for the get swagger JSON operation
func NewGetSwaggerJSON(ctx *middleware.Context, handler GetSwaggerJSONHandler) *GetSwaggerJSON {
	return &GetSwaggerJSON{Context: ctx, Handler: handler}
}

/* GetSwaggerJSON swagger:route GET /swagger.json swagger getSwaggerJson

Get swagger json

*/
type GetSwaggerJSON struct {
	Context *middleware.Context
	Handler GetSwaggerJSONHandler
}

func (o *GetSwaggerJSON) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetSwaggerJSONParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
