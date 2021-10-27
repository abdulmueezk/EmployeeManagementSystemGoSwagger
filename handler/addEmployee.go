package handler

import (
	"employeemanagementsystemgen/gen/restapi/operations/admin"
	"employeemanagementsystemgen/service"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
)

type addEmployee struct {

}

func AddEmployee() admin.AddEmployeeHandler {
	return &addEmployee{}
}

func (e *addEmployee)Handle(params admin.AddEmployeeParams) middleware.Responder  {
	fmt.Println("this is service layer")
	_,err := service.AddDetails(toEmployeeDomain(params.Body))
	if err != nil {
				fmt.Errorf("failed to add student: %s", err)
				return admin.NewAddEmployeeInternalServerError()
			}
//			return admin.NewAddEmployeeCreated().WithPayload(toEmployeeGen(params.Body))
	return admin.NewAddEmployeeCreated().WithPayload("Employee Details Successfully Added")
}
