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

// NewGetModeParams creates a new GetModeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetModeParams() *GetModeParams {
	return &GetModeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetModeParamsWithTimeout creates a new GetModeParams object
// with the ability to set a timeout on a request.
func NewGetModeParamsWithTimeout(timeout time.Duration) *GetModeParams {
	return &GetModeParams{
		timeout: timeout,
	}
}

// NewGetModeParamsWithContext creates a new GetModeParams object
// with the ability to set a context for a request.
func NewGetModeParamsWithContext(ctx context.Context) *GetModeParams {
	return &GetModeParams{
		Context: ctx,
	}
}

// NewGetModeParamsWithHTTPClient creates a new GetModeParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetModeParamsWithHTTPClient(client *http.Client) *GetModeParams {
	return &GetModeParams{
		HTTPClient: client,
	}
}

/*
GetModeParams contains all the parameters to send to the API endpoint

	for the get mode operation.

	Typically these are written to a http.Request.
*/
type GetModeParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get mode params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetModeParams) WithDefaults() *GetModeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get mode params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetModeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get mode params
func (o *GetModeParams) WithTimeout(timeout time.Duration) *GetModeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get mode params
func (o *GetModeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get mode params
func (o *GetModeParams) WithContext(ctx context.Context) *GetModeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get mode params
func (o *GetModeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get mode params
func (o *GetModeParams) WithHTTPClient(client *http.Client) *GetModeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get mode params
func (o *GetModeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetModeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}