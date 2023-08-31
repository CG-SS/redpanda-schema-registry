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
)

// NewPutConfigSubjectParams creates a new PutConfigSubjectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutConfigSubjectParams() *PutConfigSubjectParams {
	return &PutConfigSubjectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutConfigSubjectParamsWithTimeout creates a new PutConfigSubjectParams object
// with the ability to set a timeout on a request.
func NewPutConfigSubjectParamsWithTimeout(timeout time.Duration) *PutConfigSubjectParams {
	return &PutConfigSubjectParams{
		timeout: timeout,
	}
}

// NewPutConfigSubjectParamsWithContext creates a new PutConfigSubjectParams object
// with the ability to set a context for a request.
func NewPutConfigSubjectParamsWithContext(ctx context.Context) *PutConfigSubjectParams {
	return &PutConfigSubjectParams{
		Context: ctx,
	}
}

// NewPutConfigSubjectParamsWithHTTPClient creates a new PutConfigSubjectParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutConfigSubjectParamsWithHTTPClient(client *http.Client) *PutConfigSubjectParams {
	return &PutConfigSubjectParams{
		HTTPClient: client,
	}
}

/*
PutConfigSubjectParams contains all the parameters to send to the API endpoint

	for the put config subject operation.

	Typically these are written to a http.Request.
*/
type PutConfigSubjectParams struct {

	// Config.
	Config PutConfigSubjectBody

	// Subject.
	Subject string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put config subject params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutConfigSubjectParams) WithDefaults() *PutConfigSubjectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put config subject params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutConfigSubjectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put config subject params
func (o *PutConfigSubjectParams) WithTimeout(timeout time.Duration) *PutConfigSubjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put config subject params
func (o *PutConfigSubjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put config subject params
func (o *PutConfigSubjectParams) WithContext(ctx context.Context) *PutConfigSubjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put config subject params
func (o *PutConfigSubjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put config subject params
func (o *PutConfigSubjectParams) WithHTTPClient(client *http.Client) *PutConfigSubjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put config subject params
func (o *PutConfigSubjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfig adds the config to the put config subject params
func (o *PutConfigSubjectParams) WithConfig(config PutConfigSubjectBody) *PutConfigSubjectParams {
	o.SetConfig(config)
	return o
}

// SetConfig adds the config to the put config subject params
func (o *PutConfigSubjectParams) SetConfig(config PutConfigSubjectBody) {
	o.Config = config
}

// WithSubject adds the subject to the put config subject params
func (o *PutConfigSubjectParams) WithSubject(subject string) *PutConfigSubjectParams {
	o.SetSubject(subject)
	return o
}

// SetSubject adds the subject to the put config subject params
func (o *PutConfigSubjectParams) SetSubject(subject string) {
	o.Subject = subject
}

// WriteToRequest writes these params to a swagger request
func (o *PutConfigSubjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Config); err != nil {
		return err
	}

	// path param subject
	if err := r.SetPathParam("subject", o.Subject); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
