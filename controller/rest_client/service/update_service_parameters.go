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

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/cosmic-cloak/ztna/controller/rest_model"
)

// NewUpdateServiceParams creates a new UpdateServiceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateServiceParams() *UpdateServiceParams {
	return &UpdateServiceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateServiceParamsWithTimeout creates a new UpdateServiceParams object
// with the ability to set a timeout on a request.
func NewUpdateServiceParamsWithTimeout(timeout time.Duration) *UpdateServiceParams {
	return &UpdateServiceParams{
		timeout: timeout,
	}
}

// NewUpdateServiceParamsWithContext creates a new UpdateServiceParams object
// with the ability to set a context for a request.
func NewUpdateServiceParamsWithContext(ctx context.Context) *UpdateServiceParams {
	return &UpdateServiceParams{
		Context: ctx,
	}
}

// NewUpdateServiceParamsWithHTTPClient creates a new UpdateServiceParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateServiceParamsWithHTTPClient(client *http.Client) *UpdateServiceParams {
	return &UpdateServiceParams{
		HTTPClient: client,
	}
}

/* UpdateServiceParams contains all the parameters to send to the API endpoint
   for the update service operation.

   Typically these are written to a http.Request.
*/
type UpdateServiceParams struct {

	/* ID.

	   The id of the requested resource
	*/
	ID string

	/* Service.

	   A service update object
	*/
	Service *rest_model.ServiceUpdate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update service params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateServiceParams) WithDefaults() *UpdateServiceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update service params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateServiceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update service params
func (o *UpdateServiceParams) WithTimeout(timeout time.Duration) *UpdateServiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update service params
func (o *UpdateServiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update service params
func (o *UpdateServiceParams) WithContext(ctx context.Context) *UpdateServiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update service params
func (o *UpdateServiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update service params
func (o *UpdateServiceParams) WithHTTPClient(client *http.Client) *UpdateServiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update service params
func (o *UpdateServiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the update service params
func (o *UpdateServiceParams) WithID(id string) *UpdateServiceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update service params
func (o *UpdateServiceParams) SetID(id string) {
	o.ID = id
}

// WithService adds the service to the update service params
func (o *UpdateServiceParams) WithService(service *rest_model.ServiceUpdate) *UpdateServiceParams {
	o.SetService(service)
	return o
}

// SetService adds the service to the update service params
func (o *UpdateServiceParams) SetService(service *rest_model.ServiceUpdate) {
	o.Service = service
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateServiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}
	if o.Service != nil {
		if err := r.SetBodyParam(o.Service); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
