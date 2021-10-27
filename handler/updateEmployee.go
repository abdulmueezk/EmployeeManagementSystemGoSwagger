package handler

import (
	"employeemanagementsystemgen/gen/restapi/operations/admin"
	"employeemanagementsystemgen/service"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
)

type updateEmployee struct {

}

func UpdateEmployee() admin.UpdateEmployeeHandler {
	return &updateEmployee{}
}

func (e *updateEmployee)Handle(params admin.UpdateEmployeeParams) middleware.Responder  {
	fmt.Println("this is handler layer UpdateEmployee")
	_,err := service.Updateetails(params.UserID, toUpdateEmployeeDomain(params.Body))
	if err != nil {
		fmt.Errorf("failed to add student: %s", err)
		return admin.NewUpdateEmployeeInternalServerError()
	}
	//			return admin.NewAddEmployeeCreated().WithPayload(toEmployeeGen(params.Body))
	return admin.NewUpdateEmployeeOK()
}

