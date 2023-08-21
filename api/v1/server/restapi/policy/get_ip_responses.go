// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/cilium/api/v1/models"
)

// GetIPOKCode is the HTTP code returned for type GetIPOK
const GetIPOKCode int = 200

/*
GetIPOK Success

swagger:response getIpOK
*/
type GetIPOK struct {

	/*
	  In: Body
	*/
	Payload []*models.IPListEntry `json:"body,omitempty"`
}

// NewGetIPOK creates GetIPOK with default headers values
func NewGetIPOK() *GetIPOK {

	return &GetIPOK{}
}

// WithPayload adds the payload to the get Ip o k response
func (o *GetIPOK) WithPayload(payload []*models.IPListEntry) *GetIPOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Ip o k response
func (o *GetIPOK) SetPayload(payload []*models.IPListEntry) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIPOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.IPListEntry, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetIPBadRequestCode is the HTTP code returned for type GetIPBadRequest
const GetIPBadRequestCode int = 400

/*
GetIPBadRequest Invalid request (error parsing parameters)

swagger:response getIpBadRequest
*/
type GetIPBadRequest struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewGetIPBadRequest creates GetIPBadRequest with default headers values
func NewGetIPBadRequest() *GetIPBadRequest {

	return &GetIPBadRequest{}
}

// WithPayload adds the payload to the get Ip bad request response
func (o *GetIPBadRequest) WithPayload(payload models.Error) *GetIPBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Ip bad request response
func (o *GetIPBadRequest) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIPBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetIPNotFoundCode is the HTTP code returned for type GetIPNotFound
const GetIPNotFoundCode int = 404

/*
GetIPNotFound No IP cache entries with provided parameters found

swagger:response getIpNotFound
*/
type GetIPNotFound struct {
}

// NewGetIPNotFound creates GetIPNotFound with default headers values
func NewGetIPNotFound() *GetIPNotFound {

	return &GetIPNotFound{}
}

// WriteResponse to the client
func (o *GetIPNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}