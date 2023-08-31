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
	"github.com/go-openapi/swag"
)

// NewDeleteSubjectVersionParams creates a new DeleteSubjectVersionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteSubjectVersionParams() *DeleteSubjectVersionParams {
	return &DeleteSubjectVersionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSubjectVersionParamsWithTimeout creates a new DeleteSubjectVersionParams object
// with the ability to set a timeout on a request.
func NewDeleteSubjectVersionParamsWithTimeout(timeout time.Duration) *DeleteSubjectVersionParams {
	return &DeleteSubjectVersionParams{
		timeout: timeout,
	}
}

// NewDeleteSubjectVersionParamsWithContext creates a new DeleteSubjectVersionParams object
// with the ability to set a context for a request.
func NewDeleteSubjectVersionParamsWithContext(ctx context.Context) *DeleteSubjectVersionParams {
	return &DeleteSubjectVersionParams{
		Context: ctx,
	}
}

// NewDeleteSubjectVersionParamsWithHTTPClient creates a new DeleteSubjectVersionParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteSubjectVersionParamsWithHTTPClient(client *http.Client) *DeleteSubjectVersionParams {
	return &DeleteSubjectVersionParams{
		HTTPClient: client,
	}
}

/*
DeleteSubjectVersionParams contains all the parameters to send to the API endpoint

	for the delete subject version operation.

	Typically these are written to a http.Request.
*/
type DeleteSubjectVersionParams struct {

	// Permanent.
	Permanent *bool

	// Subject.
	Subject string

	// Version.
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete subject version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSubjectVersionParams) WithDefaults() *DeleteSubjectVersionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete subject version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSubjectVersionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete subject version params
func (o *DeleteSubjectVersionParams) WithTimeout(timeout time.Duration) *DeleteSubjectVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete subject version params
func (o *DeleteSubjectVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete subject version params
func (o *DeleteSubjectVersionParams) WithContext(ctx context.Context) *DeleteSubjectVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete subject version params
func (o *DeleteSubjectVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete subject version params
func (o *DeleteSubjectVersionParams) WithHTTPClient(client *http.Client) *DeleteSubjectVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete subject version params
func (o *DeleteSubjectVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPermanent adds the permanent to the delete subject version params
func (o *DeleteSubjectVersionParams) WithPermanent(permanent *bool) *DeleteSubjectVersionParams {
	o.SetPermanent(permanent)
	return o
}

// SetPermanent adds the permanent to the delete subject version params
func (o *DeleteSubjectVersionParams) SetPermanent(permanent *bool) {
	o.Permanent = permanent
}

// WithSubject adds the subject to the delete subject version params
func (o *DeleteSubjectVersionParams) WithSubject(subject string) *DeleteSubjectVersionParams {
	o.SetSubject(subject)
	return o
}

// SetSubject adds the subject to the delete subject version params
func (o *DeleteSubjectVersionParams) SetSubject(subject string) {
	o.Subject = subject
}

// WithVersion adds the version to the delete subject version params
func (o *DeleteSubjectVersionParams) WithVersion(version string) *DeleteSubjectVersionParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete subject version params
func (o *DeleteSubjectVersionParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSubjectVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Permanent != nil {

		// query param permanent
		var qrPermanent bool

		if o.Permanent != nil {
			qrPermanent = *o.Permanent
		}
		qPermanent := swag.FormatBool(qrPermanent)
		if qPermanent != "" {

			if err := r.SetQueryParam("permanent", qPermanent); err != nil {
				return err
			}
		}
	}

	// path param subject
	if err := r.SetPathParam("subject", o.Subject); err != nil {
		return err
	}

	// path param version
	if err := r.SetPathParam("version", o.Version); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}