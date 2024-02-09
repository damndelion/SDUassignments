package entity

import "fmt"

type Employee struct {
	Person
	JobTitle string
}

func (e Employee) PrintEmployeeInfo() {
	e.PrintInfo()
	fmt.Println("Job Title:", e.JobTitle)
}

func (e Employee) Print() {
	e.PrintEmployeeInfo()
}
