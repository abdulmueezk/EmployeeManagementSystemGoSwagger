package main

import (
	"employeemanagementsystemgen/gen/restapi"
	"employeemanagementsystemgen/gen/restapi/operations"
	"employeemanagementsystemgen/handler"
	"flag"
	"github.com/go-openapi/loads"
	"log"
)
var portFlag = flag.Int("port", 3001, "Port to run this service on")
func main() {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON,"")
	if err != nil {
		log.Fatalln(err)
	}
	api := operations.NewEmployeeManagemntSystemAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()
	flag.Parse()
	server.Port = *portFlag
	api.AdminAddEmployeeHandler = handler.AddEmployee()
	api.AdminShowEmployeeHandler = handler.ShowEmployeeAdmin()
	api.AdminUpdateEmployeeHandler = handler.UpdateEmployee()
	api.AdminDeleteEmployeeHandler = handler.DeleteEmployee()
	api.EmployeeShowEmployeeSelfHandler = handler.ShowEmployeeSelf()
	api.TeamLeadShowEmployeeTeamHandler = handler.ShowEmployeeTeamLead()
	if err := server.Serve(); err != nil{
		log.Fatalln(err)
	}
}
