// Code generated by go-swagger; DO NOT EDIT.

package signals

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetSignalsClassesOKCode is the HTTP code returned for type GetSignalsClassesOK
const GetSignalsClassesOKCode int = 200

/*GetSignalsClassesOK OK

swagger:response getSignalsClassesOK
*/
type GetSignalsClassesOK struct {

	/*
	  In: Body
	*/
	Payload []string `json:"body,omitempty"`
}

// NewGetSignalsClassesOK creates GetSignalsClassesOK with default headers values
func NewGetSignalsClassesOK() *GetSignalsClassesOK {

	return &GetSignalsClassesOK{}
}

// WithPayload adds the payload to the get signals classes o k response
func (o *GetSignalsClassesOK) WithPayload(payload []string) *GetSignalsClassesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get signals classes o k response
func (o *GetSignalsClassesOK) SetPayload(payload []string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSignalsClassesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetSignalsClassesBadRequestCode is the HTTP code returned for type GetSignalsClassesBadRequest
const GetSignalsClassesBadRequestCode int = 400

/*GetSignalsClassesBadRequest Bad params supplied

swagger:response getSignalsClassesBadRequest
*/
type GetSignalsClassesBadRequest struct {

	/*
	  In: Body
	*/
	Payload *GetSignalsClassesBadRequestBody `json:"body,omitempty"`
}

// NewGetSignalsClassesBadRequest creates GetSignalsClassesBadRequest with default headers values
func NewGetSignalsClassesBadRequest() *GetSignalsClassesBadRequest {

	return &GetSignalsClassesBadRequest{}
}

// WithPayload adds the payload to the get signals classes bad request response
func (o *GetSignalsClassesBadRequest) WithPayload(payload *GetSignalsClassesBadRequestBody) *GetSignalsClassesBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get signals classes bad request response
func (o *GetSignalsClassesBadRequest) SetPayload(payload *GetSignalsClassesBadRequestBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSignalsClassesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetSignalsClassesUnauthorizedCode is the HTTP code returned for type GetSignalsClassesUnauthorized
const GetSignalsClassesUnauthorizedCode int = 401

/*GetSignalsClassesUnauthorized Unauthorized

swagger:response getSignalsClassesUnauthorized
*/
type GetSignalsClassesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *GetSignalsClassesUnauthorizedBody `json:"body,omitempty"`
}

// NewGetSignalsClassesUnauthorized creates GetSignalsClassesUnauthorized with default headers values
func NewGetSignalsClassesUnauthorized() *GetSignalsClassesUnauthorized {

	return &GetSignalsClassesUnauthorized{}
}

// WithPayload adds the payload to the get signals classes unauthorized response
func (o *GetSignalsClassesUnauthorized) WithPayload(payload *GetSignalsClassesUnauthorizedBody) *GetSignalsClassesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get signals classes unauthorized response
func (o *GetSignalsClassesUnauthorized) SetPayload(payload *GetSignalsClassesUnauthorizedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSignalsClassesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetSignalsClassesInternalServerErrorCode is the HTTP code returned for type GetSignalsClassesInternalServerError
const GetSignalsClassesInternalServerErrorCode int = 500

/*GetSignalsClassesInternalServerError Internal server error

swagger:response getSignalsClassesInternalServerError
*/
type GetSignalsClassesInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *GetSignalsClassesInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetSignalsClassesInternalServerError creates GetSignalsClassesInternalServerError with default headers values
func NewGetSignalsClassesInternalServerError() *GetSignalsClassesInternalServerError {

	return &GetSignalsClassesInternalServerError{}
}

// WithPayload adds the payload to the get signals classes internal server error response
func (o *GetSignalsClassesInternalServerError) WithPayload(payload *GetSignalsClassesInternalServerErrorBody) *GetSignalsClassesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get signals classes internal server error response
func (o *GetSignalsClassesInternalServerError) SetPayload(payload *GetSignalsClassesInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSignalsClassesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}