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

// CreateCoreDirectoryHandlerFunc turns a function with the right signature into a create core directory handler
type CreateCoreDirectoryHandlerFunc func(CreateCoreDirectoryParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateCoreDirectoryHandlerFunc) Handle(params CreateCoreDirectoryParams) middleware.Responder {
	return fn(params)
}

// CreateCoreDirectoryHandler interface for that can handle valid create core directory params
type CreateCoreDirectoryHandler interface {
	Handle(CreateCoreDirectoryParams) middleware.Responder
}

// NewCreateCoreDirectory creates a new http.Handler for the create core directory operation
func NewCreateCoreDirectory(ctx *middleware.Context, handler CreateCoreDirectoryHandler) *CreateCoreDirectory {
	return &CreateCoreDirectory{Context: ctx, Handler: handler}
}

/* CreateCoreDirectory swagger:route PUT /core/directories core directories createCoreDirectory

Create new core directory

*/
type CreateCoreDirectory struct {
	Context *middleware.Context
	Handler CreateCoreDirectoryHandler
}

func (o *CreateCoreDirectory) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateCoreDirectoryParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// CreateCoreDirectoryBadRequestBody create core directory bad request body
//
// swagger:model CreateCoreDirectoryBadRequestBody
type CreateCoreDirectoryBadRequestBody struct {

	// code
	// Example: 300
	Code int64 `json:"code,omitempty"`

	// message
	// Example: Something bad happens.
	Message string `json:"message,omitempty"`
}

// Validate validates this create core directory bad request body
func (o *CreateCoreDirectoryBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create core directory bad request body based on context it is used
func (o *CreateCoreDirectoryBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateCoreDirectoryBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateCoreDirectoryBadRequestBody) UnmarshalBinary(b []byte) error {
	var res CreateCoreDirectoryBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// CreateCoreDirectoryBody create core directory body
//
// swagger:model CreateCoreDirectoryBody
type CreateCoreDirectoryBody struct {

	// acl
	ACL interface{} `json:"acl,omitempty"`

	// description
	// Example: Description
	// Pattern: ^.{2,200}$
	Description string `json:"description,omitempty"`

	// name
	// Example: directory1
	// Required: true
	// Pattern: ^[A-Za-z0-9_.]{2,50}$
	Name *string `json:"name"`
}

// Validate validates this create core directory body
func (o *CreateCoreDirectoryBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDescription(formats); err != nil {
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

func (o *CreateCoreDirectoryBody) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(o.Description) { // not required
		return nil
	}

	if err := validate.Pattern("body"+"."+"description", "body", o.Description, `^.{2,200}$`); err != nil {
		return err
	}

	return nil
}

func (o *CreateCoreDirectoryBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	if err := validate.Pattern("body"+"."+"name", "body", *o.Name, `^[A-Za-z0-9_.]{2,50}$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create core directory body based on context it is used
func (o *CreateCoreDirectoryBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateCoreDirectoryBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateCoreDirectoryBody) UnmarshalBinary(b []byte) error {
	var res CreateCoreDirectoryBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// CreateCoreDirectoryConflictBody create core directory conflict body
//
// swagger:model CreateCoreDirectoryConflictBody
type CreateCoreDirectoryConflictBody struct {

	// code
	// Example: 300
	Code int64 `json:"code,omitempty"`

	// message
	// Example: Something bad happens.
	Message string `json:"message,omitempty"`
}

// Validate validates this create core directory conflict body
func (o *CreateCoreDirectoryConflictBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create core directory conflict body based on context it is used
func (o *CreateCoreDirectoryConflictBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateCoreDirectoryConflictBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateCoreDirectoryConflictBody) UnmarshalBinary(b []byte) error {
	var res CreateCoreDirectoryConflictBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// CreateCoreDirectoryOKBody create core directory o k body
//
// swagger:model CreateCoreDirectoryOKBody
type CreateCoreDirectoryOKBody struct {

	// ISO 8601 format
	// Example: 2018-03-20T09:12:28.123Z
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"created_at"`

	// description
	// Example: Description
	// Pattern: ^.{2,200}$
	Description string `json:"description,omitempty"`

	// id
	// Example: 1
	// Required: true
	ID *int64 `json:"id"`

	// name
	// Example: some_directory
	// Required: true
	Name *string `json:"name"`

	// ISO 8601 format
	// Example: 2018-03-20T09:12:28.123Z
	// Required: true
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"updated_at"`
}

// Validate validates this create core directory o k body
func (o *CreateCoreDirectoryOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
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

func (o *CreateCoreDirectoryOKBody) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("createCoreDirectoryOK"+"."+"created_at", "body", o.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("createCoreDirectoryOK"+"."+"created_at", "body", "date-time", o.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *CreateCoreDirectoryOKBody) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(o.Description) { // not required
		return nil
	}

	if err := validate.Pattern("createCoreDirectoryOK"+"."+"description", "body", o.Description, `^.{2,200}$`); err != nil {
		return err
	}

	return nil
}

func (o *CreateCoreDirectoryOKBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("createCoreDirectoryOK"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *CreateCoreDirectoryOKBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("createCoreDirectoryOK"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *CreateCoreDirectoryOKBody) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("createCoreDirectoryOK"+"."+"updated_at", "body", o.UpdatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("createCoreDirectoryOK"+"."+"updated_at", "body", "date-time", o.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create core directory o k body based on context it is used
func (o *CreateCoreDirectoryOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateCoreDirectoryOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateCoreDirectoryOKBody) UnmarshalBinary(b []byte) error {
	var res CreateCoreDirectoryOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// CreateCoreDirectoryUnauthorizedBody create core directory unauthorized body
//
// swagger:model CreateCoreDirectoryUnauthorizedBody
type CreateCoreDirectoryUnauthorizedBody struct {

	// code
	// Example: 300
	Code int64 `json:"code,omitempty"`

	// message
	// Example: Something bad happens.
	Message string `json:"message,omitempty"`
}

// Validate validates this create core directory unauthorized body
func (o *CreateCoreDirectoryUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create core directory unauthorized body based on context it is used
func (o *CreateCoreDirectoryUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateCoreDirectoryUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateCoreDirectoryUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res CreateCoreDirectoryUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}