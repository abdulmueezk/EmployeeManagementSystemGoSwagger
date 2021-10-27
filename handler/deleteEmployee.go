package handler

import (
	"employeemanagementsystemgen/gen/restapi/operations/admin"
	"employeemanagementsystemgen/service"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
)

type deleteEmployee struct {

}

func DeleteEmployee() admin.DeleteEmployeeHandler {
	return &deleteEmployee{}
}

func (e *deleteEmployee)Handle(params admin.DeleteEmployeeParams) middleware.Responder  {
	fmt.Println("this is handler layer deleteEmployee")
	_,err := service.DeleteDetails(params.UserID , swag.StringValue(params.Body.JobType))
	if err != nil {
		fmt.Errorf("Failed to Delete Employee: %s", err)
		return admin.NewDeleteEmployeeInternalServerError()
	}
	return admin.NewDeleteEmployeeCreated()
}

