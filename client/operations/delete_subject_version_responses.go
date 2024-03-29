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

// DeleteSubjectVersionReader is a Reader for the DeleteSubjectVersion structure.
type DeleteSubjectVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSubjectVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteSubjectVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteSubjectVersionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewDeleteSubjectVersionUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteSubjectVersionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /subjects/{subject}/versions/{version}] delete_subject_version", response, response.Code())
	}
}

// NewDeleteSubjectVersionOK creates a DeleteSubjectVersionOK with default headers values
func NewDeleteSubjectVersionOK() *DeleteSubjectVersionOK {
	return &DeleteSubjectVersionOK{}
}

/*
DeleteSubjectVersionOK describes a response with status code 200, with default header values.

OK
*/
type DeleteSubjectVersionOK struct {
	Payload int64
}

// IsSuccess returns true when this delete subject version o k response has a 2xx status code
func (o *DeleteSubjectVersionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete subject version o k response has a 3xx status code
func (o *DeleteSubjectVersionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete subject version o k response has a 4xx status code
func (o *DeleteSubjectVersionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete subject version o k response has a 5xx status code
func (o *DeleteSubjectVersionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete subject version o k response a status code equal to that given
func (o *DeleteSubjectVersionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete subject version o k response
func (o *DeleteSubjectVersionOK) Code() int {
	return 200
}

func (o *DeleteSubjectVersionOK) Error() string {
	return fmt.Sprintf("[DELETE /subjects/{subject}/versions/{version}][%d] deleteSubjectVersionOK  %+v", 200, o.Payload)
}

func (o *DeleteSubjectVersionOK) String() string {
	return fmt.Sprintf("[DELETE /subjects/{subject}/versions/{version}][%d] deleteSubjectVersionOK  %+v", 200, o.Payload)
}

func (o *DeleteSubjectVersionOK) GetPayload() int64 {
	return o.Payload
}

func (o *DeleteSubjectVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSubjectVersionNotFound creates a DeleteSubjectVersionNotFound with default headers values
func NewDeleteSubjectVersionNotFound() *DeleteSubjectVersionNotFound {
	return &DeleteSubjectVersionNotFound{}
}

/*
DeleteSubjectVersionNotFound describes a response with status code 404, with default header values.

Schema not found
*/
type DeleteSubjectVersionNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete subject version not found response has a 2xx status code
func (o *DeleteSubjectVersionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete subject version not found response has a 3xx status code
func (o *DeleteSubjectVersionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete subject version not found response has a 4xx status code
func (o *DeleteSubjectVersionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete subject version not found response has a 5xx status code
func (o *DeleteSubjectVersionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete subject version not found response a status code equal to that given
func (o *DeleteSubjectVersionNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete subject version not found response
func (o *DeleteSubjectVersionNotFound) Code() int {
	return 404
}

func (o *DeleteSubjectVersionNotFound) Error() string {
	return fmt.Sprintf("[DELETE /subjects/{subject}/versions/{version}][%d] deleteSubjectVersionNotFound  %+v", 404, o.Payload)
}

func (o *DeleteSubjectVersionNotFound) String() string {
	return fmt.Sprintf("[DELETE /subjects/{subject}/versions/{version}][%d] deleteSubjectVersionNotFound  %+v", 404, o.Payload)
}

func (o *DeleteSubjectVersionNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteSubjectVersionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSubjectVersionUnprocessableEntity creates a DeleteSubjectVersionUnprocessableEntity with default headers values
func NewDeleteSubjectVersionUnprocessableEntity() *DeleteSubjectVersionUnprocessableEntity {
	return &DeleteSubjectVersionUnprocessableEntity{}
}

/*
DeleteSubjectVersionUnprocessableEntity describes a response with status code 422, with default header values.

Invalid version
*/
type DeleteSubjectVersionUnprocessableEntity struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete subject version unprocessable entity response has a 2xx status code
func (o *DeleteSubjectVersionUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete subject version unprocessable entity response has a 3xx status code
func (o *DeleteSubjectVersionUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete subject version unprocessable entity response has a 4xx status code
func (o *DeleteSubjectVersionUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete subject version unprocessable entity response has a 5xx status code
func (o *DeleteSubjectVersionUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this delete subject version unprocessable entity response a status code equal to that given
func (o *DeleteSubjectVersionUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the delete subject version unprocessable entity response
func (o *DeleteSubjectVersionUnprocessableEntity) Code() int {
	return 422
}

func (o *DeleteSubjectVersionUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /subjects/{subject}/versions/{version}][%d] deleteSubjectVersionUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *DeleteSubjectVersionUnprocessableEntity) String() string {
	return fmt.Sprintf("[DELETE /subjects/{subject}/versions/{version}][%d] deleteSubjectVersionUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *DeleteSubjectVersionUnprocessableEntity) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteSubjectVersionUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSubjectVersionInternalServerError creates a DeleteSubjectVersionInternalServerError with default headers values
func NewDeleteSubjectVersionInternalServerError() *DeleteSubjectVersionInternalServerError {
	return &DeleteSubjectVersionInternalServerError{}
}

/*
DeleteSubjectVersionInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteSubjectVersionInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete subject version internal server error response has a 2xx status code
func (o *DeleteSubjectVersionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete subject version internal server error response has a 3xx status code
func (o *DeleteSubjectVersionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete subject version internal server error response has a 4xx status code
func (o *DeleteSubjectVersionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete subject version internal server error response has a 5xx status code
func (o *DeleteSubjectVersionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete subject version internal server error response a status code equal to that given
func (o *DeleteSubjectVersionInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete subject version internal server error response
func (o *DeleteSubjectVersionInternalServerError) Code() int {
	return 500
}

func (o *DeleteSubjectVersionInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /subjects/{subject}/versions/{version}][%d] deleteSubjectVersionInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteSubjectVersionInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /subjects/{subject}/versions/{version}][%d] deleteSubjectVersionInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteSubjectVersionInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteSubjectVersionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
