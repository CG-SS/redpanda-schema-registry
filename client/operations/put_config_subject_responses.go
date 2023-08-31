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

// PutConfigSubjectReader is a Reader for the PutConfigSubject structure.
type PutConfigSubjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutConfigSubjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutConfigSubjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewPutConfigSubjectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /config/{subject}] put_config_subject", response, response.Code())
	}
}

// NewPutConfigSubjectOK creates a PutConfigSubjectOK with default headers values
func NewPutConfigSubjectOK() *PutConfigSubjectOK {
	return &PutConfigSubjectOK{}
}

/*
PutConfigSubjectOK describes a response with status code 200, with default header values.

OK
*/
type PutConfigSubjectOK struct {
	Payload *PutConfigSubjectOKBody
}

// IsSuccess returns true when this put config subject o k response has a 2xx status code
func (o *PutConfigSubjectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put config subject o k response has a 3xx status code
func (o *PutConfigSubjectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put config subject o k response has a 4xx status code
func (o *PutConfigSubjectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put config subject o k response has a 5xx status code
func (o *PutConfigSubjectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put config subject o k response a status code equal to that given
func (o *PutConfigSubjectOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put config subject o k response
func (o *PutConfigSubjectOK) Code() int {
	return 200
}

func (o *PutConfigSubjectOK) Error() string {
	return fmt.Sprintf("[PUT /config/{subject}][%d] putConfigSubjectOK  %+v", 200, o.Payload)
}

func (o *PutConfigSubjectOK) String() string {
	return fmt.Sprintf("[PUT /config/{subject}][%d] putConfigSubjectOK  %+v", 200, o.Payload)
}

func (o *PutConfigSubjectOK) GetPayload() *PutConfigSubjectOKBody {
	return o.Payload
}

func (o *PutConfigSubjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PutConfigSubjectOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConfigSubjectInternalServerError creates a PutConfigSubjectInternalServerError with default headers values
func NewPutConfigSubjectInternalServerError() *PutConfigSubjectInternalServerError {
	return &PutConfigSubjectInternalServerError{}
}

/*
PutConfigSubjectInternalServerError describes a response with status code 500, with default header values.

Internal Server error
*/
type PutConfigSubjectInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put config subject internal server error response has a 2xx status code
func (o *PutConfigSubjectInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put config subject internal server error response has a 3xx status code
func (o *PutConfigSubjectInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put config subject internal server error response has a 4xx status code
func (o *PutConfigSubjectInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put config subject internal server error response has a 5xx status code
func (o *PutConfigSubjectInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put config subject internal server error response a status code equal to that given
func (o *PutConfigSubjectInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the put config subject internal server error response
func (o *PutConfigSubjectInternalServerError) Code() int {
	return 500
}

func (o *PutConfigSubjectInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /config/{subject}][%d] putConfigSubjectInternalServerError  %+v", 500, o.Payload)
}

func (o *PutConfigSubjectInternalServerError) String() string {
	return fmt.Sprintf("[PUT /config/{subject}][%d] putConfigSubjectInternalServerError  %+v", 500, o.Payload)
}

func (o *PutConfigSubjectInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConfigSubjectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PutConfigSubjectBody put config subject body
swagger:model PutConfigSubjectBody
*/
type PutConfigSubjectBody struct {

	// compatibility
	Compatibility string `json:"compatibility,omitempty"`
}

// Validate validates this put config subject body
func (o *PutConfigSubjectBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this put config subject body based on context it is used
func (o *PutConfigSubjectBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutConfigSubjectBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutConfigSubjectBody) UnmarshalBinary(b []byte) error {
	var res PutConfigSubjectBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PutConfigSubjectOKBody put config subject o k body
swagger:model PutConfigSubjectOKBody
*/
type PutConfigSubjectOKBody struct {

	// compatibility
	Compatibility string `json:"compatibility,omitempty"`
}

// Validate validates this put config subject o k body
func (o *PutConfigSubjectOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this put config subject o k body based on context it is used
func (o *PutConfigSubjectOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutConfigSubjectOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutConfigSubjectOKBody) UnmarshalBinary(b []byte) error {
	var res PutConfigSubjectOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
