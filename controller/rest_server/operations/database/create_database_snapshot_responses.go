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

package database

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"ztna-core/ztna/controller/rest_model"
)

// CreateDatabaseSnapshotOKCode is the HTTP code returned for type CreateDatabaseSnapshotOK
const CreateDatabaseSnapshotOKCode int = 200

/*CreateDatabaseSnapshotOK Base empty response

swagger:response createDatabaseSnapshotOK
*/
type CreateDatabaseSnapshotOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.Empty `json:"body,omitempty"`
}

// NewCreateDatabaseSnapshotOK creates CreateDatabaseSnapshotOK with default headers values
func NewCreateDatabaseSnapshotOK() *CreateDatabaseSnapshotOK {

	return &CreateDatabaseSnapshotOK{}
}

// WithPayload adds the payload to the create database snapshot o k response
func (o *CreateDatabaseSnapshotOK) WithPayload(payload *rest_model.Empty) *CreateDatabaseSnapshotOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create database snapshot o k response
func (o *CreateDatabaseSnapshotOK) SetPayload(payload *rest_model.Empty) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateDatabaseSnapshotOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateDatabaseSnapshotUnauthorizedCode is the HTTP code returned for type CreateDatabaseSnapshotUnauthorized
const CreateDatabaseSnapshotUnauthorizedCode int = 401

/*CreateDatabaseSnapshotUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response createDatabaseSnapshotUnauthorized
*/
type CreateDatabaseSnapshotUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewCreateDatabaseSnapshotUnauthorized creates CreateDatabaseSnapshotUnauthorized with default headers values
func NewCreateDatabaseSnapshotUnauthorized() *CreateDatabaseSnapshotUnauthorized {

	return &CreateDatabaseSnapshotUnauthorized{}
}

// WithPayload adds the payload to the create database snapshot unauthorized response
func (o *CreateDatabaseSnapshotUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *CreateDatabaseSnapshotUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create database snapshot unauthorized response
func (o *CreateDatabaseSnapshotUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateDatabaseSnapshotUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateDatabaseSnapshotTooManyRequestsCode is the HTTP code returned for type CreateDatabaseSnapshotTooManyRequests
const CreateDatabaseSnapshotTooManyRequestsCode int = 429

/*CreateDatabaseSnapshotTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response createDatabaseSnapshotTooManyRequests
*/
type CreateDatabaseSnapshotTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewCreateDatabaseSnapshotTooManyRequests creates CreateDatabaseSnapshotTooManyRequests with default headers values
func NewCreateDatabaseSnapshotTooManyRequests() *CreateDatabaseSnapshotTooManyRequests {

	return &CreateDatabaseSnapshotTooManyRequests{}
}

// WithPayload adds the payload to the create database snapshot too many requests response
func (o *CreateDatabaseSnapshotTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *CreateDatabaseSnapshotTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create database snapshot too many requests response
func (o *CreateDatabaseSnapshotTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateDatabaseSnapshotTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
