// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"redpanda-schema-registry/models"
)

// PutConfigReader is a Reader for the PutConfig structure.
type PutConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewPutConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /config] put_config", response, response.Code())
	}
}

// NewPutConfigOK creates a PutConfigOK with default headers values
func NewPutConfigOK() *PutConfigOK {
	return &PutConfigOK{}
}

/*
PutConfigOK describes a response with status code 200, with default header values.

OK
*/
type PutConfigOK struct {
	Payload *PutConfigOKBody
}

// IsSuccess returns true when this put config o k response has a 2xx status code
func (o *PutConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put config o k response has a 3xx status code
func (o *PutConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put config o k response has a 4xx status code
func (o *PutConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put config o k response has a 5xx status code
func (o *PutConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put config o k response a status code equal to that given
func (o *PutConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put config o k response
func (o *PutConfigOK) Code() int {
	return 200
}

func (o *PutConfigOK) Error() string {
	return fmt.Sprintf("[PUT /config][%d] putConfigOK  %+v", 200, o.Payload)
}

func (o *PutConfigOK) String() string {
	return fmt.Sprintf("[PUT /config][%d] putConfigOK  %+v", 200, o.Payload)
}

func (o *PutConfigOK) GetPayload() *PutConfigOKBody {
	return o.Payload
}

func (o *PutConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PutConfigOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConfigInternalServerError creates a PutConfigInternalServerError with default headers values
func NewPutConfigInternalServerError() *PutConfigInternalServerError {
	return &PutConfigInternalServerError{}
}

/*
PutConfigInternalServerError describes a response with status code 500, with default header values.

Internal Server error
*/
type PutConfigInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put config internal server error response has a 2xx status code
func (o *PutConfigInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put config internal server error response has a 3xx status code
func (o *PutConfigInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put config internal server error response has a 4xx status code
func (o *PutConfigInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put config internal server error response has a 5xx status code
func (o *PutConfigInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put config internal server error response a status code equal to that given
func (o *PutConfigInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the put config internal server error response
func (o *PutConfigInternalServerError) Code() int {
	return 500
}

func (o *PutConfigInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /config][%d] putConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *PutConfigInternalServerError) String() string {
	return fmt.Sprintf("[PUT /config][%d] putConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *PutConfigInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PutConfigBody put config body
swagger:model PutConfigBody
*/
type PutConfigBody struct {

	// compatibility
	Compatibility string `json:"compatibility,omitempty"`
}

// Validate validates this put config body
func (o *PutConfigBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this put config body based on context it is used
func (o *PutConfigBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutConfigBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutConfigBody) UnmarshalBinary(b []byte) error {
	var res PutConfigBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PutConfigOKBody put config o k body
swagger:model PutConfigOKBody
*/
type PutConfigOKBody struct {

	// compatibility
	Compatibility string `json:"compatibility,omitempty"`
}

// Validate validates this put config o k body
func (o *PutConfigOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this put config o k body based on context it is used
func (o *PutConfigOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutConfigOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutConfigOKBody) UnmarshalBinary(b []byte) error {
	var res PutConfigOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}