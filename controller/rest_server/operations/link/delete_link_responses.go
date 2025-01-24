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

	"github.com/cosmic-cloak/ztna/controller/rest_model"
)

// DeleteLinkOKCode is the HTTP code returned for type DeleteLinkOK
const DeleteLinkOKCode int = 200

/*DeleteLinkOK The delete request was successful and the resource has been removed

swagger:response deleteLinkOK
*/
type DeleteLinkOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.Empty `json:"body,omitempty"`
}

// NewDeleteLinkOK creates DeleteLinkOK with default headers values
func NewDeleteLinkOK() *DeleteLinkOK {

	return &DeleteLinkOK{}
}

// WithPayload adds the payload to the delete link o k response
func (o *DeleteLinkOK) WithPayload(payload *rest_model.Empty) *DeleteLinkOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete link o k response
func (o *DeleteLinkOK) SetPayload(payload *rest_model.Empty) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteLinkOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteLinkBadRequestCode is the HTTP code returned for type DeleteLinkBadRequest
const DeleteLinkBadRequestCode int = 400

/*DeleteLinkBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response deleteLinkBadRequest
*/
type DeleteLinkBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDeleteLinkBadRequest creates DeleteLinkBadRequest with default headers values
func NewDeleteLinkBadRequest() *DeleteLinkBadRequest {

	return &DeleteLinkBadRequest{}
}

// WithPayload adds the payload to the delete link bad request response
func (o *DeleteLinkBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *DeleteLinkBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete link bad request response
func (o *DeleteLinkBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteLinkBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteLinkUnauthorizedCode is the HTTP code returned for type DeleteLinkUnauthorized
const DeleteLinkUnauthorizedCode int = 401

/*DeleteLinkUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response deleteLinkUnauthorized
*/
type DeleteLinkUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDeleteLinkUnauthorized creates DeleteLinkUnauthorized with default headers values
func NewDeleteLinkUnauthorized() *DeleteLinkUnauthorized {

	return &DeleteLinkUnauthorized{}
}

// WithPayload adds the payload to the delete link unauthorized response
func (o *DeleteLinkUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *DeleteLinkUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete link unauthorized response
func (o *DeleteLinkUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteLinkUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteLinkTooManyRequestsCode is the HTTP code returned for type DeleteLinkTooManyRequests
const DeleteLinkTooManyRequestsCode int = 429

/*DeleteLinkTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response deleteLinkTooManyRequests
*/
type DeleteLinkTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDeleteLinkTooManyRequests creates DeleteLinkTooManyRequests with default headers values
func NewDeleteLinkTooManyRequests() *DeleteLinkTooManyRequests {

	return &DeleteLinkTooManyRequests{}
}

// WithPayload adds the payload to the delete link too many requests response
func (o *DeleteLinkTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *DeleteLinkTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete link too many requests response
func (o *DeleteLinkTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteLinkTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
