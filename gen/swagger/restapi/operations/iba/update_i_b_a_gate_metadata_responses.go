// Code generated by go-swagger; DO NOT EDIT.

package iba

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// UpdateIBAGateMetadataOKCode is the HTTP code returned for type UpdateIBAGateMetadataOK
const UpdateIBAGateMetadataOKCode int = 200

/*UpdateIBAGateMetadataOK OK

swagger:response updateIBAGateMetadataOK
*/
type UpdateIBAGateMetadataOK struct {

	/*
	  In: Body
	*/
	Payload *UpdateIBAGateMetadataOKBody `json:"body,omitempty"`
}

// NewUpdateIBAGateMetadataOK creates UpdateIBAGateMetadataOK with default headers values
func NewUpdateIBAGateMetadataOK() *UpdateIBAGateMetadataOK {

	return &UpdateIBAGateMetadataOK{}
}

// WithPayload adds the payload to the update i b a gate metadata o k response
func (o *UpdateIBAGateMetadataOK) WithPayload(payload *UpdateIBAGateMetadataOKBody) *UpdateIBAGateMetadataOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update i b a gate metadata o k response
func (o *UpdateIBAGateMetadataOK) SetPayload(payload *UpdateIBAGateMetadataOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateIBAGateMetadataOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateIBAGateMetadataBadRequestCode is the HTTP code returned for type UpdateIBAGateMetadataBadRequest
const UpdateIBAGateMetadataBadRequestCode int = 400

/*UpdateIBAGateMetadataBadRequest Bad params supplied

swagger:response updateIBAGateMetadataBadRequest
*/
type UpdateIBAGateMetadataBadRequest struct {

	/*
	  In: Body
	*/
	Payload *UpdateIBAGateMetadataBadRequestBody `json:"body,omitempty"`
}

// NewUpdateIBAGateMetadataBadRequest creates UpdateIBAGateMetadataBadRequest with default headers values
func NewUpdateIBAGateMetadataBadRequest() *UpdateIBAGateMetadataBadRequest {

	return &UpdateIBAGateMetadataBadRequest{}
}

// WithPayload adds the payload to the update i b a gate metadata bad request response
func (o *UpdateIBAGateMetadataBadRequest) WithPayload(payload *UpdateIBAGateMetadataBadRequestBody) *UpdateIBAGateMetadataBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update i b a gate metadata bad request response
func (o *UpdateIBAGateMetadataBadRequest) SetPayload(payload *UpdateIBAGateMetadataBadRequestBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateIBAGateMetadataBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateIBAGateMetadataUnauthorizedCode is the HTTP code returned for type UpdateIBAGateMetadataUnauthorized
const UpdateIBAGateMetadataUnauthorizedCode int = 401

/*UpdateIBAGateMetadataUnauthorized Unauthorized

swagger:response updateIBAGateMetadataUnauthorized
*/
type UpdateIBAGateMetadataUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *UpdateIBAGateMetadataUnauthorizedBody `json:"body,omitempty"`
}

// NewUpdateIBAGateMetadataUnauthorized creates UpdateIBAGateMetadataUnauthorized with default headers values
func NewUpdateIBAGateMetadataUnauthorized() *UpdateIBAGateMetadataUnauthorized {

	return &UpdateIBAGateMetadataUnauthorized{}
}

// WithPayload adds the payload to the update i b a gate metadata unauthorized response
func (o *UpdateIBAGateMetadataUnauthorized) WithPayload(payload *UpdateIBAGateMetadataUnauthorizedBody) *UpdateIBAGateMetadataUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update i b a gate metadata unauthorized response
func (o *UpdateIBAGateMetadataUnauthorized) SetPayload(payload *UpdateIBAGateMetadataUnauthorizedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateIBAGateMetadataUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateIBAGateMetadataNotFoundCode is the HTTP code returned for type UpdateIBAGateMetadataNotFound
const UpdateIBAGateMetadataNotFoundCode int = 404

/*UpdateIBAGateMetadataNotFound Not found

swagger:response updateIBAGateMetadataNotFound
*/
type UpdateIBAGateMetadataNotFound struct {

	/*
	  In: Body
	*/
	Payload *UpdateIBAGateMetadataNotFoundBody `json:"body,omitempty"`
}

// NewUpdateIBAGateMetadataNotFound creates UpdateIBAGateMetadataNotFound with default headers values
func NewUpdateIBAGateMetadataNotFound() *UpdateIBAGateMetadataNotFound {

	return &UpdateIBAGateMetadataNotFound{}
}

// WithPayload adds the payload to the update i b a gate metadata not found response
func (o *UpdateIBAGateMetadataNotFound) WithPayload(payload *UpdateIBAGateMetadataNotFoundBody) *UpdateIBAGateMetadataNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update i b a gate metadata not found response
func (o *UpdateIBAGateMetadataNotFound) SetPayload(payload *UpdateIBAGateMetadataNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateIBAGateMetadataNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}