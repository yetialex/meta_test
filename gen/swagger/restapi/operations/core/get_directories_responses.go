// Code generated by go-swagger; DO NOT EDIT.

package core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/yetialex/meta_test/gen/swagger/models"
)

// GetDirectoriesOKCode is the HTTP code returned for type GetDirectoriesOK
const GetDirectoriesOKCode int = 200

/*GetDirectoriesOK OK

swagger:response getDirectoriesOK
*/
type GetDirectoriesOK struct {

	/*
	  In: Body
	*/
	Payload *GetDirectoriesOKBody `json:"body,omitempty"`
}

// NewGetDirectoriesOK creates GetDirectoriesOK with default headers values
func NewGetDirectoriesOK() *GetDirectoriesOK {

	return &GetDirectoriesOK{}
}

// WithPayload adds the payload to the get directories o k response
func (o *GetDirectoriesOK) WithPayload(payload *GetDirectoriesOKBody) *GetDirectoriesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get directories o k response
func (o *GetDirectoriesOK) SetPayload(payload *GetDirectoriesOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDirectoriesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetDirectoriesBadRequestCode is the HTTP code returned for type GetDirectoriesBadRequest
const GetDirectoriesBadRequestCode int = 400

/*GetDirectoriesBadRequest Bad params supplied

swagger:response getDirectoriesBadRequest
*/
type GetDirectoriesBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetDirectoriesBadRequest creates GetDirectoriesBadRequest with default headers values
func NewGetDirectoriesBadRequest() *GetDirectoriesBadRequest {

	return &GetDirectoriesBadRequest{}
}

// WithPayload adds the payload to the get directories bad request response
func (o *GetDirectoriesBadRequest) WithPayload(payload *models.ErrorResponse) *GetDirectoriesBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get directories bad request response
func (o *GetDirectoriesBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDirectoriesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetDirectoriesUnauthorizedCode is the HTTP code returned for type GetDirectoriesUnauthorized
const GetDirectoriesUnauthorizedCode int = 401

/*GetDirectoriesUnauthorized Unauthorized

swagger:response getDirectoriesUnauthorized
*/
type GetDirectoriesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetDirectoriesUnauthorized creates GetDirectoriesUnauthorized with default headers values
func NewGetDirectoriesUnauthorized() *GetDirectoriesUnauthorized {

	return &GetDirectoriesUnauthorized{}
}

// WithPayload adds the payload to the get directories unauthorized response
func (o *GetDirectoriesUnauthorized) WithPayload(payload *models.ErrorResponse) *GetDirectoriesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get directories unauthorized response
func (o *GetDirectoriesUnauthorized) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDirectoriesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetDirectoriesNotFoundCode is the HTTP code returned for type GetDirectoriesNotFound
const GetDirectoriesNotFoundCode int = 404

/*GetDirectoriesNotFound Core directories not found

swagger:response getDirectoriesNotFound
*/
type GetDirectoriesNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetDirectoriesNotFound creates GetDirectoriesNotFound with default headers values
func NewGetDirectoriesNotFound() *GetDirectoriesNotFound {

	return &GetDirectoriesNotFound{}
}

// WithPayload adds the payload to the get directories not found response
func (o *GetDirectoriesNotFound) WithPayload(payload *models.ErrorResponse) *GetDirectoriesNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get directories not found response
func (o *GetDirectoriesNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDirectoriesNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
