// Code generated by go-swagger; DO NOT EDIT.

package iba

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/yetialex/meta_test/gen/swagger/models"
)

// GetIBAGateMntOKCode is the HTTP code returned for type GetIBAGateMntOK
const GetIBAGateMntOKCode int = 200

/*GetIBAGateMntOK OK

swagger:response getIBAGateMntOK
*/
type GetIBAGateMntOK struct {

	/*
	  In: Body
	*/
	Payload *GetIBAGateMntOKBody `json:"body,omitempty"`
}

// NewGetIBAGateMntOK creates GetIBAGateMntOK with default headers values
func NewGetIBAGateMntOK() *GetIBAGateMntOK {

	return &GetIBAGateMntOK{}
}

// WithPayload adds the payload to the get i b a gate mnt o k response
func (o *GetIBAGateMntOK) WithPayload(payload *GetIBAGateMntOKBody) *GetIBAGateMntOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get i b a gate mnt o k response
func (o *GetIBAGateMntOK) SetPayload(payload *GetIBAGateMntOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIBAGateMntOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetIBAGateMntBadRequestCode is the HTTP code returned for type GetIBAGateMntBadRequest
const GetIBAGateMntBadRequestCode int = 400

/*GetIBAGateMntBadRequest Bad params supplied

swagger:response getIBAGateMntBadRequest
*/
type GetIBAGateMntBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetIBAGateMntBadRequest creates GetIBAGateMntBadRequest with default headers values
func NewGetIBAGateMntBadRequest() *GetIBAGateMntBadRequest {

	return &GetIBAGateMntBadRequest{}
}

// WithPayload adds the payload to the get i b a gate mnt bad request response
func (o *GetIBAGateMntBadRequest) WithPayload(payload *models.ErrorResponse) *GetIBAGateMntBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get i b a gate mnt bad request response
func (o *GetIBAGateMntBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIBAGateMntBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetIBAGateMntUnauthorizedCode is the HTTP code returned for type GetIBAGateMntUnauthorized
const GetIBAGateMntUnauthorizedCode int = 401

/*GetIBAGateMntUnauthorized Unauthorized

swagger:response getIBAGateMntUnauthorized
*/
type GetIBAGateMntUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetIBAGateMntUnauthorized creates GetIBAGateMntUnauthorized with default headers values
func NewGetIBAGateMntUnauthorized() *GetIBAGateMntUnauthorized {

	return &GetIBAGateMntUnauthorized{}
}

// WithPayload adds the payload to the get i b a gate mnt unauthorized response
func (o *GetIBAGateMntUnauthorized) WithPayload(payload *models.ErrorResponse) *GetIBAGateMntUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get i b a gate mnt unauthorized response
func (o *GetIBAGateMntUnauthorized) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIBAGateMntUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetIBAGateMntNotFoundCode is the HTTP code returned for type GetIBAGateMntNotFound
const GetIBAGateMntNotFoundCode int = 404

/*GetIBAGateMntNotFound Mnt for Gate not found

swagger:response getIBAGateMntNotFound
*/
type GetIBAGateMntNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetIBAGateMntNotFound creates GetIBAGateMntNotFound with default headers values
func NewGetIBAGateMntNotFound() *GetIBAGateMntNotFound {

	return &GetIBAGateMntNotFound{}
}

// WithPayload adds the payload to the get i b a gate mnt not found response
func (o *GetIBAGateMntNotFound) WithPayload(payload *models.ErrorResponse) *GetIBAGateMntNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get i b a gate mnt not found response
func (o *GetIBAGateMntNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIBAGateMntNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
