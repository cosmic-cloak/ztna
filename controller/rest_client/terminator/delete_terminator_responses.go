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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cosmic-cloak/ztna/controller/rest_model"
)

// DeleteTerminatorReader is a Reader for the DeleteTerminator structure.
type DeleteTerminatorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTerminatorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteTerminatorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteTerminatorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteTerminatorUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteTerminatorConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteTerminatorTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteTerminatorOK creates a DeleteTerminatorOK with default headers values
func NewDeleteTerminatorOK() *DeleteTerminatorOK {
	return &DeleteTerminatorOK{}
}

/* DeleteTerminatorOK describes a response with status code 200, with default header values.

The delete request was successful and the resource has been removed
*/
type DeleteTerminatorOK struct {
	Payload *rest_model.Empty
}

func (o *DeleteTerminatorOK) Error() string {
	return fmt.Sprintf("[DELETE /terminators/{id}][%d] deleteTerminatorOK  %+v", 200, o.Payload)
}
func (o *DeleteTerminatorOK) GetPayload() *rest_model.Empty {
	return o.Payload
}

func (o *DeleteTerminatorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.Empty)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTerminatorBadRequest creates a DeleteTerminatorBadRequest with default headers values
func NewDeleteTerminatorBadRequest() *DeleteTerminatorBadRequest {
	return &DeleteTerminatorBadRequest{}
}

/* DeleteTerminatorBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type DeleteTerminatorBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DeleteTerminatorBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /terminators/{id}][%d] deleteTerminatorBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteTerminatorBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteTerminatorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTerminatorUnauthorized creates a DeleteTerminatorUnauthorized with default headers values
func NewDeleteTerminatorUnauthorized() *DeleteTerminatorUnauthorized {
	return &DeleteTerminatorUnauthorized{}
}

/* DeleteTerminatorUnauthorized describes a response with status code 401, with default header values.

The currently supplied session does not have the correct access rights to request this resource
*/
type DeleteTerminatorUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DeleteTerminatorUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /terminators/{id}][%d] deleteTerminatorUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteTerminatorUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteTerminatorUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTerminatorConflict creates a DeleteTerminatorConflict with default headers values
func NewDeleteTerminatorConflict() *DeleteTerminatorConflict {
	return &DeleteTerminatorConflict{}
}

/* DeleteTerminatorConflict describes a response with status code 409, with default header values.

The resource requested to be removed/altered cannot be as it is referenced by another object.
*/
type DeleteTerminatorConflict struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DeleteTerminatorConflict) Error() string {
	return fmt.Sprintf("[DELETE /terminators/{id}][%d] deleteTerminatorConflict  %+v", 409, o.Payload)
}
func (o *DeleteTerminatorConflict) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteTerminatorConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTerminatorTooManyRequests creates a DeleteTerminatorTooManyRequests with default headers values
func NewDeleteTerminatorTooManyRequests() *DeleteTerminatorTooManyRequests {
	return &DeleteTerminatorTooManyRequests{}
}

/* DeleteTerminatorTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type DeleteTerminatorTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DeleteTerminatorTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /terminators/{id}][%d] deleteTerminatorTooManyRequests  %+v", 429, o.Payload)
}
func (o *DeleteTerminatorTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteTerminatorTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
