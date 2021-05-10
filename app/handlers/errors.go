package handlers

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/yetialex/meta_test/gen/swagger/models"
)

type InternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

func NewCommonInternalServerError() *InternalServerError {
	return &InternalServerError{}
}

func (o *InternalServerError) WithPayload(payload *models.ErrorResponse) *InternalServerError {
	o.Payload = payload
	return o
}

func (o *InternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *InternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(int(o.Payload.Code))
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

type NotFoundServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

func NewCommonNotFoundServerError() *NotFoundServerError {
	return &NotFoundServerError{}
}

func (o *NotFoundServerError) WithPayload(payload *models.ErrorResponse) *NotFoundServerError {
	o.Payload = payload
	return o
}

func (o *NotFoundServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NotFoundServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(int(o.Payload.Code))
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

type UnauthorizedError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

func NewCommonUnauthorizedError() *UnauthorizedError {
	return &UnauthorizedError{}
}

func (o *UnauthorizedError) WithPayload(payload *models.ErrorResponse) *UnauthorizedError {
	o.Payload = payload
	return o
}

func (o *UnauthorizedError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UnauthorizedError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(int(o.Payload.Code))
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
