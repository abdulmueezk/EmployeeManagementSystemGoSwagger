package db

import (
	"employeemanagementsystemgen/database"
	"employeemanagementsystemgen/models"
	"fmt"
	"strconv"
)

func Db_AddEmployeeDetails(employee *models.Employee) (string,error) {
	fmt.Println("Hello this is DB layer Add EMployee")
fmt.Println(employee)
	healthInsurance, err := strconv.Atoi(strconv.FormatBool(employee.HealthInsurance))
		if err != nil {
			fmt.Errorf("failed to convert Bool to String, error : %v", err)
		}
		lifeInsurance, _ := strconv.Atoi(strconv.FormatBool(employee.HealthInsurance))
		if err != nil {
			fmt.Errorf("failed to convert Bool to String, error : %v", err)
		}
db := database.Sqlclient()
	result, err := db.Query("INSERT INTO employeeOfficial (designation,department,teamLead,jobType,healthInsurance,lifeInsurance,userId,salary)VALUES ( '"+employee.Designation+"' ,'"+employee.Department+"','"+employee.TeamLead+"','"+employee.JobType+"','"+fmt.Sprintf("%v",healthInsurance)+"','"+fmt.Sprintf("%v",lifeInsurance)+"','"+employee.UserId+"','"+fmt.Sprintf("%v",employee.Salary)+"')")
		if err != nil {
			return "" , err
		}
		defer result.Close()
	fmt.Println(result)
	return "Task performed", nil
}