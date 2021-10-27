// Code generated by go-swagger; DO NOT EDIT.

package team_lead

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewShowEmployeeTeamParams creates a new ShowEmployeeTeamParams object
//
// There are no default values defined in the spec.
func NewShowEmployeeTeamParams() ShowEmployeeTeamParams {

	return ShowEmployeeTeamParams{}
}

// ShowEmployeeTeamParams contains all the bound params for the show employee team operation
// typically these are obtained from a http.Request
//
// swagger:parameters showEmployeeTeam
type ShowEmployeeTeamParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Enter user_id to Show the Recoard of Team Member
	  Required: true
	  In: path
	*/
	UserID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewShowEmployeeTeamParams() beforehand.
func (o *ShowEmployeeTeamParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rUserID, rhkUserID, _ := route.Params.GetOK("user_id")
	if err := o.bindUserID(rUserID, rhkUserID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindUserID binds and validates parameter UserID from path.
func (o *ShowEmployeeTeamParams) bindUserID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.UserID = raw

	return nil
}