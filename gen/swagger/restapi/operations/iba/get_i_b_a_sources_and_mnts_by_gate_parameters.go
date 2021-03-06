// Code generated by go-swagger; DO NOT EDIT.

package iba

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewGetIBASourcesAndMntsByGateParams creates a new GetIBASourcesAndMntsByGateParams object
//
// There are no default values defined in the spec.
func NewGetIBASourcesAndMntsByGateParams() GetIBASourcesAndMntsByGateParams {

	return GetIBASourcesAndMntsByGateParams{}
}

// GetIBASourcesAndMntsByGateParams contains all the bound params for the get i b a sources and mnts by gate operation
// typically these are obtained from a http.Request
//
// swagger:parameters getIBASourcesAndMntsByGate
type GetIBASourcesAndMntsByGateParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*IBA Gate name
	  Required: true
	  In: path
	*/
	GateName string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetIBASourcesAndMntsByGateParams() beforehand.
func (o *GetIBASourcesAndMntsByGateParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rGateName, rhkGateName, _ := route.Params.GetOK("gate_name")
	if err := o.bindGateName(rGateName, rhkGateName, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindGateName binds and validates parameter GateName from path.
func (o *GetIBASourcesAndMntsByGateParams) bindGateName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.GateName = raw

	return nil
}
