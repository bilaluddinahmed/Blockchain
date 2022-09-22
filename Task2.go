package main

import "fmt"

type Company struct {
	companyName string
	employees   []Employee
}

type Employee struct {
	name     string
	salary   int
	position string
}

func main() {
	Bilal := Employee{"Bilal => ", 10000, " Game Developer\n"}
	Taha := Employee{"Taha => ", 5000, " Back-end Developer\n"}
	Musa := Employee{"Musa => ", 9000, " Front-end Developer\n"}

	employees := []Employee{Bilal, Taha, Musa}
	company := Company{"\nToyota\nEmployees Details : \n", employees}

	fmt.Printf("Company Name :  %v\n", company)
}
