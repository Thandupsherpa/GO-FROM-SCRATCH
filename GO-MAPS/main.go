package main

import "fmt"

func main()  {
	// GO-MAPS
	// Create Maps Using var and :=

	var a = map[string]string{"name":"Thandup","course":"BCA"}
	b := map[string]int{"age":19,"Reg-id":23041000020}

	fmt.Printf("a\t%v\n",a)
	fmt.Printf("b\t%v\n",b)

	// Create Maps Using the make() Function:
	var c = make(map[string]string) // This is a empty map

	c["food"] = "Momo"
	c["language"] = "GO"
	c["place"] = "singtam"
	fmt.Printf("c\t%v\n",c)

	// Accessing map element
	d := c["food"]
	fmt.Println(d)

	// Removing emelent from map by delete()
	delete(c,"food")
	fmt.Println(c)

	// Check for specific elements in map

	car := map[string]string{"brand":"Ford","model":"mustang","year":"1964","day":""}

	val1,ok1 := car["brand"]
	val2,ok2 := car["color"]
	val3,ok3 := car["day"]
	_,ok4 := car["model"]

	fmt.Println(val1,ok1)
	fmt.Println(val2,ok2)
	fmt.Println(val3,ok3)
	fmt.Println(ok4)
}