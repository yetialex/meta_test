// Code generated by go-swagger; DO NOT EDIT.

package core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteDirectoryHandlerFunc turns a function with the right signature into a delete directory handler
type DeleteDirectoryHandlerFunc func(DeleteDirectoryParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteDirectoryHandlerFunc) Handle(params DeleteDirectoryParams) middleware.Responder {
	return fn(params)
}

// DeleteDirectoryHandler interface for that can handle valid delete directory params
type DeleteDirectoryHandler interface {
	Handle(DeleteDirectoryParams) middleware.Responder
}

// NewDeleteDirectory creates a new http.Handler for the delete directory operation
func NewDeleteDirectory(ctx *middleware.Context, handler DeleteDirectoryHandler) *DeleteDirectory {
	return &DeleteDirectory{Context: ctx, Handler: handler}
}

/* DeleteDirectory swagger:route DELETE /core/directories/{directory_id} core directories deleteDirectory

Delete directory

*/
type DeleteDirectory struct {
	Context *middleware.Context
	Handler DeleteDirectoryHandler
}

func (o *DeleteDirectory) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteDirectoryParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
