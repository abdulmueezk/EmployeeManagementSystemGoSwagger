package service

import (
	"employeemanagementsystemgen/database/db"
	"employeemanagementsystemgen/models"
	"fmt"
)

func AddDetails(employee *models.Employee) (string,error) {
	fmt.Println("this is service layer AddEmployee")
	return db.Db_AddEmployeeDetails(employee)
}
func Updateetails(s string, employee *models.Employee) (string,error) {
	fmt.Println("this is service layer UpdateEmployee")
	return db.Db_UpdateEmployeeDetails(s, employee)
}
func DeleteDetails(id string,s string) (string, error) {
	fmt.Println("this is service layer DeleteEmployee")
	return db.Db_DeleteEmployeeDetails(id, s)
}
func ShowDetailsAdmin(employee string) (models.Employee,error) {
	fmt.Println("this is service layer ShowEmployeeAdmin")
	return db.Db_ShowDetailsAdmin(employee)
}

func ShowDetailsTeamLead(employee string) (models.Employee,error) {
	fmt.Println("this is service layer ShowEmployeeTeamLead")
	return db.Db_ShowDetailsTeamLead(employee)
}

func ShowDetailsEmployeeSelf(employee string) (models.Employee,error) {
	fmt.Println("this is service layer ShowEmployeeSelf")
	return db.Db_ShowDetailsSelf(employee)
}
