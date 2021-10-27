package handler

import (
	"employeemanagementsystemgen/gen/restapi/operations/admin"
	"employeemanagementsystemgen/service"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
)

type showEmployeeAdmin struct {

}

func ShowEmployeeAdmin() admin.ShowEmployeeHandler {
	return &showEmployeeAdmin{}
}

func (e *showEmployeeAdmin) Handle(params admin.ShowEmployeeParams) middleware.Responder  {
	fmt.Println("this is handler")
	employee ,err := service.ShowDetailsAdmin(params.UserID)
	if err != nil {
		fmt.Errorf("failed to add student: %s", err)
		return admin.NewShowEmployeeInternalServerError()
	}
	//			return admin.NewAddEmployeeCreated().WithPayload(toEmployeeGen(params.Body))
	return admin.NewShowEmployeeOK().WithPayload(toEmployeeGen(employee))
}

