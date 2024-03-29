// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/CG-SS/redpanda-schema-registry/models"
)

// GetSubjectVersionsReader is a Reader for the GetSubjectVersions structure.
type GetSubjectVersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSubjectVersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSubjectVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetSubjectVersionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSubjectVersionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /subjects/{subject}/versions] get_subject_versions", response, response.Code())
	}
}

// NewGetSubjectVersionsOK creates a GetSubjectVersionsOK with default headers values
func NewGetSubjectVersionsOK() *GetSubjectVersionsOK {
	return &GetSubjectVersionsOK{}
}

/*
GetSubjectVersionsOK describes a response with status code 200, with default header values.

OK
*/
type GetSubjectVersionsOK struct {
	Payload []int64
}

// IsSuccess returns true when this get subject versions o k response has a 2xx status code
func (o *GetSubjectVersionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get subject versions o k response has a 3xx status code
func (o *GetSubjectVersionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get subject versions o k response has a 4xx status code
func (o *GetSubjectVersionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get subject versions o k response has a 5xx status code
func (o *GetSubjectVersionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get subject versions o k response a status code equal to that given
func (o *GetSubjectVersionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get subject versions o k response
func (o *GetSubjectVersionsOK) Code() int {
	return 200
}

func (o *GetSubjectVersionsOK) Error() string {
	return fmt.Sprintf("[GET /subjects/{subject}/versions][%d] getSubjectVersionsOK  %+v", 200, o.Payload)
}

func (o *GetSubjectVersionsOK) String() string {
	return fmt.Sprintf("[GET /subjects/{subject}/versions][%d] getSubjectVersionsOK  %+v", 200, o.Payload)
}

func (o *GetSubjectVersionsOK) GetPayload() []int64 {
	return o.Payload
}

func (o *GetSubjectVersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubjectVersionsNotFound creates a GetSubjectVersionsNotFound with default headers values
func NewGetSubjectVersionsNotFound() *GetSubjectVersionsNotFound {
	return &GetSubjectVersionsNotFound{}
}

/*
GetSubjectVersionsNotFound describes a response with status code 404, with default header values.

Subject not found
*/
type GetSubjectVersionsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get subject versions not found response has a 2xx status code
func (o *GetSubjectVersionsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get subject versions not found response has a 3xx status code
func (o *GetSubjectVersionsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get subject versions not found response has a 4xx status code
func (o *GetSubjectVersionsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get subject versions not found response has a 5xx status code
func (o *GetSubjectVersionsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get subject versions not found response a status code equal to that given
func (o *GetSubjectVersionsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get subject versions not found response
func (o *GetSubjectVersionsNotFound) Code() int {
	return 404
}

func (o *GetSubjectVersionsNotFound) Error() string {
	return fmt.Sprintf("[GET /subjects/{subject}/versions][%d] getSubjectVersionsNotFound  %+v", 404, o.Payload)
}

func (o *GetSubjectVersionsNotFound) String() string {
	return fmt.Sprintf("[GET /subjects/{subject}/versions][%d] getSubjectVersionsNotFound  %+v", 404, o.Payload)
}

func (o *GetSubjectVersionsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSubjectVersionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubjectVersionsInternalServerError creates a GetSubjectVersionsInternalServerError with default headers values
func NewGetSubjectVersionsInternalServerError() *GetSubjectVersionsInternalServerError {
	return &GetSubjectVersionsInternalServerError{}
}

/*
GetSubjectVersionsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetSubjectVersionsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get subject versions internal server error response has a 2xx status code
func (o *GetSubjectVersionsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get subject versions internal server error response has a 3xx status code
func (o *GetSubjectVersionsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get subject versions internal server error response has a 4xx status code
func (o *GetSubjectVersionsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get subject versions internal server error response has a 5xx status code
func (o *GetSubjectVersionsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get subject versions internal server error response a status code equal to that given
func (o *GetSubjectVersionsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get subject versions internal server error response
func (o *GetSubjectVersionsInternalServerError) Code() int {
	return 500
}

func (o *GetSubjectVersionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /subjects/{subject}/versions][%d] getSubjectVersionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSubjectVersionsInternalServerError) String() string {
	return fmt.Sprintf("[GET /subjects/{subject}/versions][%d] getSubjectVersionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSubjectVersionsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSubjectVersionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
