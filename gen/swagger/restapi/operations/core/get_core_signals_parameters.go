// Code generated by go-swagger; DO NOT EDIT.

package core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewGetCoreSignalsParams creates a new GetCoreSignalsParams object
//
// There are no default values defined in the spec.
func NewGetCoreSignalsParams() GetCoreSignalsParams {

	return GetCoreSignalsParams{}
}

// GetCoreSignalsParams contains all the bound params for the get core signals operation
// typically these are obtained from a http.Request
//
// swagger:parameters getCoreSignals
type GetCoreSignalsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Datetime as ISO 8601, for example 2018-03-20T09:12:28.123
	  In: query
	*/
	Datetime *strfmt.DateTime
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetCoreSignalsParams() beforehand.
func (o *GetCoreSignalsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qDatetime, qhkDatetime, _ := qs.GetOK("datetime")
	if err := o.bindDatetime(qDatetime, qhkDatetime, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindDatetime binds and validates parameter Datetime from query.
func (o *GetCoreSignalsParams) bindDatetime(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *GetCoreSignalsParams) validateDatetime(formats strfmt.Registry) error {

	if err := validate.FormatOf("datetime", "query", "date-time", o.Datetime.String(), formats); err != nil {
		return err
	}
	return nil
}
