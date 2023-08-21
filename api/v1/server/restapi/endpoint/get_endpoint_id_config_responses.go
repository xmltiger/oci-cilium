// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/cilium/api/v1/models"
)

// GetEndpointIDConfigOKCode is the HTTP code returned for type GetEndpointIDConfigOK
const GetEndpointIDConfigOKCode int = 200

/*
GetEndpointIDConfigOK Success

swagger:response getEndpointIdConfigOK
*/
type GetEndpointIDConfigOK struct {

	/*
	  In: Body
	*/
	Payload *models.EndpointConfigurationStatus `json:"body,omitempty"`
}

// NewGetEndpointIDConfigOK creates GetEndpointIDConfigOK with default headers values
func NewGetEndpointIDConfigOK() *GetEndpointIDConfigOK {

	return &GetEndpointIDConfigOK{}
}

// WithPayload adds the payload to the get endpoint Id config o k response
func (o *GetEndpointIDConfigOK) WithPayload(payload *models.EndpointConfigurationStatus) *GetEndpointIDConfigOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get endpoint Id config o k response
func (o *GetEndpointIDConfigOK) SetPayload(payload *models.EndpointConfigurationStatus) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEndpointIDConfigOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetEndpointIDConfigNotFoundCode is the HTTP code returned for type GetEndpointIDConfigNotFound
const GetEndpointIDConfigNotFoundCode int = 404

/*
GetEndpointIDConfigNotFound Endpoint not found

swagger:response getEndpointIdConfigNotFound
*/
type GetEndpointIDConfigNotFound struct {
}

// NewGetEndpointIDConfigNotFound creates GetEndpointIDConfigNotFound with default headers values
func NewGetEndpointIDConfigNotFound() *GetEndpointIDConfigNotFound {

	return &GetEndpointIDConfigNotFound{}
}

// WriteResponse to the client
func (o *GetEndpointIDConfigNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetEndpointIDConfigTooManyRequestsCode is the HTTP code returned for type GetEndpointIDConfigTooManyRequests
const GetEndpointIDConfigTooManyRequestsCode int = 429

/*
GetEndpointIDConfigTooManyRequests Rate-limiting too many requests in the given time frame

swagger:response getEndpointIdConfigTooManyRequests
*/
type GetEndpointIDConfigTooManyRequests struct {
}

// NewGetEndpointIDConfigTooManyRequests creates GetEndpointIDConfigTooManyRequests with default headers values
func NewGetEndpointIDConfigTooManyRequests() *GetEndpointIDConfigTooManyRequests {

	return &GetEndpointIDConfigTooManyRequests{}
}

// WriteResponse to the client
func (o *GetEndpointIDConfigTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(429)
}
