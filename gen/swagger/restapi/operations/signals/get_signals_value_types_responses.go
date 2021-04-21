// Code generated by go-swagger; DO NOT EDIT.

package signals

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetSignalsValueTypesOKCode is the HTTP code returned for type GetSignalsValueTypesOK
const GetSignalsValueTypesOKCode int = 200

/*GetSignalsValueTypesOK OK

swagger:response getSignalsValueTypesOK
*/
type GetSignalsValueTypesOK struct {

	/*
	  In: Body
	*/
	Payload []string `json:"body,omitempty"`
}

// NewGetSignalsValueTypesOK creates GetSignalsValueTypesOK with default headers values
func NewGetSignalsValueTypesOK() *GetSignalsValueTypesOK {

	return &GetSignalsValueTypesOK{}
}

// WithPayload adds the payload to the get signals value types o k response
func (o *GetSignalsValueTypesOK) WithPayload(payload []string) *GetSignalsValueTypesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get signals value types o k response
func (o *GetSignalsValueTypesOK) SetPayload(payload []string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSignalsValueTypesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]string, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetSignalsValueTypesBadRequestCode is the HTTP code returned for type GetSignalsValueTypesBadRequest
const GetSignalsValueTypesBadRequestCode int = 400

/*GetSignalsValueTypesBadRequest Bad params supplied

swagger:response getSignalsValueTypesBadRequest
*/
type GetSignalsValueTypesBadRequest struct {

	/*
	  In: Body
	*/
	Payload *GetSignalsValueTypesBadRequestBody `json:"body,omitempty"`
}

// NewGetSignalsValueTypesBadRequest creates GetSignalsValueTypesBadRequest with default headers values
func NewGetSignalsValueTypesBadRequest() *GetSignalsValueTypesBadRequest {

	return &GetSignalsValueTypesBadRequest{}
}

// WithPayload adds the payload to the get signals value types bad request response
func (o *GetSignalsValueTypesBadRequest) WithPayload(payload *GetSignalsValueTypesBadRequestBody) *GetSignalsValueTypesBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get signals value types bad request response
func (o *GetSignalsValueTypesBadRequest) SetPayload(payload *GetSignalsValueTypesBadRequestBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSignalsValueTypesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetSignalsValueTypesUnauthorizedCode is the HTTP code returned for type GetSignalsValueTypesUnauthorized
const GetSignalsValueTypesUnauthorizedCode int = 401

/*GetSignalsValueTypesUnauthorized Unauthorized

swagger:response getSignalsValueTypesUnauthorized
*/
type GetSignalsValueTypesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *GetSignalsValueTypesUnauthorizedBody `json:"body,omitempty"`
}

// NewGetSignalsValueTypesUnauthorized creates GetSignalsValueTypesUnauthorized with default headers values
func NewGetSignalsValueTypesUnauthorized() *GetSignalsValueTypesUnauthorized {

	return &GetSignalsValueTypesUnauthorized{}
}

// WithPayload adds the payload to the get signals value types unauthorized response
func (o *GetSignalsValueTypesUnauthorized) WithPayload(payload *GetSignalsValueTypesUnauthorizedBody) *GetSignalsValueTypesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get signals value types unauthorized response
func (o *GetSignalsValueTypesUnauthorized) SetPayload(payload *GetSignalsValueTypesUnauthorizedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSignalsValueTypesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetSignalsValueTypesInternalServerErrorCode is the HTTP code returned for type GetSignalsValueTypesInternalServerError
const GetSignalsValueTypesInternalServerErrorCode int = 500

/*GetSignalsValueTypesInternalServerError Internal server error

swagger:response getSignalsValueTypesInternalServerError
*/
type GetSignalsValueTypesInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *GetSignalsValueTypesInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetSignalsValueTypesInternalServerError creates GetSignalsValueTypesInternalServerError with default headers values
func NewGetSignalsValueTypesInternalServerError() *GetSignalsValueTypesInternalServerError {

	return &GetSignalsValueTypesInternalServerError{}
}

// WithPayload adds the payload to the get signals value types internal server error response
func (o *GetSignalsValueTypesInternalServerError) WithPayload(payload *GetSignalsValueTypesInternalServerErrorBody) *GetSignalsValueTypesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get signals value types internal server error response
func (o *GetSignalsValueTypesInternalServerError) SetPayload(payload *GetSignalsValueTypesInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSignalsValueTypesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}