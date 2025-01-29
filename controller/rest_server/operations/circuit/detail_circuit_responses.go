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

package circuit

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"ztna-core/ztna/controller/rest_model"
)

// DetailCircuitOKCode is the HTTP code returned for type DetailCircuitOK
const DetailCircuitOKCode int = 200

/*DetailCircuitOK A single circuit

swagger:response detailCircuitOK
*/
type DetailCircuitOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.DetailCircuitEnvelope `json:"body,omitempty"`
}

// NewDetailCircuitOK creates DetailCircuitOK with default headers values
func NewDetailCircuitOK() *DetailCircuitOK {

	return &DetailCircuitOK{}
}

// WithPayload adds the payload to the detail circuit o k response
func (o *DetailCircuitOK) WithPayload(payload *rest_model.DetailCircuitEnvelope) *DetailCircuitOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail circuit o k response
func (o *DetailCircuitOK) SetPayload(payload *rest_model.DetailCircuitEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailCircuitOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailCircuitUnauthorizedCode is the HTTP code returned for type DetailCircuitUnauthorized
const DetailCircuitUnauthorizedCode int = 401

/*DetailCircuitUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response detailCircuitUnauthorized
*/
type DetailCircuitUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailCircuitUnauthorized creates DetailCircuitUnauthorized with default headers values
func NewDetailCircuitUnauthorized() *DetailCircuitUnauthorized {

	return &DetailCircuitUnauthorized{}
}

// WithPayload adds the payload to the detail circuit unauthorized response
func (o *DetailCircuitUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailCircuitUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail circuit unauthorized response
func (o *DetailCircuitUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailCircuitUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailCircuitNotFoundCode is the HTTP code returned for type DetailCircuitNotFound
const DetailCircuitNotFoundCode int = 404

/*DetailCircuitNotFound The requested resource does not exist

swagger:response detailCircuitNotFound
*/
type DetailCircuitNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailCircuitNotFound creates DetailCircuitNotFound with default headers values
func NewDetailCircuitNotFound() *DetailCircuitNotFound {

	return &DetailCircuitNotFound{}
}

// WithPayload adds the payload to the detail circuit not found response
func (o *DetailCircuitNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailCircuitNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail circuit not found response
func (o *DetailCircuitNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailCircuitNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailCircuitTooManyRequestsCode is the HTTP code returned for type DetailCircuitTooManyRequests
const DetailCircuitTooManyRequestsCode int = 429

/*DetailCircuitTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response detailCircuitTooManyRequests
*/
type DetailCircuitTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailCircuitTooManyRequests creates DetailCircuitTooManyRequests with default headers values
func NewDetailCircuitTooManyRequests() *DetailCircuitTooManyRequests {

	return &DetailCircuitTooManyRequests{}
}

// WithPayload adds the payload to the detail circuit too many requests response
func (o *DetailCircuitTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailCircuitTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail circuit too many requests response
func (o *DetailCircuitTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailCircuitTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
