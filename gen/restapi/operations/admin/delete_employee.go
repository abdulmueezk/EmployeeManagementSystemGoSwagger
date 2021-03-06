// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteEmployeeHandlerFunc turns a function with the right signature into a delete employee handler
type DeleteEmployeeHandlerFunc func(DeleteEmployeeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteEmployeeHandlerFunc) Handle(params DeleteEmployeeParams) middleware.Responder {
	return fn(params)
}

// DeleteEmployeeHandler interface for that can handle valid delete employee params
type DeleteEmployeeHandler interface {
	Handle(DeleteEmployeeParams) middleware.Responder
}

// NewDeleteEmployee creates a new http.Handler for the delete employee operation
func NewDeleteEmployee(ctx *middleware.Context, handler DeleteEmployeeHandler) *DeleteEmployee {
	return &DeleteEmployee{Context: ctx, Handler: handler}
}

/* DeleteEmployee swagger:route DELETE /admin/delete/{user_id} admin deleteEmployee

Update Official Information to the employee account

*/
type DeleteEmployee struct {
	Context *middleware.Context
	Handler DeleteEmployeeHandler
}

func (o *DeleteEmployee) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteEmployeeParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
