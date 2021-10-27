package handler

import (
	"employeemanagementsystemgen/gen/restapi/operations/team_lead"
	"employeemanagementsystemgen/service"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
)

type showEmployeeTeamLead struct {

}

func ShowEmployeeTeamLead() team_lead.ShowEmployeeTeamHandler {
	return &showEmployeeTeamLead{}
}

func (e *showEmployeeTeamLead) Handle(params team_lead.ShowEmployeeTeamParams) middleware.Responder  {
	fmt.Println("this is handler")
	emply ,err := service.ShowDetailsTeamLead(params.UserID)
	if err != nil {
		fmt.Errorf("failed to add student: %s", err)
		return team_lead.NewShowEmployeeTeamInternalServerError()
	}
	return team_lead.NewShowEmployeeTeamOK().WithPayload(toTeamLeadEmployeeGen(emply))
}
