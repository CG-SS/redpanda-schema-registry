// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"redpanda-schema-registry/models"
)

// GetSubjectsReader is a Reader for the GetSubjects structure.
type GetSubjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSubjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSubjectsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetSubjectsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /subjects] get_subjects", response, response.Code())
	}
}

// NewGetSubjectsOK creates a GetSubjectsOK with default headers values
func NewGetSubjectsOK() *GetSubjectsOK {
	return &GetSubjectsOK{}
}

/*
GetSubjectsOK describes a response with status code 200, with default header values.

OK
*/
type GetSubjectsOK struct {
	Payload []string
}

// IsSuccess returns true when this get subjects o k response has a 2xx status code
func (o *GetSubjectsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get subjects o k response has a 3xx status code
func (o *GetSubjectsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get subjects o k response has a 4xx status code
func (o *GetSubjectsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get subjects o k response has a 5xx status code
func (o *GetSubjectsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get subjects o k response a status code equal to that given
func (o *GetSubjectsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get subjects o k response
func (o *GetSubjectsOK) Code() int {
	return 200
}

func (o *GetSubjectsOK) Error() string {
	return fmt.Sprintf("[GET /subjects][%d] getSubjectsOK  %+v", 200, o.Payload)
}

func (o *GetSubjectsOK) String() string {
	return fmt.Sprintf("[GET /subjects][%d] getSubjectsOK  %+v", 200, o.Payload)
}

func (o *GetSubjectsOK) GetPayload() []string {
	return o.Payload
}

func (o *GetSubjectsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubjectsInternalServerError creates a GetSubjectsInternalServerError with default headers values
func NewGetSubjectsInternalServerError() *GetSubjectsInternalServerError {
	return &GetSubjectsInternalServerError{}
}

/*
GetSubjectsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetSubjectsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get subjects internal server error response has a 2xx status code
func (o *GetSubjectsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get subjects internal server error response has a 3xx status code
func (o *GetSubjectsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get subjects internal server error response has a 4xx status code
func (o *GetSubjectsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get subjects internal server error response has a 5xx status code
func (o *GetSubjectsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get subjects internal server error response a status code equal to that given
func (o *GetSubjectsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get subjects internal server error response
func (o *GetSubjectsInternalServerError) Code() int {
	return 500
}

func (o *GetSubjectsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /subjects][%d] getSubjectsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSubjectsInternalServerError) String() string {
	return fmt.Sprintf("[GET /subjects][%d] getSubjectsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSubjectsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSubjectsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}