// Code generated by go-swagger; DO NOT EDIT.

package core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeleteNodeHandlerFunc turns a function with the right signature into a delete node handler
type DeleteNodeHandlerFunc func(DeleteNodeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteNodeHandlerFunc) Handle(params DeleteNodeParams) middleware.Responder {
	return fn(params)
}

// DeleteNodeHandler interface for that can handle valid delete node params
type DeleteNodeHandler interface {
	Handle(DeleteNodeParams) middleware.Responder
}

// NewDeleteNode creates a new http.Handler for the delete node operation
func NewDeleteNode(ctx *middleware.Context, handler DeleteNodeHandler) *DeleteNode {
	return &DeleteNode{Context: ctx, Handler: handler}
}

/* DeleteNode swagger:route DELETE /core/nodes/{node_id} core nodes deleteNode

Delete node

*/
type DeleteNode struct {
	Context *middleware.Context
	Handler DeleteNodeHandler
}

func (o *DeleteNode) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteNodeParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// DeleteNodeBadRequestBody delete node bad request body
//
// swagger:model DeleteNodeBadRequestBody
type DeleteNodeBadRequestBody struct {

	// code
	// Example: 300
	Code int64 `json:"code,omitempty"`

	// message
	// Example: Something bad happens.
	Message string `json:"message,omitempty"`
}

// Validate validates this delete node bad request body
func (o *DeleteNodeBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete node bad request body based on context it is used
func (o *DeleteNodeBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteNodeBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteNodeBadRequestBody) UnmarshalBinary(b []byte) error {
	var res DeleteNodeBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// DeleteNodeInternalServerErrorBody delete node internal server error body
//
// swagger:model DeleteNodeInternalServerErrorBody
type DeleteNodeInternalServerErrorBody struct {

	// code
	// Example: 300
	Code int64 `json:"code,omitempty"`

	// message
	// Example: Something bad happens.
	Message string `json:"message,omitempty"`
}

// Validate validates this delete node internal server error body
func (o *DeleteNodeInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete node internal server error body based on context it is used
func (o *DeleteNodeInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteNodeInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteNodeInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res DeleteNodeInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// DeleteNodeNotFoundBody delete node not found body
//
// swagger:model DeleteNodeNotFoundBody
type DeleteNodeNotFoundBody struct {

	// code
	// Example: 300
	Code int64 `json:"code,omitempty"`

	// message
	// Example: Something bad happens.
	Message string `json:"message,omitempty"`
}

// Validate validates this delete node not found body
func (o *DeleteNodeNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete node not found body based on context it is used
func (o *DeleteNodeNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteNodeNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteNodeNotFoundBody) UnmarshalBinary(b []byte) error {
	var res DeleteNodeNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// DeleteNodeUnauthorizedBody delete node unauthorized body
//
// swagger:model DeleteNodeUnauthorizedBody
type DeleteNodeUnauthorizedBody struct {

	// code
	// Example: 300
	Code int64 `json:"code,omitempty"`

	// message
	// Example: Something bad happens.
	Message string `json:"message,omitempty"`
}

// Validate validates this delete node unauthorized body
func (o *DeleteNodeUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete node unauthorized body based on context it is used
func (o *DeleteNodeUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteNodeUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteNodeUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res DeleteNodeUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
