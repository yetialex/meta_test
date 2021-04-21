// Code generated by go-swagger; DO NOT EDIT.

package core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetRootNodeOKCode is the HTTP code returned for type GetRootNodeOK
const GetRootNodeOKCode int = 200

/*GetRootNodeOK OK

swagger:response getRootNodeOK
*/
type GetRootNodeOK struct {

	/*
	  In: Body
	*/
	Payload *GetRootNodeOKBody `json:"body,omitempty"`
}

// NewGetRootNodeOK creates GetRootNodeOK with default headers values
func NewGetRootNodeOK() *GetRootNodeOK {

	return &GetRootNodeOK{}
}

// WithPayload adds the payload to the get root node o k response
func (o *GetRootNodeOK) WithPayload(payload *GetRootNodeOKBody) *GetRootNodeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get root node o k response
func (o *GetRootNodeOK) SetPayload(payload *GetRootNodeOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRootNodeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetRootNodeBadRequestCode is the HTTP code returned for type GetRootNodeBadRequest
const GetRootNodeBadRequestCode int = 400

/*GetRootNodeBadRequest Bad params supplied

swagger:response getRootNodeBadRequest
*/
type GetRootNodeBadRequest struct {

	/*
	  In: Body
	*/
	Payload *GetRootNodeBadRequestBody `json:"body,omitempty"`
}

// NewGetRootNodeBadRequest creates GetRootNodeBadRequest with default headers values
func NewGetRootNodeBadRequest() *GetRootNodeBadRequest {

	return &GetRootNodeBadRequest{}
}

// WithPayload adds the payload to the get root node bad request response
func (o *GetRootNodeBadRequest) WithPayload(payload *GetRootNodeBadRequestBody) *GetRootNodeBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get root node bad request response
func (o *GetRootNodeBadRequest) SetPayload(payload *GetRootNodeBadRequestBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRootNodeBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetRootNodeUnauthorizedCode is the HTTP code returned for type GetRootNodeUnauthorized
const GetRootNodeUnauthorizedCode int = 401

/*GetRootNodeUnauthorized Unauthorized

swagger:response getRootNodeUnauthorized
*/
type GetRootNodeUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *GetRootNodeUnauthorizedBody `json:"body,omitempty"`
}

// NewGetRootNodeUnauthorized creates GetRootNodeUnauthorized with default headers values
func NewGetRootNodeUnauthorized() *GetRootNodeUnauthorized {

	return &GetRootNodeUnauthorized{}
}

// WithPayload adds the payload to the get root node unauthorized response
func (o *GetRootNodeUnauthorized) WithPayload(payload *GetRootNodeUnauthorizedBody) *GetRootNodeUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get root node unauthorized response
func (o *GetRootNodeUnauthorized) SetPayload(payload *GetRootNodeUnauthorizedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRootNodeUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetRootNodeInternalServerErrorCode is the HTTP code returned for type GetRootNodeInternalServerError
const GetRootNodeInternalServerErrorCode int = 500

/*GetRootNodeInternalServerError Internal server error

swagger:response getRootNodeInternalServerError
*/
type GetRootNodeInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *GetRootNodeInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetRootNodeInternalServerError creates GetRootNodeInternalServerError with default headers values
func NewGetRootNodeInternalServerError() *GetRootNodeInternalServerError {

	return &GetRootNodeInternalServerError{}
}

// WithPayload adds the payload to the get root node internal server error response
func (o *GetRootNodeInternalServerError) WithPayload(payload *GetRootNodeInternalServerErrorBody) *GetRootNodeInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get root node internal server error response
func (o *GetRootNodeInternalServerError) SetPayload(payload *GetRootNodeInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRootNodeInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
