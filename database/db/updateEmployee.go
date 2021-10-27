package db

import (
	"employeemanagementsystemgen/database"
	"employeemanagementsystemgen/models"
	"fmt"
	"strconv"
)

func Db_UpdateEmployeeDetails(s string, employee *models.Employee) (string,error){
	fmt.Println("This is Db layer of Update Employee")
	healthInsurance, err := strconv.Atoi(strconv.FormatBool(employee.HealthInsurance))
	if err != nil {
		fmt.Errorf("failed to convert string to int, error : %v", err)
	}
	lifeInsurance, _ := strconv.Atoi(strconv.FormatBool(employee.HealthInsurance))
	if err != nil {
		fmt.Errorf("failed to convert string to int, error : %v", err)
	}
	var abcd error
	db := database.Sqlclient()
	fmt.Println("Update Employee Error ")
	employee.UserId = s
	query := "UPDATE employeeOfficial SET designation='"+employee.Designation+"' ,department='"+employee.Department+"', teamLead='"+employee.TeamLead+"',jobType='"+employee.JobType+"',healthInsurance='"+fmt.Sprintf("%v",healthInsurance)+"',lifeInsurance='"+fmt.Sprintf("%v",lifeInsurance)+"',salary='"+fmt.Sprintf("%v",employee.Salary)+"' where userId='"+s+"'"
	querychk := "select * from employeeOfficial where userid ='" +s+ "'"
	resultchk, err := db.Query(querychk)
	if err != nil {
		fmt.Println("Update Employee Error 1")
		panic(err)
	}
	defer resultchk.Close()
	var emp models.Employee
	/*healthInsurance1, err := strconv.Atoi(strconv.FormatBool(employee.HealthInsurance))
	if err != nil {
		fmt.Errorf("failed to convert string to int, error : %v", err)
	}
	lifeInsurance1, err := strconv.Atoi(strconv.FormatBool(employee.HealthInsurance))
	if err != nil {
		fmt.Errorf("failed to convert string to int, error : %v", err)
	}*/
	for resultchk.Next() {
		err := resultchk.Scan(&emp.Designation, &emp.Department, &emp.TeamLead, &emp.JobType, &emp.HealthInsurance, &emp.LifeInsurance, &emp.UserId,&emp.Salary)
		if err != nil {
			fmt.Println("Update Employee Error 2")
			panic(err.Error())
		}
	}
	if employee.UserId == emp.UserId {
		result, err := db.Exec(query)
		if err != nil {
			fmt.Println("Update Employee Error 3")
			panic(err)
		}
		defer db.Close()
		fmt.Println(result)
	} else {
		var x int
		err := fmt.Errorf("math: square root of negative number %v", x)
		abcd = err
	}
	return "",abcd
}
