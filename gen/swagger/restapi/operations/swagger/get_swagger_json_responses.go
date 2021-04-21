// Code generated by go-swagger; DO NOT EDIT.

package swagger

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetSwaggerJSONOKCode is the HTTP code returned for type GetSwaggerJSONOK
const GetSwaggerJSONOKCode int = 200

/*GetSwaggerJSONOK OK

swagger:response getSwaggerJsonOK
*/
type GetSwaggerJSONOK struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewGetSwaggerJSONOK creates GetSwaggerJSONOK with default headers values
func NewGetSwaggerJSONOK() *GetSwaggerJSONOK {

	return &GetSwaggerJSONOK{}
}

// WithPayload adds the payload to the get swagger Json o k response
func (o *GetSwaggerJSONOK) WithPayload(payload interface{}) *GetSwaggerJSONOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get swagger Json o k response
func (o *GetSwaggerJSONOK) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSwaggerJSONOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetSwaggerJSONUnauthorizedCode is the HTTP code returned for type GetSwaggerJSONUnauthorized
const GetSwaggerJSONUnauthorizedCode int = 401

/*GetSwaggerJSONUnauthorized Unauthorized

swagger:response getSwaggerJsonUnauthorized
*/
type GetSwaggerJSONUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *GetSwaggerJSONUnauthorizedBody `json:"body,omitempty"`
}

// NewGetSwaggerJSONUnauthorized creates GetSwaggerJSONUnauthorized with default headers values
func NewGetSwaggerJSONUnauthorized() *GetSwaggerJSONUnauthorized {

	return &GetSwaggerJSONUnauthorized{}
}

// WithPayload adds the payload to the get swagger Json unauthorized response
func (o *GetSwaggerJSONUnauthorized) WithPayload(payload *GetSwaggerJSONUnauthorizedBody) *GetSwaggerJSONUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get swagger Json unauthorized response
func (o *GetSwaggerJSONUnauthorized) SetPayload(payload *GetSwaggerJSONUnauthorizedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSwaggerJSONUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetSwaggerJSONInternalServerErrorCode is the HTTP code returned for type GetSwaggerJSONInternalServerError
const GetSwaggerJSONInternalServerErrorCode int = 500

/*GetSwaggerJSONInternalServerError Internal server error

swagger:response getSwaggerJsonInternalServerError
*/
type GetSwaggerJSONInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *GetSwaggerJSONInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetSwaggerJSONInternalServerError creates GetSwaggerJSONInternalServerError with default headers values
func NewGetSwaggerJSONInternalServerError() *GetSwaggerJSONInternalServerError {

	return &GetSwaggerJSONInternalServerError{}
}

// WithPayload adds the payload to the get swagger Json internal server error response
func (o *GetSwaggerJSONInternalServerError) WithPayload(payload *GetSwaggerJSONInternalServerErrorBody) *GetSwaggerJSONInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get swagger Json internal server error response
func (o *GetSwaggerJSONInternalServerError) SetPayload(payload *GetSwaggerJSONInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSwaggerJSONInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
