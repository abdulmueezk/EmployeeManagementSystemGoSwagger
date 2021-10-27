package handler

import (
	"employeemanagementsystemgen/gen/restapi/operations/employee"
	"employeemanagementsystemgen/service"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
)

type showEmployeeSelf struct {

}

func ShowEmployeeSelf() employee.ShowEmployeeSelfHandler {
	return &showEmployeeSelf{}
}

func (e *showEmployeeSelf) Handle(params employee.ShowEmployeeSelfParams) middleware.Responder  {
	fmt.Println("this is handler")
	empl ,err := service.ShowDetailsEmployeeSelf(params.UserID)
	if err != nil {
		fmt.Errorf("failed to add student: %s", err)
		return employee.NewShowEmployeeSelfInternalServerError()
	}
	//			return admin.NewAddEmployeeCreated().WithPayload(toEmployeeGen(params.Body))
	return employee.NewShowEmployeeSelfOK().WithPayload(toEmployeeGen(empl))
}
