// Code generated by go-swagger; DO NOT EDIT.

package core

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

// UpdateNodeHandlerFunc turns a function with the right signature into a update node handler
type UpdateNodeHandlerFunc func(UpdateNodeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateNodeHandlerFunc) Handle(params UpdateNodeParams) middleware.Responder {
	return fn(params)
}

// UpdateNodeHandler interface for that can handle valid update node params
type UpdateNodeHandler interface {
	Handle(UpdateNodeParams) middleware.Responder
}

// NewUpdateNode creates a new http.Handler for the update node operation
func NewUpdateNode(ctx *middleware.Context, handler UpdateNodeHandler) *UpdateNode {
	return &UpdateNode{Context: ctx, Handler: handler}
}

/* UpdateNode swagger:route PATCH /core/nodes/{node_id} core nodes updateNode

Update node

*/
type UpdateNode struct {
	Context *middleware.Context
	Handler UpdateNodeHandler
}

func (o *UpdateNode) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateNodeParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// UpdateNodeBadRequestBody update node bad request body
//
// swagger:model UpdateNodeBadRequestBody
type UpdateNodeBadRequestBody struct {

	// code
	// Example: 300
	Code int64 `json:"code,omitempty"`

	// message
	// Example: Something bad happens.
	Message string `json:"message,omitempty"`
}

// Validate validates this update node bad request body
func (o *UpdateNodeBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update node bad request body based on context it is used
func (o *UpdateNodeBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNodeBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNodeBadRequestBody) UnmarshalBinary(b []byte) error {
	var res UpdateNodeBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// UpdateNodeBody update node body
//
// swagger:model UpdateNodeBody
type UpdateNodeBody struct {

	// acl
	ACL interface{} `json:"acl,omitempty"`

	// comment
	// Example: Comment
	// Pattern: ^.{2,200}$
	Comment *string `json:"comment,omitempty"`

	// description
	// Example: Description
	// Pattern: ^.{2,200}$
	Description *string `json:"description,omitempty"`

	// directory id
	// Example: 1
	DirectoryID *int64 `json:"directory_id,omitempty"`

	// meta
	// Example: meta
	// Pattern: ^.{2,200}$
	Meta *string `json:"meta,omitempty"`

	// name
	// Example: node1
	// Pattern: ^[A-Za-z0-9_.]{2,50}$
	Name *string `json:"name,omitempty"`

	// prev id
	// Example: 1
	PrevID *int64 `json:"prev_id,omitempty"`

	// signal id
	// Example: 1
	SignalID *int64 `json:"signal_id,omitempty"`
}

// Validate validates this update node body
func (o *UpdateNodeBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateComment(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMeta(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNodeBody) validateComment(formats strfmt.Registry) error {
	if swag.IsZero(o.Comment) { // not required
		return nil
	}

	if err := validate.Pattern("body"+"."+"comment", "body", *o.Comment, `^.{2,200}$`); err != nil {
		return err
	}

	return nil
}

func (o *UpdateNodeBody) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(o.Description) { // not required
		return nil
	}

	if err := validate.Pattern("body"+"."+"description", "body", *o.Description, `^.{2,200}$`); err != nil {
		return err
	}

	return nil
}

func (o *UpdateNodeBody) validateMeta(formats strfmt.Registry) error {
	if swag.IsZero(o.Meta) { // not required
		return nil
	}

	if err := validate.Pattern("body"+"."+"meta", "body", *o.Meta, `^.{2,200}$`); err != nil {
		return err
	}

	return nil
}

func (o *UpdateNodeBody) validateName(formats strfmt.Registry) error {
	if swag.IsZero(o.Name) { // not required
		return nil
	}

	if err := validate.Pattern("body"+"."+"name", "body", *o.Name, `^[A-Za-z0-9_.]{2,50}$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update node body based on context it is used
func (o *UpdateNodeBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNodeBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNodeBody) UnmarshalBinary(b []byte) error {
	var res UpdateNodeBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// UpdateNodeInternalServerErrorBody update node internal server error body
//
// swagger:model UpdateNodeInternalServerErrorBody
type UpdateNodeInternalServerErrorBody struct {

	// code
	// Example: 300
	Code int64 `json:"code,omitempty"`

	// message
	// Example: Something bad happens.
	Message string `json:"message,omitempty"`
}

// Validate validates this update node internal server error body
func (o *UpdateNodeInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update node internal server error body based on context it is used
func (o *UpdateNodeInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNodeInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNodeInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res UpdateNodeInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// UpdateNodeNotFoundBody update node not found body
//
// swagger:model UpdateNodeNotFoundBody
type UpdateNodeNotFoundBody struct {

	// code
	// Example: 300
	Code int64 `json:"code,omitempty"`

	// message
	// Example: Something bad happens.
	Message string `json:"message,omitempty"`
}

// Validate validates this update node not found body
func (o *UpdateNodeNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update node not found body based on context it is used
func (o *UpdateNodeNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNodeNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNodeNotFoundBody) UnmarshalBinary(b []byte) error {
	var res UpdateNodeNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// UpdateNodeOKBody update node o k body
//
// swagger:model UpdateNodeOKBody
type UpdateNodeOKBody struct {

	// ISO 8601 format
	// Example: 2018-03-20T09:12:28.123Z
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"created_at"`

	// directory id
	// Example: 1
	DirectoryID int64 `json:"directory_id,omitempty"`

	// has children
	// Example: true
	HasChildren *bool `json:"has_children,omitempty"`

	// id
	// Example: 1
	// Required: true
	ID *int64 `json:"id"`

	// name
	// Example: some_node
	// Required: true
	Name *string `json:"name"`

	// prev id
	// Example: 0
	// Required: true
	PrevID *int64 `json:"prev_id"`

	// signal id
	// Example: 1
	SignalID int64 `json:"signal_id,omitempty"`

	// ISO 8601 format
	// Example: 2018-03-20T09:12:28.123Z
	// Required: true
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"updated_at"`
}

// Validate validates this update node o k body
func (o *UpdateNodeOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePrevID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNodeOKBody) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updateNodeOK"+"."+"created_at", "body", o.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("updateNodeOK"+"."+"created_at", "body", "date-time", o.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *UpdateNodeOKBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("updateNodeOK"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *UpdateNodeOKBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("updateNodeOK"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *UpdateNodeOKBody) validatePrevID(formats strfmt.Registry) error {

	if err := validate.Required("updateNodeOK"+"."+"prev_id", "body", o.PrevID); err != nil {
		return err
	}

	return nil
}

func (o *UpdateNodeOKBody) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updateNodeOK"+"."+"updated_at", "body", o.UpdatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("updateNodeOK"+"."+"updated_at", "body", "date-time", o.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update node o k body based on context it is used
func (o *UpdateNodeOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNodeOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNodeOKBody) UnmarshalBinary(b []byte) error {
	var res UpdateNodeOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// UpdateNodeUnauthorizedBody update node unauthorized body
//
// swagger:model UpdateNodeUnauthorizedBody
type UpdateNodeUnauthorizedBody struct {

	// code
	// Example: 300
	Code int64 `json:"code,omitempty"`

	// message
	// Example: Something bad happens.
	Message string `json:"message,omitempty"`
}

// Validate validates this update node unauthorized body
func (o *UpdateNodeUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update node unauthorized body based on context it is used
func (o *UpdateNodeUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNodeUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNodeUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res UpdateNodeUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}