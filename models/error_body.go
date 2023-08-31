// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ErrorBody error body
//
// swagger:model error_body
type ErrorBody struct {

	// error code
	ErrorCode int64 `json:"error_code,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this error body
func (m *ErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this error body based on context it is used
func (m *ErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ErrorBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ErrorBody) UnmarshalBinary(b []byte) error {
	var res ErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}