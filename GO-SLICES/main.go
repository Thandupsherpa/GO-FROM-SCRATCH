package main

import "fmt"

func main() {

	//Go  Slices
	// Creating aslice with []datatype{value} format
	myslice := []int{1, 2, 3}
	fmt.Println(myslice)
	fmt.Println(len(myslice))
	fmt.Println(cap(myslice))

	//Creating a  slice from an array
	var myarr = [3]int{4,5,6}
	arrSlice := myarr[0:2]
	fmt.Println(arrSlice)

	//Creating a slice with the make() func
	sliceMake := make([]int,10)
	fmt.Println(sliceMake)
	fmt.Println(len(sliceMake))
	fmt.Println(cap(sliceMake))
	sliceMake[0] = 55
	fmt.Println(sliceMake)
	sliceMake = append(sliceMake, 33,66)
	fmt.Println(sliceMake)
	fmt.Println(sliceMake[0:3])

}
