package db

import (
	"employeemanagementsystemgen/database"
	"employeemanagementsystemgen/models"
	"fmt"
)

func Db_ShowDetailsTeamLead(employee string) (models.Employee,error) {
	fmt.Println("Hello this is DB layer Show EMployee Admin")
	fmt.Println(employee)
	db := database.Sqlclient()
	var Empl models.Employee
	query := "SELECT * FROM employeeOfficial WHERE userid='" + employee+ "'"
	querychk := "select userId from employeeOfficial where userid ='" + employee+ "'"
	resultchk, err := db.Query(querychk)
	if err != nil {
		fmt.Println("This is querry1 Error",err)
		return Empl , err
	}
	defer resultchk.Close()
	var emp models.Employee
	for resultchk.Next() {
		err := resultchk.Scan(&emp.UserId)
		if err != nil {
			fmt.Println("This is querry2 Error",err)
			return Empl , err
		}
	}
	if employee == emp.UserId {
		result, err := db.Query(query)
		if err != nil {
			fmt.Println("This is querry3 Error",err)
			return Empl, err

		}
		defer result.Close()


		for result.Next() {
			err := result.Scan(&Empl.Designation, &Empl.Department, &Empl.TeamLead, &Empl.JobType, &Empl.HealthInsurance, &Empl.LifeInsurance, &Empl.UserId,&Empl.Salary)
			if err != nil {
				fmt.Println("This is querry4 Error",err)
				return Empl, err
			}
		}
		defer db.Close()
	} else {
		var x int
		err := fmt.Errorf("math: square root of negative number %v", x)
		return Empl,err
	}
	return Empl , nil
}

