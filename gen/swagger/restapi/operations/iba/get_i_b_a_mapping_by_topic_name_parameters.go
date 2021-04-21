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

// NewGetIBAMappingByTopicNameParams creates a new GetIBAMappingByTopicNameParams object
//
// There are no default values defined in the spec.
func NewGetIBAMappingByTopicNameParams() GetIBAMappingByTopicNameParams {

	return GetIBAMappingByTopicNameParams{}
}

// GetIBAMappingByTopicNameParams contains all the bound params for the get i b a mapping by topic name operation
// typically these are obtained from a http.Request
//
// swagger:parameters getIBAMappingByTopicName
type GetIBAMappingByTopicNameParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*IBA topic name
	  Required: true
	  In: path
	*/
	TopicName string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetIBAMappingByTopicNameParams() beforehand.
func (o *GetIBAMappingByTopicNameParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rTopicName, rhkTopicName, _ := route.Params.GetOK("topic_name")
	if err := o.bindTopicName(rTopicName, rhkTopicName, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindTopicName binds and validates parameter TopicName from path.
func (o *GetIBAMappingByTopicNameParams) bindTopicName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.TopicName = raw

	return nil
}
