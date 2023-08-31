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

// NewGetSchemasTypesParams creates a new GetSchemasTypesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSchemasTypesParams() *GetSchemasTypesParams {
	return &GetSchemasTypesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSchemasTypesParamsWithTimeout creates a new GetSchemasTypesParams object
// with the ability to set a timeout on a request.
func NewGetSchemasTypesParamsWithTimeout(timeout time.Duration) *GetSchemasTypesParams {
	return &GetSchemasTypesParams{
		timeout: timeout,
	}
}

// NewGetSchemasTypesParamsWithContext creates a new GetSchemasTypesParams object
// with the ability to set a context for a request.
func NewGetSchemasTypesParamsWithContext(ctx context.Context) *GetSchemasTypesParams {
	return &GetSchemasTypesParams{
		Context: ctx,
	}
}

// NewGetSchemasTypesParamsWithHTTPClient creates a new GetSchemasTypesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSchemasTypesParamsWithHTTPClient(client *http.Client) *GetSchemasTypesParams {
	return &GetSchemasTypesParams{
		HTTPClient: client,
	}
}

/*
GetSchemasTypesParams contains all the parameters to send to the API endpoint

	for the get schemas types operation.

	Typically these are written to a http.Request.
*/
type GetSchemasTypesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get schemas types params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSchemasTypesParams) WithDefaults() *GetSchemasTypesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get schemas types params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSchemasTypesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get schemas types params
func (o *GetSchemasTypesParams) WithTimeout(timeout time.Duration) *GetSchemasTypesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get schemas types params
func (o *GetSchemasTypesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get schemas types params
func (o *GetSchemasTypesParams) WithContext(ctx context.Context) *GetSchemasTypesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get schemas types params
func (o *GetSchemasTypesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get schemas types params
func (o *GetSchemasTypesParams) WithHTTPClient(client *http.Client) *GetSchemasTypesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get schemas types params
func (o *GetSchemasTypesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetSchemasTypesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
