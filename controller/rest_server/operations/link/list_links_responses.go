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

package link

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"ztna-core/ztna/controller/rest_model"
)

// ListLinksOKCode is the HTTP code returned for type ListLinksOK
const ListLinksOKCode int = 200

/*ListLinksOK A list of links

swagger:response listLinksOK
*/
type ListLinksOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.ListLinksEnvelope `json:"body,omitempty"`
}

// NewListLinksOK creates ListLinksOK with default headers values
func NewListLinksOK() *ListLinksOK {

	return &ListLinksOK{}
}

// WithPayload adds the payload to the list links o k response
func (o *ListLinksOK) WithPayload(payload *rest_model.ListLinksEnvelope) *ListLinksOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list links o k response
func (o *ListLinksOK) SetPayload(payload *rest_model.ListLinksEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListLinksOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListLinksUnauthorizedCode is the HTTP code returned for type ListLinksUnauthorized
const ListLinksUnauthorizedCode int = 401

/*ListLinksUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response listLinksUnauthorized
*/
type ListLinksUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListLinksUnauthorized creates ListLinksUnauthorized with default headers values
func NewListLinksUnauthorized() *ListLinksUnauthorized {

	return &ListLinksUnauthorized{}
}

// WithPayload adds the payload to the list links unauthorized response
func (o *ListLinksUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *ListLinksUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list links unauthorized response
func (o *ListLinksUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListLinksUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListLinksTooManyRequestsCode is the HTTP code returned for type ListLinksTooManyRequests
const ListLinksTooManyRequestsCode int = 429

/*ListLinksTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response listLinksTooManyRequests
*/
type ListLinksTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListLinksTooManyRequests creates ListLinksTooManyRequests with default headers values
func NewListLinksTooManyRequests() *ListLinksTooManyRequests {

	return &ListLinksTooManyRequests{}
}

// WithPayload adds the payload to the list links too many requests response
func (o *ListLinksTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *ListLinksTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list links too many requests response
func (o *ListLinksTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListLinksTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
