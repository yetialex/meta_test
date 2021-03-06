// Code generated by go-swagger; DO NOT EDIT.

package core

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

	"github.com/yetialex/meta_test/gen/swagger/models"
)

// NewUpdateNodeParams creates a new UpdateNodeParams object
//
// There are no default values defined in the spec.
func NewUpdateNodeParams() UpdateNodeParams {

	return UpdateNodeParams{}
}

// UpdateNodeParams contains all the bound params for the update node operation
// typically these are obtained from a http.Request
//
// swagger:parameters updateNode
type UpdateNodeParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Update node params object.
	  Required: true
	  In: body
	*/
	Body *models.UpdateNode
	/*Datetime as ISO 8601, for example 2018-03-20T09:12:28.123
	  In: query
	*/
	Datetime *strfmt.DateTime
	/*Node ID
	  Required: true
	  In: path
	*/
	NodeID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUpdateNodeParams() beforehand.
func (o *UpdateNodeParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.UpdateNode
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
				o.Body = &body
			}
		}
	} else {
		res = append(res, errors.Required("body", "body", ""))
	}

	qDatetime, qhkDatetime, _ := qs.GetOK("datetime")
	if err := o.bindDatetime(qDatetime, qhkDatetime, route.Formats); err != nil {
		res = append(res, err)
	}

	rNodeID, rhkNodeID, _ := route.Params.GetOK("node_id")
	if err := o.bindNodeID(rNodeID, rhkNodeID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindDatetime binds and validates parameter Datetime from query.
func (o *UpdateNodeParams) bindDatetime(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	// Format: date-time
	value, err := formats.Parse("date-time", raw)
	if err != nil {
		return errors.InvalidType("datetime", "query", "strfmt.DateTime", raw)
	}
	o.Datetime = (value.(*strfmt.DateTime))

	if err := o.validateDatetime(formats); err != nil {
		return err
	}

	return nil
}

// validateDatetime carries on validations for parameter Datetime
func (o *UpdateNodeParams) validateDatetime(formats strfmt.Registry) error {

	if err := validate.FormatOf("datetime", "query", "date-time", o.Datetime.String(), formats); err != nil {
		return err
	}
	return nil
}

// bindNodeID binds and validates parameter NodeID from path.
func (o *UpdateNodeParams) bindNodeID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("node_id", "path", "int64", raw)
	}
	o.NodeID = value

	return nil
}
