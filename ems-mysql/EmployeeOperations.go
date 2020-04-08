package ems

import (
	"database/sql"
	"fmt"
)

func selectAllRecords() *sql.Rows {

	db := openDatabaseConnection()
	results, err := db.Query("SELECT * FROM Employees")

	if err != nil {
		fmt.Println("Error! Getting records...", err)
	}
	defer db.Close()
	return results
}

func selectEmployeeRecordById(EmpId int) *sql.Row {

	db := openDatabaseConnection()
	result := db.QueryRow("SELECT * FROM Employees WHERE EmpId = ?", EmpId)
	defer db.Close()
	return result
}

func insertEmployeeRecord(employee Employee) {
	db := openDatabaseConnection()
	insert, err := db.Query("INSERT INTO Employees VALUES (?, ?, ?)", employee.EmpId, employee.FirstName, employee.LastName)

	if err != nil {
		fmt.Println("Error! Inserting records...")
	}
	defer insert.Close()
	defer db.Close()
}

func deleteEmployeeRecord(EmpId int) {
	db := openDatabaseConnection()
	db.QueryRow("DELETE FROM Employees WHERE EmpId = ?", EmpId)
}

func updateEmployeeById(employee Employee) {
	db := openDatabaseConnection()
	db.QueryRow("UPDATE Employees SET LastName = ? WHERE EmpId = ?", employee.FirstName, employee.EmpId)
}
