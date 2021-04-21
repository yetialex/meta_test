// Code generated by go-swagger; DO NOT EDIT.

package iba

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewRegisterSignalParams creates a new RegisterSignalParams object
//
// There are no default values defined in the spec.
func NewRegisterSignalParams() RegisterSignalParams {

	return RegisterSignalParams{}
}

// RegisterSignalParams contains all the bound params for the register signal operation
// typically these are obtained from a http.Request
//
// swagger:parameters registerSignal
type RegisterSignalParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Signal object.
	  Required: true
	  In: body
	*/
	Body RegisterSignalBody
	/*IBA Server ID
	  Required: true
	  In: path
	*/
	IbaServerID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewRegisterSignalParams() beforehand.
func (o *RegisterSignalParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body RegisterSignalBody
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("body", "body", ""))
			} else {
				res = append(res, errors.NewParseError("body", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(context.Background())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = body
			}
		}
	} else {
		res = append(res, errors.Required("body", "body", ""))
	}

	rIbaServerID, rhkIbaServerID, _ := route.Params.GetOK("iba_server_id")
	if err := o.bindIbaServerID(rIbaServerID, rhkIbaServerID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindIbaServerID binds and validates parameter IbaServerID from path.
func (o *RegisterSignalParams) bindIbaServerID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("iba_server_id", "path", "int64", raw)
	}
	o.IbaServerID = value

	return nil
}
