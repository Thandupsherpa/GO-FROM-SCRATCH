package main

import "fmt"

// Declaring a function in GO
func test() {
	fmt.Println("This is a func declaration")
}

// Function with parameter
func familyName(fname string){
	fmt.Println("Hello ",fname, "Sherpa")
}

// Function with multiple parameters
func grade(name string, score int){
	fmt.Println("Hello ",name," Your score is: ",score)
}

//Return values
func sum(x int, y int) int{
	return  x+y
}

//Specifying the variable for returned value
func add(x int, y int) (sum int){
	sum = x + y
	return sum
}

//Naked Return
func naked_ret(x int, y int) (total int){
	total = x * y
	return
}

// test func
func sum2(x int, y int)(total int){
	total = x + y
	return

}

// Omit
func omit(age int, name string)(res1 int, res2 string){
	res1 = age + 1
	res2 = "Hello " + name
	return
}

//Recursive ()
func test_count(x int) int{
	if x == 11{
		return 0
	}
	fmt.Println(x)
	return test_count(x+1)
}



func main() {

	// Calling a Function
	test()
	test()
	test()
	
	//Calling familyName()
	familyName("Thandup")
	familyName("Dichen")

	//Calling grade()
	grade("Thandup",80)
	grade("Sherpa",66)

	//Calling sum()
	fmt.Println(sum(10,20))

	// Calling add()
	fmt.Println(add(5,6))

	//Calling naked_ret
	fmt.Println(naked_ret(2,3))

	// Calling sum2() and storing the return value in a variable
	total := sum2(2,3)
	fmt.Println(total)

	//calling omit() and omitting age
	_,name := omit(19,"Thandup")
	fmt.Println(name)

	//calling test_count()
	test_count(1)
}
