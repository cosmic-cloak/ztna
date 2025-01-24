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
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/cosmic-cloak/ztna/controller/rest_model"
)

// NewCreateDatabaseSnapshotWithPathParams creates a new CreateDatabaseSnapshotWithPathParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateDatabaseSnapshotWithPathParams() *CreateDatabaseSnapshotWithPathParams {
	return &CreateDatabaseSnapshotWithPathParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateDatabaseSnapshotWithPathParamsWithTimeout creates a new CreateDatabaseSnapshotWithPathParams object
// with the ability to set a timeout on a request.
func NewCreateDatabaseSnapshotWithPathParamsWithTimeout(timeout time.Duration) *CreateDatabaseSnapshotWithPathParams {
	return &CreateDatabaseSnapshotWithPathParams{
		timeout: timeout,
	}
}

// NewCreateDatabaseSnapshotWithPathParamsWithContext creates a new CreateDatabaseSnapshotWithPathParams object
// with the ability to set a context for a request.
func NewCreateDatabaseSnapshotWithPathParamsWithContext(ctx context.Context) *CreateDatabaseSnapshotWithPathParams {
	return &CreateDatabaseSnapshotWithPathParams{
		Context: ctx,
	}
}

// NewCreateDatabaseSnapshotWithPathParamsWithHTTPClient creates a new CreateDatabaseSnapshotWithPathParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateDatabaseSnapshotWithPathParamsWithHTTPClient(client *http.Client) *CreateDatabaseSnapshotWithPathParams {
	return &CreateDatabaseSnapshotWithPathParams{
		HTTPClient: client,
	}
}

/* CreateDatabaseSnapshotWithPathParams contains all the parameters to send to the API endpoint
   for the create database snapshot with path operation.

   Typically these are written to a http.Request.
*/
type CreateDatabaseSnapshotWithPathParams struct {

	/* Snapshot.

	   snapshot parameters
	*/
	Snapshot *rest_model.DatabaseSnapshotCreate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create database snapshot with path params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateDatabaseSnapshotWithPathParams) WithDefaults() *CreateDatabaseSnapshotWithPathParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create database snapshot with path params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateDatabaseSnapshotWithPathParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create database snapshot with path params
func (o *CreateDatabaseSnapshotWithPathParams) WithTimeout(timeout time.Duration) *CreateDatabaseSnapshotWithPathParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create database snapshot with path params
func (o *CreateDatabaseSnapshotWithPathParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create database snapshot with path params
func (o *CreateDatabaseSnapshotWithPathParams) WithContext(ctx context.Context) *CreateDatabaseSnapshotWithPathParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create database snapshot with path params
func (o *CreateDatabaseSnapshotWithPathParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create database snapshot with path params
func (o *CreateDatabaseSnapshotWithPathParams) WithHTTPClient(client *http.Client) *CreateDatabaseSnapshotWithPathParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create database snapshot with path params
func (o *CreateDatabaseSnapshotWithPathParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSnapshot adds the snapshot to the create database snapshot with path params
func (o *CreateDatabaseSnapshotWithPathParams) WithSnapshot(snapshot *rest_model.DatabaseSnapshotCreate) *CreateDatabaseSnapshotWithPathParams {
	o.SetSnapshot(snapshot)
	return o
}

// SetSnapshot adds the snapshot to the create database snapshot with path params
func (o *CreateDatabaseSnapshotWithPathParams) SetSnapshot(snapshot *rest_model.DatabaseSnapshotCreate) {
	o.Snapshot = snapshot
}

// WriteToRequest writes these params to a swagger request
func (o *CreateDatabaseSnapshotWithPathParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Snapshot != nil {
		if err := r.SetBodyParam(o.Snapshot); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
