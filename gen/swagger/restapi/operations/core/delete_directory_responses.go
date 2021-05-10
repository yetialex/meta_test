// Code generated by go-swagger; DO NOT EDIT.

package core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/yetialex/meta_test/gen/swagger/models"
)

// DeleteDirectoryOKCode is the HTTP code returned for type DeleteDirectoryOK
const DeleteDirectoryOKCode int = 200

/*DeleteDirectoryOK OK

swagger:response deleteDirectoryOK
*/
type DeleteDirectoryOK struct {
}

// NewDeleteDirectoryOK creates DeleteDirectoryOK with default headers values
func NewDeleteDirectoryOK() *DeleteDirectoryOK {

	return &DeleteDirectoryOK{}
}

// WriteResponse to the client
func (o *DeleteDirectoryOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// DeleteDirectoryBadRequestCode is the HTTP code returned for type DeleteDirectoryBadRequest
const DeleteDirectoryBadRequestCode int = 400

/*DeleteDirectoryBadRequest Bad params supplied

swagger:response deleteDirectoryBadRequest
*/
type DeleteDirectoryBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewDeleteDirectoryBadRequest creates DeleteDirectoryBadRequest with default headers values
func NewDeleteDirectoryBadRequest() *DeleteDirectoryBadRequest {

	return &DeleteDirectoryBadRequest{}
}

// WithPayload adds the payload to the delete directory bad request response
func (o *DeleteDirectoryBadRequest) WithPayload(payload *models.ErrorResponse) *DeleteDirectoryBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete directory bad request response
func (o *DeleteDirectoryBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteDirectoryBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteDirectoryUnauthorizedCode is the HTTP code returned for type DeleteDirectoryUnauthorized
const DeleteDirectoryUnauthorizedCode int = 401

/*DeleteDirectoryUnauthorized Unauthorized

swagger:response deleteDirectoryUnauthorized
*/
type DeleteDirectoryUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewDeleteDirectoryUnauthorized creates DeleteDirectoryUnauthorized with default headers values
func NewDeleteDirectoryUnauthorized() *DeleteDirectoryUnauthorized {

	return &DeleteDirectoryUnauthorized{}
}

// WithPayload adds the payload to the delete directory unauthorized response
func (o *DeleteDirectoryUnauthorized) WithPayload(payload *models.ErrorResponse) *DeleteDirectoryUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete directory unauthorized response
func (o *DeleteDirectoryUnauthorized) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteDirectoryUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteDirectoryNotFoundCode is the HTTP code returned for type DeleteDirectoryNotFound
const DeleteDirectoryNotFoundCode int = 404

/*DeleteDirectoryNotFound Directory not found

swagger:response deleteDirectoryNotFound
*/
type DeleteDirectoryNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewDeleteDirectoryNotFound creates DeleteDirectoryNotFound with default headers values
func NewDeleteDirectoryNotFound() *DeleteDirectoryNotFound {

	return &DeleteDirectoryNotFound{}
}

// WithPayload adds the payload to the delete directory not found response
func (o *DeleteDirectoryNotFound) WithPayload(payload *models.ErrorResponse) *DeleteDirectoryNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete directory not found response
func (o *DeleteDirectoryNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteDirectoryNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteDirectoryInternalServerErrorCode is the HTTP code returned for type DeleteDirectoryInternalServerError
const DeleteDirectoryInternalServerErrorCode int = 500

/*DeleteDirectoryInternalServerError Internal server error

swagger:response deleteDirectoryInternalServerError
*/
type DeleteDirectoryInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewDeleteDirectoryInternalServerError creates DeleteDirectoryInternalServerError with default headers values
func NewDeleteDirectoryInternalServerError() *DeleteDirectoryInternalServerError {

	return &DeleteDirectoryInternalServerError{}
}

// WithPayload adds the payload to the delete directory internal server error response
func (o *DeleteDirectoryInternalServerError) WithPayload(payload *models.ErrorResponse) *DeleteDirectoryInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete directory internal server error response
func (o *DeleteDirectoryInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteDirectoryInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
