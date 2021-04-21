// Code generated by go-swagger; DO NOT EDIT.

package iba

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetIBAServersOKCode is the HTTP code returned for type GetIBAServersOK
const GetIBAServersOKCode int = 200

/*GetIBAServersOK OK

swagger:response getIBAServersOK
*/
type GetIBAServersOK struct {

	/*
	  In: Body
	*/
	Payload *GetIBAServersOKBody `json:"body,omitempty"`
}

// NewGetIBAServersOK creates GetIBAServersOK with default headers values
func NewGetIBAServersOK() *GetIBAServersOK {

	return &GetIBAServersOK{}
}

// WithPayload adds the payload to the get i b a servers o k response
func (o *GetIBAServersOK) WithPayload(payload *GetIBAServersOKBody) *GetIBAServersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get i b a servers o k response
func (o *GetIBAServersOK) SetPayload(payload *GetIBAServersOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIBAServersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetIBAServersBadRequestCode is the HTTP code returned for type GetIBAServersBadRequest
const GetIBAServersBadRequestCode int = 400

/*GetIBAServersBadRequest Bad params supplied

swagger:response getIBAServersBadRequest
*/
type GetIBAServersBadRequest struct {

	/*
	  In: Body
	*/
	Payload *GetIBAServersBadRequestBody `json:"body,omitempty"`
}

// NewGetIBAServersBadRequest creates GetIBAServersBadRequest with default headers values
func NewGetIBAServersBadRequest() *GetIBAServersBadRequest {

	return &GetIBAServersBadRequest{}
}

// WithPayload adds the payload to the get i b a servers bad request response
func (o *GetIBAServersBadRequest) WithPayload(payload *GetIBAServersBadRequestBody) *GetIBAServersBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get i b a servers bad request response
func (o *GetIBAServersBadRequest) SetPayload(payload *GetIBAServersBadRequestBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIBAServersBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetIBAServersUnauthorizedCode is the HTTP code returned for type GetIBAServersUnauthorized
const GetIBAServersUnauthorizedCode int = 401

/*GetIBAServersUnauthorized Unauthorized

swagger:response getIBAServersUnauthorized
*/
type GetIBAServersUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *GetIBAServersUnauthorizedBody `json:"body,omitempty"`
}

// NewGetIBAServersUnauthorized creates GetIBAServersUnauthorized with default headers values
func NewGetIBAServersUnauthorized() *GetIBAServersUnauthorized {

	return &GetIBAServersUnauthorized{}
}

// WithPayload adds the payload to the get i b a servers unauthorized response
func (o *GetIBAServersUnauthorized) WithPayload(payload *GetIBAServersUnauthorizedBody) *GetIBAServersUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get i b a servers unauthorized response
func (o *GetIBAServersUnauthorized) SetPayload(payload *GetIBAServersUnauthorizedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIBAServersUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetIBAServersNotFoundCode is the HTTP code returned for type GetIBAServersNotFound
const GetIBAServersNotFoundCode int = 404

/*GetIBAServersNotFound IBA Servers not found

swagger:response getIBAServersNotFound
*/
type GetIBAServersNotFound struct {

	/*
	  In: Body
	*/
	Payload *GetIBAServersNotFoundBody `json:"body,omitempty"`
}

// NewGetIBAServersNotFound creates GetIBAServersNotFound with default headers values
func NewGetIBAServersNotFound() *GetIBAServersNotFound {

	return &GetIBAServersNotFound{}
}

// WithPayload adds the payload to the get i b a servers not found response
func (o *GetIBAServersNotFound) WithPayload(payload *GetIBAServersNotFoundBody) *GetIBAServersNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get i b a servers not found response
func (o *GetIBAServersNotFound) SetPayload(payload *GetIBAServersNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIBAServersNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
