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

// PostSubjectReader is a Reader for the PostSubject structure.
type PostSubjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSubjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostSubjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewPostSubjectConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPostSubjectUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostSubjectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /subjects/{subject}] post_subject", response, response.Code())
	}
}

// NewPostSubjectOK creates a PostSubjectOK with default headers values
func NewPostSubjectOK() *PostSubjectOK {
	return &PostSubjectOK{}
}

/*
PostSubjectOK describes a response with status code 200, with default header values.

OK
*/
type PostSubjectOK struct {
	Payload *models.SubjectSchema
}

// IsSuccess returns true when this post subject o k response has a 2xx status code
func (o *PostSubjectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post subject o k response has a 3xx status code
func (o *PostSubjectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post subject o k response has a 4xx status code
func (o *PostSubjectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post subject o k response has a 5xx status code
func (o *PostSubjectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post subject o k response a status code equal to that given
func (o *PostSubjectOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post subject o k response
func (o *PostSubjectOK) Code() int {
	return 200
}

func (o *PostSubjectOK) Error() string {
	return fmt.Sprintf("[POST /subjects/{subject}][%d] postSubjectOK  %+v", 200, o.Payload)
}

func (o *PostSubjectOK) String() string {
	return fmt.Sprintf("[POST /subjects/{subject}][%d] postSubjectOK  %+v", 200, o.Payload)
}

func (o *PostSubjectOK) GetPayload() *models.SubjectSchema {
	return o.Payload
}

func (o *PostSubjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SubjectSchema)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSubjectConflict creates a PostSubjectConflict with default headers values
func NewPostSubjectConflict() *PostSubjectConflict {
	return &PostSubjectConflict{}
}

/*
PostSubjectConflict describes a response with status code 409, with default header values.

Incompatible schema
*/
type PostSubjectConflict struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post subject conflict response has a 2xx status code
func (o *PostSubjectConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post subject conflict response has a 3xx status code
func (o *PostSubjectConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post subject conflict response has a 4xx status code
func (o *PostSubjectConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this post subject conflict response has a 5xx status code
func (o *PostSubjectConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this post subject conflict response a status code equal to that given
func (o *PostSubjectConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the post subject conflict response
func (o *PostSubjectConflict) Code() int {
	return 409
}

func (o *PostSubjectConflict) Error() string {
	return fmt.Sprintf("[POST /subjects/{subject}][%d] postSubjectConflict  %+v", 409, o.Payload)
}

func (o *PostSubjectConflict) String() string {
	return fmt.Sprintf("[POST /subjects/{subject}][%d] postSubjectConflict  %+v", 409, o.Payload)
}

func (o *PostSubjectConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostSubjectConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSubjectUnprocessableEntity creates a PostSubjectUnprocessableEntity with default headers values
func NewPostSubjectUnprocessableEntity() *PostSubjectUnprocessableEntity {
	return &PostSubjectUnprocessableEntity{}
}

/*
PostSubjectUnprocessableEntity describes a response with status code 422, with default header values.

Invalid schema
*/
type PostSubjectUnprocessableEntity struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post subject unprocessable entity response has a 2xx status code
func (o *PostSubjectUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post subject unprocessable entity response has a 3xx status code
func (o *PostSubjectUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post subject unprocessable entity response has a 4xx status code
func (o *PostSubjectUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this post subject unprocessable entity response has a 5xx status code
func (o *PostSubjectUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this post subject unprocessable entity response a status code equal to that given
func (o *PostSubjectUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the post subject unprocessable entity response
func (o *PostSubjectUnprocessableEntity) Code() int {
	return 422
}

func (o *PostSubjectUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /subjects/{subject}][%d] postSubjectUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PostSubjectUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /subjects/{subject}][%d] postSubjectUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PostSubjectUnprocessableEntity) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostSubjectUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSubjectInternalServerError creates a PostSubjectInternalServerError with default headers values
func NewPostSubjectInternalServerError() *PostSubjectInternalServerError {
	return &PostSubjectInternalServerError{}
}

/*
PostSubjectInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostSubjectInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post subject internal server error response has a 2xx status code
func (o *PostSubjectInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post subject internal server error response has a 3xx status code
func (o *PostSubjectInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post subject internal server error response has a 4xx status code
func (o *PostSubjectInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post subject internal server error response has a 5xx status code
func (o *PostSubjectInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post subject internal server error response a status code equal to that given
func (o *PostSubjectInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post subject internal server error response
func (o *PostSubjectInternalServerError) Code() int {
	return 500
}

func (o *PostSubjectInternalServerError) Error() string {
	return fmt.Sprintf("[POST /subjects/{subject}][%d] postSubjectInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSubjectInternalServerError) String() string {
	return fmt.Sprintf("[POST /subjects/{subject}][%d] postSubjectInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSubjectInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostSubjectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
