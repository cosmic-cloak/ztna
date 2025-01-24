// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright NetFoundry Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// __          __              _
// \ \        / /             (_)
//  \ \  /\  / /_ _ _ __ _ __  _ _ __   __ _
//   \ \/  \/ / _` | '__| '_ \| | '_ \ / _` |
//    \  /\  / (_| | |  | | | | | | | | (_| | : This file is generated, do not edit it.
//     \/  \/ \__,_|_|  |_| |_|_|_| |_|\__, |
//                                      __/ |
//                                     |___/

package terminator

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cosmic-cloak/ztna/controller/rest_model"
)

// DetailTerminatorOKCode is the HTTP code returned for type DetailTerminatorOK
const DetailTerminatorOKCode int = 200

/*DetailTerminatorOK A single terminator

swagger:response detailTerminatorOK
*/
type DetailTerminatorOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.DetailTerminatorEnvelope `json:"body,omitempty"`
}

// NewDetailTerminatorOK creates DetailTerminatorOK with default headers values
func NewDetailTerminatorOK() *DetailTerminatorOK {

	return &DetailTerminatorOK{}
}

// WithPayload adds the payload to the detail terminator o k response
func (o *DetailTerminatorOK) WithPayload(payload *rest_model.DetailTerminatorEnvelope) *DetailTerminatorOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail terminator o k response
func (o *DetailTerminatorOK) SetPayload(payload *rest_model.DetailTerminatorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailTerminatorOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailTerminatorUnauthorizedCode is the HTTP code returned for type DetailTerminatorUnauthorized
const DetailTerminatorUnauthorizedCode int = 401

/*DetailTerminatorUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response detailTerminatorUnauthorized
*/
type DetailTerminatorUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailTerminatorUnauthorized creates DetailTerminatorUnauthorized with default headers values
func NewDetailTerminatorUnauthorized() *DetailTerminatorUnauthorized {

	return &DetailTerminatorUnauthorized{}
}

// WithPayload adds the payload to the detail terminator unauthorized response
func (o *DetailTerminatorUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailTerminatorUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail terminator unauthorized response
func (o *DetailTerminatorUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailTerminatorUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailTerminatorNotFoundCode is the HTTP code returned for type DetailTerminatorNotFound
const DetailTerminatorNotFoundCode int = 404

/*DetailTerminatorNotFound The requested resource does not exist

swagger:response detailTerminatorNotFound
*/
type DetailTerminatorNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailTerminatorNotFound creates DetailTerminatorNotFound with default headers values
func NewDetailTerminatorNotFound() *DetailTerminatorNotFound {

	return &DetailTerminatorNotFound{}
}

// WithPayload adds the payload to the detail terminator not found response
func (o *DetailTerminatorNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailTerminatorNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail terminator not found response
func (o *DetailTerminatorNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailTerminatorNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailTerminatorTooManyRequestsCode is the HTTP code returned for type DetailTerminatorTooManyRequests
const DetailTerminatorTooManyRequestsCode int = 429

/*DetailTerminatorTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response detailTerminatorTooManyRequests
*/
type DetailTerminatorTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailTerminatorTooManyRequests creates DetailTerminatorTooManyRequests with default headers values
func NewDetailTerminatorTooManyRequests() *DetailTerminatorTooManyRequests {

	return &DetailTerminatorTooManyRequests{}
}

// WithPayload adds the payload to the detail terminator too many requests response
func (o *DetailTerminatorTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailTerminatorTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail terminator too many requests response
func (o *DetailTerminatorTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailTerminatorTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
