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

		Nice to have features:
			1. Add functionality to have AND and OR conditionals for update method
			2. Updating row via primary key(Primary key concept)
*/

func main() {

	empTable := table.CreateTable("employee")

	empTable.AddColumn("empId", table.NumberDataType, []table.Constraint{table.UniqueConstraintType}).
		AddColumn("empName", table.VarcharDataType, []table.Constraint{table.NotNullConstraintType})

	empTable.InsertRow(map[string]interface{}{
		"empId":   int64(1),
		"empName": "Lovelesh",
	})

	err := empTable.InsertRow(map[string]interface{}{
		"empId":   int64(2),
		"empName": "mahima",
	})

	err = empTable.InsertRow(map[string]interface{}{
		"empId":   int64(3),
		"empName": "Khushvii",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	if err = empTable.DeleteRow("empName", "Lovelesh"); err != nil {
		fmt.Println(err)
		return
	}

	empTable.Update(map[string]interface{}{
		"empName": "KHUSHVII",
	}, "empName", "Khushvii")

	fmt.Printf("%v", empTable.SelectAll())
}
