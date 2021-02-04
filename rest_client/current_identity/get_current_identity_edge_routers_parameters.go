// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright NetFoundry, Inc.
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

package current_identity

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
)

// NewGetCurrentIdentityEdgeRoutersParams creates a new GetCurrentIdentityEdgeRoutersParams object
// with the default values initialized.
func NewGetCurrentIdentityEdgeRoutersParams() *GetCurrentIdentityEdgeRoutersParams {

	return &GetCurrentIdentityEdgeRoutersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCurrentIdentityEdgeRoutersParamsWithTimeout creates a new GetCurrentIdentityEdgeRoutersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCurrentIdentityEdgeRoutersParamsWithTimeout(timeout time.Duration) *GetCurrentIdentityEdgeRoutersParams {

	return &GetCurrentIdentityEdgeRoutersParams{

		timeout: timeout,
	}
}

// NewGetCurrentIdentityEdgeRoutersParamsWithContext creates a new GetCurrentIdentityEdgeRoutersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCurrentIdentityEdgeRoutersParamsWithContext(ctx context.Context) *GetCurrentIdentityEdgeRoutersParams {

	return &GetCurrentIdentityEdgeRoutersParams{

		Context: ctx,
	}
}

// NewGetCurrentIdentityEdgeRoutersParamsWithHTTPClient creates a new GetCurrentIdentityEdgeRoutersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCurrentIdentityEdgeRoutersParamsWithHTTPClient(client *http.Client) *GetCurrentIdentityEdgeRoutersParams {

	return &GetCurrentIdentityEdgeRoutersParams{
		HTTPClient: client,
	}
}

/*GetCurrentIdentityEdgeRoutersParams contains all the parameters to send to the API endpoint
for the get current identity edge routers operation typically these are written to a http.Request
*/
type GetCurrentIdentityEdgeRoutersParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get current identity edge routers params
func (o *GetCurrentIdentityEdgeRoutersParams) WithTimeout(timeout time.Duration) *GetCurrentIdentityEdgeRoutersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get current identity edge routers params
func (o *GetCurrentIdentityEdgeRoutersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get current identity edge routers params
func (o *GetCurrentIdentityEdgeRoutersParams) WithContext(ctx context.Context) *GetCurrentIdentityEdgeRoutersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get current identity edge routers params
func (o *GetCurrentIdentityEdgeRoutersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get current identity edge routers params
func (o *GetCurrentIdentityEdgeRoutersParams) WithHTTPClient(client *http.Client) *GetCurrentIdentityEdgeRoutersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get current identity edge routers params
func (o *GetCurrentIdentityEdgeRoutersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetCurrentIdentityEdgeRoutersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
