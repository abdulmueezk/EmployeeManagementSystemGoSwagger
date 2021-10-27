// Code generated by go-swagger; DO NOT EDIT.

package team_lead

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"employeemanagementsystemgen/gen/models"
)

// ShowEmployeeTeamReader is a Reader for the ShowEmployeeTeam structure.
type ShowEmployeeTeamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShowEmployeeTeamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShowEmployeeTeamOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewShowEmployeeTeamNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewShowEmployeeTeamInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewShowEmployeeTeamOK creates a ShowEmployeeTeamOK with default headers values
func NewShowEmployeeTeamOK() *ShowEmployeeTeamOK {
	return &ShowEmployeeTeamOK{}
}

/* ShowEmployeeTeamOK describes a response with status code 200, with default header values.

Operation Successful
*/
type ShowEmployeeTeamOK struct {
	Payload *models.TeamLeadEmployee
}

func (o *ShowEmployeeTeamOK) Error() string {
	return fmt.Sprintf("[GET /teamLead/showteam/{user_id}][%d] showEmployeeTeamOK  %+v", 200, o.Payload)
}
func (o *ShowEmployeeTeamOK) GetPayload() *models.TeamLeadEmployee {
	return o.Payload
}

func (o *ShowEmployeeTeamOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TeamLeadEmployee)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowEmployeeTeamNotFound creates a ShowEmployeeTeamNotFound with default headers values
func NewShowEmployeeTeamNotFound() *ShowEmployeeTeamNotFound {
	return &ShowEmployeeTeamNotFound{}
}

/* ShowEmployeeTeamNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ShowEmployeeTeamNotFound struct {
}

func (o *ShowEmployeeTeamNotFound) Error() string {
	return fmt.Sprintf("[GET /teamLead/showteam/{user_id}][%d] showEmployeeTeamNotFound ", 404)
}

func (o *ShowEmployeeTeamNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewShowEmployeeTeamInternalServerError creates a ShowEmployeeTeamInternalServerError with default headers values
func NewShowEmployeeTeamInternalServerError() *ShowEmployeeTeamInternalServerError {
	return &ShowEmployeeTeamInternalServerError{}
}

/* ShowEmployeeTeamInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ShowEmployeeTeamInternalServerError struct {
}

func (o *ShowEmployeeTeamInternalServerError) Error() string {
	return fmt.Sprintf("[GET /teamLead/showteam/{user_id}][%d] showEmployeeTeamInternalServerError ", 500)
}

func (o *ShowEmployeeTeamInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
