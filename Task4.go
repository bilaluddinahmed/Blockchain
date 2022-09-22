package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

type Student struct {
	rollnumber int
	name       string
	address    string
}

var globe int

func NewStudent(rollno int, name string, address string) *Student {
	s := new(Student)
	s.rollnumber = rollno
	s.name = name
	s.address = address
	return s
}

type StudentList struct {
	list []*Student
}

func (ls *StudentList) CreateStudent(rollno int, name string, address string) *Student {
	st := NewStudent(rollno, name, address)
	ls.list = append(ls.list, st)
	globe = globe + 1
	return st
}
func Print(ls *StudentList) {
	for i := 0; i < globe; i++ {
		fmt.Printf("\n\n%s List %d %s\n\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		fmt.Printf("Student roll no	    %d\n", ls.list[i].rollnumber)
		fmt.Printf("Student name	    %s\n", ls.list[i].name)
		fmt.Printf("Student address	    %s\n", ls.list[i].address)
		hash := fmt.Sprintf("%x", sha256.Sum256([]byte(ls.list[i].name)))
		fmt.Printf("Hash: %x", hash)
	}
}

func main() {
	student := new(StudentList)
	student.CreateStudent(10, "Bilal", "Home")
	student.CreateStudent(20, "Taha", "Not Home")
	student.CreateStudent(30, "Musa", "Home Again")
	student.CreateStudent(40, "Saad", "Home 3")

	Print(student)

}
