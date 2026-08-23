package main

// 	A  simple Student gpa showcase using GO structs

import "fmt"



type Student struct {
	name            string
	registration_id int
	course          string
	semester        int
	gpa             float32
}

func main() {

	var Student1 Student
	var Student2 Student

	Student1.name = "Thandup"
	Student1.registration_id = 230410000020
	Student1.course = "BCA(Software Development)"
	Student1.semester = 5
	Student1.gpa = 7.5

	Student2.name = "Ankit"
	Student2.registration_id = 230420132546
	Student2.course = "BCA(Software Development)"
	Student2.semester = 5
	Student2.gpa = 8.5

	fmt.Println("FROM - MEDHAVI SKILLS UNIVERSITY")
	fmt.Printf("********************************\n")
	fmt.Printf("Hello %v ", Student1.name)
	fmt.Printf("(%v)\n", Student1.registration_id)
	fmt.Printf("%vth semester (%v)\n", Student1.semester, Student1.course)
	fmt.Printf("Your gpa for %vth semester is %v\n", Student1.semester, Student1.gpa)

	fmt.Printf("\n")

	fmt.Println("FROM - MEDHAVI SKILLS UNIVERSITY")
	fmt.Printf("********************************\n")
	fmt.Printf("Hello %v ", Student2.name)
	fmt.Printf("(%v)\n", Student2.registration_id)
	fmt.Printf("%vth semester (%v)\n", Student2.semester, Student2.course)
	fmt.Printf("Your gpa for %vth semester is %v\n", Student2.semester, Student2.gpa)


}
