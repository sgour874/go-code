package main

type Employee struct {
	EmpName   string  `json: "empname"`
	EmpSalary float64 `json: "empsalary"`
	Email     string  `json: "email"`
	City      string  `json: "city"`
	Id        int     `json: "id"`
}
