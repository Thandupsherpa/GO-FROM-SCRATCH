package main

import "fmt"

func main(){
	//With var keyword

	var arr = [3]int{1,2,3} 

	var arr1 = [...]int{1,2,3}
	fmt.Println(arr)
	fmt.Println(arr1)


	//With walrus operator :=

	arr2 := [4]int{1,2,3,4}
	fmt.Println(arr2)

	//Accessing a element of an array
	fmt.Println(arr2[2])	

	//Changing element of an array

	prices := [...]int{20,30,50}

	prices[2] = 40

	fmt.Println(prices)

	// Initializing a specific elements of an array

	amount := [5]int{2:30,4:50}
	fmt.Println(amount)
	fmt.Println(len(amount))


}
