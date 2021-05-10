// Code generated by go-swagger; DO NOT EDIT.

package core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdateDirectoryHandlerFunc turns a function with the right signature into a update directory handler
type UpdateDirectoryHandlerFunc func(UpdateDirectoryParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateDirectoryHandlerFunc) Handle(params UpdateDirectoryParams) middleware.Responder {
	return fn(params)
}

// UpdateDirectoryHandler interface for that can handle valid update directory params
type UpdateDirectoryHandler interface {
	Handle(UpdateDirectoryParams) middleware.Responder
}

// NewUpdateDirectory creates a new http.Handler for the update directory operation
func NewUpdateDirectory(ctx *middleware.Context, handler UpdateDirectoryHandler) *UpdateDirectory {
	return &UpdateDirectory{Context: ctx, Handler: handler}
}

/* UpdateDirectory swagger:route PATCH /core/directories/{directory_id} core directories updateDirectory

Update directory

*/
type UpdateDirectory struct {
	Context *middleware.Context
	Handler UpdateDirectoryHandler
}

func (o *UpdateDirectory) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateDirectoryParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
