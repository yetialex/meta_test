// Code generated by go-swagger; DO NOT EDIT.

package iba

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetIBAMappingByServerIDOKCode is the HTTP code returned for type GetIBAMappingByServerIDOK
const GetIBAMappingByServerIDOKCode int = 200

/*GetIBAMappingByServerIDOK OK

swagger:response getIBAMappingByServerIdOK
*/
type GetIBAMappingByServerIDOK struct {

	/*
	  In: Body
	*/
	Payload *GetIBAMappingByServerIDOKBody `json:"body,omitempty"`
}

// NewGetIBAMappingByServerIDOK creates GetIBAMappingByServerIDOK with default headers values
func NewGetIBAMappingByServerIDOK() *GetIBAMappingByServerIDOK {

	return &GetIBAMappingByServerIDOK{}
}

// WithPayload adds the payload to the get i b a mapping by server Id o k response
func (o *GetIBAMappingByServerIDOK) WithPayload(payload *GetIBAMappingByServerIDOKBody) *GetIBAMappingByServerIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get i b a mapping by server Id o k response
func (o *GetIBAMappingByServerIDOK) SetPayload(payload *GetIBAMappingByServerIDOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIBAMappingByServerIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetIBAMappingByServerIDBadRequestCode is the HTTP code returned for type GetIBAMappingByServerIDBadRequest
const GetIBAMappingByServerIDBadRequestCode int = 400

/*GetIBAMappingByServerIDBadRequest Bad params supplied

swagger:response getIBAMappingByServerIdBadRequest
*/
type GetIBAMappingByServerIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload *GetIBAMappingByServerIDBadRequestBody `json:"body,omitempty"`
}

// NewGetIBAMappingByServerIDBadRequest creates GetIBAMappingByServerIDBadRequest with default headers values
func NewGetIBAMappingByServerIDBadRequest() *GetIBAMappingByServerIDBadRequest {

	return &GetIBAMappingByServerIDBadRequest{}
}

// WithPayload adds the payload to the get i b a mapping by server Id bad request response
func (o *GetIBAMappingByServerIDBadRequest) WithPayload(payload *GetIBAMappingByServerIDBadRequestBody) *GetIBAMappingByServerIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get i b a mapping by server Id bad request response
func (o *GetIBAMappingByServerIDBadRequest) SetPayload(payload *GetIBAMappingByServerIDBadRequestBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIBAMappingByServerIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetIBAMappingByServerIDUnauthorizedCode is the HTTP code returned for type GetIBAMappingByServerIDUnauthorized
const GetIBAMappingByServerIDUnauthorizedCode int = 401

/*GetIBAMappingByServerIDUnauthorized Unauthorized

swagger:response getIBAMappingByServerIdUnauthorized
*/
type GetIBAMappingByServerIDUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *GetIBAMappingByServerIDUnauthorizedBody `json:"body,omitempty"`
}

// NewGetIBAMappingByServerIDUnauthorized creates GetIBAMappingByServerIDUnauthorized with default headers values
func NewGetIBAMappingByServerIDUnauthorized() *GetIBAMappingByServerIDUnauthorized {

	return &GetIBAMappingByServerIDUnauthorized{}
}

// WithPayload adds the payload to the get i b a mapping by server Id unauthorized response
func (o *GetIBAMappingByServerIDUnauthorized) WithPayload(payload *GetIBAMappingByServerIDUnauthorizedBody) *GetIBAMappingByServerIDUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get i b a mapping by server Id unauthorized response
func (o *GetIBAMappingByServerIDUnauthorized) SetPayload(payload *GetIBAMappingByServerIDUnauthorizedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIBAMappingByServerIDUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetIBAMappingByServerIDNotFoundCode is the HTTP code returned for type GetIBAMappingByServerIDNotFound
const GetIBAMappingByServerIDNotFoundCode int = 404

/*GetIBAMappingByServerIDNotFound Server not found

swagger:response getIBAMappingByServerIdNotFound
*/
type GetIBAMappingByServerIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *GetIBAMappingByServerIDNotFoundBody `json:"body,omitempty"`
}

// NewGetIBAMappingByServerIDNotFound creates GetIBAMappingByServerIDNotFound with default headers values
func NewGetIBAMappingByServerIDNotFound() *GetIBAMappingByServerIDNotFound {

	return &GetIBAMappingByServerIDNotFound{}
}

// WithPayload adds the payload to the get i b a mapping by server Id not found response
func (o *GetIBAMappingByServerIDNotFound) WithPayload(payload *GetIBAMappingByServerIDNotFoundBody) *GetIBAMappingByServerIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get i b a mapping by server Id not found response
func (o *GetIBAMappingByServerIDNotFound) SetPayload(payload *GetIBAMappingByServerIDNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIBAMappingByServerIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetIBAMappingByServerIDInternalServerErrorCode is the HTTP code returned for type GetIBAMappingByServerIDInternalServerError
const GetIBAMappingByServerIDInternalServerErrorCode int = 500

/*GetIBAMappingByServerIDInternalServerError Internal server error

swagger:response getIBAMappingByServerIdInternalServerError
*/
type GetIBAMappingByServerIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *GetIBAMappingByServerIDInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetIBAMappingByServerIDInternalServerError creates GetIBAMappingByServerIDInternalServerError with default headers values
func NewGetIBAMappingByServerIDInternalServerError() *GetIBAMappingByServerIDInternalServerError {

	return &GetIBAMappingByServerIDInternalServerError{}
}

// WithPayload adds the payload to the get i b a mapping by server Id internal server error response
func (o *GetIBAMappingByServerIDInternalServerError) WithPayload(payload *GetIBAMappingByServerIDInternalServerErrorBody) *GetIBAMappingByServerIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get i b a mapping by server Id internal server error response
func (o *GetIBAMappingByServerIDInternalServerError) SetPayload(payload *GetIBAMappingByServerIDInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIBAMappingByServerIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}