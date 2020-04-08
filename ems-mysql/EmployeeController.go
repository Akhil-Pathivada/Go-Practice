package ems

import "fmt"

func getEmployeeRecordById() {

	result := selectEmployeeRecordById(101)
	var employee Employee
	result.Scan(&employee.EmpId, &employee.FirstName, &employee.LastName)
	fmt.Println(employee.EmpId, employee.FirstName, employee.LastName)
}

func getEmployeeRecords() {

	results := selectAllRecords()

	for results.Next() {
		var employee Employee
		results.Scan(&employee.EmpId, &employee.FirstName, &employee.LastName)
		fmt.Println(employee.EmpId, employee.FirstName, employee.LastName)
	}
}

func EmployeeMain() {
	getEmployeeRecords()
	getEmployeeRecordById()
	insertEmployeeRecord(Employee{104, "Sachin", "Tendulkar"})
	deleteEmployeeRecord(103)
	updateEmployeeById(Employee{102, "RavindraBabu", "RAVULA"})
}
