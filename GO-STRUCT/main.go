package main

import "fmt"

// To declare a structure in GO, use type and struct keywords:
type Person struct {
	name   string
	reg_id int
	course string
	sem    int
}

func main() {

	var person1 Person
	var person2 Person

	person1.name = "Thandup"
	person1.reg_id = 230410000020
	person1.course = "BCA"
	person1.sem = 5

	person2.name = "Ankit"
	person2.reg_id = 23041000555
	person2.course = "BCA"
	person2.sem = 5

	fmt.Println("Name: ", person1.name)
	fmt.Println("Reg-id: ", person1.reg_id)
	fmt.Println("course: ", person1.course)
	fmt.Println("sem: ", person1.sem)

	fmt.Printf("\n")

	fmt.Println("Name: ", person2.name)
	fmt.Println("Reg-id: ", person2.reg_id)
	fmt.Println("course: ", person2.course)
	fmt.Println("sem: ", person2.sem)

	fmt.Printf("\n")


	printPerson(person1)

	fmt.Printf("\n")

	printPerson(person2)
}
// Pass struct as functionn arguments


func printPerson(pers Person)  {
	fmt.Println("Name: ", pers.name)
	fmt.Println("Reg-id: ", pers.reg_id)
	fmt.Println("course: ", pers.course)
	fmt.Println("sem: ", pers.sem)
	
}
