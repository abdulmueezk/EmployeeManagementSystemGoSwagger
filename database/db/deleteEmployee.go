package db

import (
	"employeemanagementsystemgen/database"
	"employeemanagementsystemgen/models"
	"fmt"
	//"strconv"
)

func Db_DeleteEmployeeDetails(id string, s string) (string,error){
	var abcd error
	db := database.Sqlclient()
	var employee models.Employee
/*
	healthInsurance, err := strconv.Atoi(strconv.FormatBool(employee.HealthInsurance))
	if err != nil {
		fmt.Errorf("failed to convert Bool to String, error : %v", err)
	}
	lifeInsurance, _ := strconv.Atoi(strconv.FormatBool(employee.HealthInsurance))
	if err != nil {
		fmt.Errorf("failed to convert Bool to String, error : %v", err)
	}*/
	employee.UserId = id
	employee.JobType = s
/*	healthInsurance = 0
	lifeInsurance = 0*/
	fmt.Println(id,s)
	query := "UPDATE employeeOfficial SET jobType='"+employee.JobType+"' where userid='"+employee.UserId+"'"
	fmt.Println("DB Error 1")
	querychk := "select * from employeeOfficial where userid ='" +id+ "'"
	fmt.Println("DB Error 2")
	resultchk, err := db.Query(querychk)
	fmt.Println("DB Error 3")
	if err != nil {
		panic(err)
	}
	defer resultchk.Close()
	fmt.Println("DB Error 4")
	var emp models.Employee
	fmt.Println("DB Error 5")
	for resultchk.Next() {
		fmt.Println("DB Error 6")
		err := resultchk.Scan(&emp.Designation, &emp.Department, &emp.TeamLead, &emp.JobType, &emp.HealthInsurance, &emp.LifeInsurance, &emp.UserId,&emp.Salary)
		fmt.Println("DB Error 7")
		if err != nil {
			fmt.Println("DB Error 8")
			panic(err.Error())
		}
		fmt.Println("scan Complete")
	}
	if id == emp.UserId {
		fmt.Println("DB Error 9")
		result, err := db.Exec(query)
		if err != nil {
			fmt.Println("DB Error 10")
			panic(err)
		}
		defer db.Close()
		fmt.Println(result)
	} else {
		fmt.Println("DB Error 11")
		var x int
		err := fmt.Errorf("math: square root of negative number %v", x)
		abcd = err
	}
	fmt.Println("DB Error 12")
	return "",abcd
}

