package main

//https://github.com/bilaluddinahmed/Blockchain
import (
	"fmt"
)

type StudentRecord struct {
	rollnumber int
	name       string
	address    string
}

func DisplayStudent(s StudentRecord) {
	fmt.Println("Roll Number:", s.rollnumber)
	fmt.Println("Name:", s.name)
	fmt.Println("Address:", s.address)
}

func main() {
	var s1 StudentRecord
	s1.rollnumber = 10
	s1.name = "Bilal"
	s1.address = "Home"

	DisplayStudent(s1)
}
