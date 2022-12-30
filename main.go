package main

import (
	"fmt"
	"in-mem-sql/database/table"
)

/*
	In-memory database
		Components:
			1. Database -> It will hold the tables
			2. SQL -> It will take commands from user and translate them and query the database
*/

func main() {

	empTable := table.CreateTable("employee")

	empTable.AddColumn("empId", table.NumberDataType, []table.Constraint{table.UniqueConstraintType}).
		AddColumn("empName", table.VarcharDataType, []table.Constraint{table.NotNullConstraintType})

	empTable.Insert(map[string]interface{}{
		"empId":   int64(1),
		"empName": "Lovelesh",
	})

	err := empTable.Insert(map[string]interface{}{
		"empId":   int64(2),
		"empName": "mahima",
	})

	err = empTable.Insert(map[string]interface{}{
		"empId":   int64(3),
		"empName": "Khushvii",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%v", empTable.SelectAll())

}
