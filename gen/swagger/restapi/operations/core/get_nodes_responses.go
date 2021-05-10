// Code generated by go-swagger; DO NOT EDIT.

package core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/yetialex/meta_test/gen/swagger/models"
)

// GetNodesOKCode is the HTTP code returned for type GetNodesOK
const GetNodesOKCode int = 200

/*GetNodesOK OK

swagger:response getNodesOK
*/
type GetNodesOK struct {

	/*
	  In: Body
	*/
	Payload *GetNodesOKBody `json:"body,omitempty"`
}

// NewGetNodesOK creates GetNodesOK with default headers values
func NewGetNodesOK() *GetNodesOK {

	return &GetNodesOK{}
}

// WithPayload adds the payload to the get nodes o k response
func (o *GetNodesOK) WithPayload(payload *GetNodesOKBody) *GetNodesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get nodes o k response
func (o *GetNodesOK) SetPayload(payload *GetNodesOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNodesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetNodesBadRequestCode is the HTTP code returned for type GetNodesBadRequest
const GetNodesBadRequestCode int = 400

/*GetNodesBadRequest Bad params supplied

swagger:response getNodesBadRequest
*/
type GetNodesBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetNodesBadRequest creates GetNodesBadRequest with default headers values
func NewGetNodesBadRequest() *GetNodesBadRequest {

	return &GetNodesBadRequest{}
}

// WithPayload adds the payload to the get nodes bad request response
func (o *GetNodesBadRequest) WithPayload(payload *models.ErrorResponse) *GetNodesBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get nodes bad request response
func (o *GetNodesBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNodesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetNodesUnauthorizedCode is the HTTP code returned for type GetNodesUnauthorized
const GetNodesUnauthorizedCode int = 401

/*GetNodesUnauthorized Unauthorized

swagger:response getNodesUnauthorized
*/
type GetNodesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetNodesUnauthorized creates GetNodesUnauthorized with default headers values
func NewGetNodesUnauthorized() *GetNodesUnauthorized {

	return &GetNodesUnauthorized{}
}

// WithPayload adds the payload to the get nodes unauthorized response
func (o *GetNodesUnauthorized) WithPayload(payload *models.ErrorResponse) *GetNodesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get nodes unauthorized response
func (o *GetNodesUnauthorized) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNodesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetNodesNotFoundCode is the HTTP code returned for type GetNodesNotFound
const GetNodesNotFoundCode int = 404

/*GetNodesNotFound Core nodes not found

swagger:response getNodesNotFound
*/
type GetNodesNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetNodesNotFound creates GetNodesNotFound with default headers values
func NewGetNodesNotFound() *GetNodesNotFound {

	return &GetNodesNotFound{}
}

// WithPayload adds the payload to the get nodes not found response
func (o *GetNodesNotFound) WithPayload(payload *models.ErrorResponse) *GetNodesNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get nodes not found response
func (o *GetNodesNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNodesNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}