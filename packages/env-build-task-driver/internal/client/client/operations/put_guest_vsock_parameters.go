// Code generated by go-swagger; DO NOT EDIT.

package operations

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

	"github.com/e2b-dev/api/packages/env-build-task-driver/internal/client/models"
)

// NewPutGuestVsockParams creates a new PutGuestVsockParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutGuestVsockParams() *PutGuestVsockParams {
	return &PutGuestVsockParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutGuestVsockParamsWithTimeout creates a new PutGuestVsockParams object
// with the ability to set a timeout on a request.
func NewPutGuestVsockParamsWithTimeout(timeout time.Duration) *PutGuestVsockParams {
	return &PutGuestVsockParams{
		timeout: timeout,
	}
}

// NewPutGuestVsockParamsWithContext creates a new PutGuestVsockParams object
// with the ability to set a context for a request.
func NewPutGuestVsockParamsWithContext(ctx context.Context) *PutGuestVsockParams {
	return &PutGuestVsockParams{
		Context: ctx,
	}
}

// NewPutGuestVsockParamsWithHTTPClient creates a new PutGuestVsockParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutGuestVsockParamsWithHTTPClient(client *http.Client) *PutGuestVsockParams {
	return &PutGuestVsockParams{
		HTTPClient: client,
	}
}

/*
PutGuestVsockParams contains all the parameters to send to the API endpoint

	for the put guest vsock operation.

	Typically these are written to a http.Request.
*/
type PutGuestVsockParams struct {

	/* Body.

	   Guest vsock properties
	*/
	Body *models.Vsock

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put guest vsock params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutGuestVsockParams) WithDefaults() *PutGuestVsockParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put guest vsock params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutGuestVsockParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put guest vsock params
func (o *PutGuestVsockParams) WithTimeout(timeout time.Duration) *PutGuestVsockParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put guest vsock params
func (o *PutGuestVsockParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put guest vsock params
func (o *PutGuestVsockParams) WithContext(ctx context.Context) *PutGuestVsockParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put guest vsock params
func (o *PutGuestVsockParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put guest vsock params
func (o *PutGuestVsockParams) WithHTTPClient(client *http.Client) *PutGuestVsockParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put guest vsock params
func (o *PutGuestVsockParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put guest vsock params
func (o *PutGuestVsockParams) WithBody(body *models.Vsock) *PutGuestVsockParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put guest vsock params
func (o *PutGuestVsockParams) SetBody(body *models.Vsock) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PutGuestVsockParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
