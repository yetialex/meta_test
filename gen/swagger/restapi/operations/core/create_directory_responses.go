// Code generated by go-swagger; DO NOT EDIT.

package core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/yetialex/meta_test/gen/swagger/models"
)

// CreateDirectoryOKCode is the HTTP code returned for type CreateDirectoryOK
const CreateDirectoryOKCode int = 200

/*CreateDirectoryOK OK

swagger:response createDirectoryOK
*/
type CreateDirectoryOK struct {

	/*
	  In: Body
	*/
	Payload *models.DirectoryObject `json:"body,omitempty"`
}

// NewCreateDirectoryOK creates CreateDirectoryOK with default headers values
func NewCreateDirectoryOK() *CreateDirectoryOK {

	return &CreateDirectoryOK{}
}

// WithPayload adds the payload to the create directory o k response
func (o *CreateDirectoryOK) WithPayload(payload *models.DirectoryObject) *CreateDirectoryOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create directory o k response
func (o *CreateDirectoryOK) SetPayload(payload *models.DirectoryObject) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateDirectoryOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateDirectoryBadRequestCode is the HTTP code returned for type CreateDirectoryBadRequest
const CreateDirectoryBadRequestCode int = 400

/*CreateDirectoryBadRequest Bad params supplied

swagger:response createDirectoryBadRequest
*/
type CreateDirectoryBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCreateDirectoryBadRequest creates CreateDirectoryBadRequest with default headers values
func NewCreateDirectoryBadRequest() *CreateDirectoryBadRequest {

	return &CreateDirectoryBadRequest{}
}

// WithPayload adds the payload to the create directory bad request response
func (o *CreateDirectoryBadRequest) WithPayload(payload *models.ErrorResponse) *CreateDirectoryBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create directory bad request response
func (o *CreateDirectoryBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateDirectoryBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateDirectoryUnauthorizedCode is the HTTP code returned for type CreateDirectoryUnauthorized
const CreateDirectoryUnauthorizedCode int = 401

/*CreateDirectoryUnauthorized Unauthorized

swagger:response createDirectoryUnauthorized
*/
type CreateDirectoryUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCreateDirectoryUnauthorized creates CreateDirectoryUnauthorized with default headers values
func NewCreateDirectoryUnauthorized() *CreateDirectoryUnauthorized {

	return &CreateDirectoryUnauthorized{}
}

// WithPayload adds the payload to the create directory unauthorized response
func (o *CreateDirectoryUnauthorized) WithPayload(payload *models.ErrorResponse) *CreateDirectoryUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create directory unauthorized response
func (o *CreateDirectoryUnauthorized) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateDirectoryUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateDirectoryConflictCode is the HTTP code returned for type CreateDirectoryConflict
const CreateDirectoryConflictCode int = 409

/*CreateDirectoryConflict Conflict, gate already registered

swagger:response createDirectoryConflict
*/
type CreateDirectoryConflict struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCreateDirectoryConflict creates CreateDirectoryConflict with default headers values
func NewCreateDirectoryConflict() *CreateDirectoryConflict {

	return &CreateDirectoryConflict{}
}

// WithPayload adds the payload to the create directory conflict response
func (o *CreateDirectoryConflict) WithPayload(payload *models.ErrorResponse) *CreateDirectoryConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create directory conflict response
func (o *CreateDirectoryConflict) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateDirectoryConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}