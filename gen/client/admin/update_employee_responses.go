// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateEmployeeReader is a Reader for the UpdateEmployee structure.
type UpdateEmployeeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateEmployeeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateEmployeeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateEmployeeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateEmployeeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateEmployeeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateEmployeeOK creates a UpdateEmployeeOK with default headers values
func NewUpdateEmployeeOK() *UpdateEmployeeOK {
	return &UpdateEmployeeOK{}
}

/* UpdateEmployeeOK describes a response with status code 200, with default header values.

Successfully Update
*/
type UpdateEmployeeOK struct {
}

func (o *UpdateEmployeeOK) Error() string {
	return fmt.Sprintf("[PATCH /admin/update/{user_id}][%d] updateEmployeeOK ", 200)
}

func (o *UpdateEmployeeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateEmployeeBadRequest creates a UpdateEmployeeBadRequest with default headers values
func NewUpdateEmployeeBadRequest() *UpdateEmployeeBadRequest {
	return &UpdateEmployeeBadRequest{}
}

/* UpdateEmployeeBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateEmployeeBadRequest struct {
}

func (o *UpdateEmployeeBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /admin/update/{user_id}][%d] updateEmployeeBadRequest ", 400)
}

func (o *UpdateEmployeeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateEmployeeUnauthorized creates a UpdateEmployeeUnauthorized with default headers values
func NewUpdateEmployeeUnauthorized() *UpdateEmployeeUnauthorized {
	return &UpdateEmployeeUnauthorized{}
}

/* UpdateEmployeeUnauthorized describes a response with status code 401, with default header values.

Unotherized
*/
type UpdateEmployeeUnauthorized struct {
}

func (o *UpdateEmployeeUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /admin/update/{user_id}][%d] updateEmployeeUnauthorized ", 401)
}

func (o *UpdateEmployeeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateEmployeeInternalServerError creates a UpdateEmployeeInternalServerError with default headers values
func NewUpdateEmployeeInternalServerError() *UpdateEmployeeInternalServerError {
	return &UpdateEmployeeInternalServerError{}
}

/* UpdateEmployeeInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type UpdateEmployeeInternalServerError struct {
}

func (o *UpdateEmployeeInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /admin/update/{user_id}][%d] updateEmployeeInternalServerError ", 500)
}

func (o *UpdateEmployeeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
