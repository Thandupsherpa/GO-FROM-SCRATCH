package main

import (
	"fmt"

	
)

func main() {
	// GO-For Loop

	for i := 0; i < 5; i++ {
		fmt.Print(i," ")
	}

	for j := 0; j < 10; j+=2{
		// fmt.Println(j)
		if j == 4{
			continue
		}
		fmt.Println(j)
	}

	//Nested For
	adj := [2]string{"big","small"}
	fruit := [3]string{"Apple","Banana","Orange"}
	for a := 0; a < len(adj); a++{
		for b:=0; b < len(fruit); b++{
			fmt.Println(adj[a],fruit[b])
		}
	}

	// Range Keyword
	food := [3]string{"chips","cola","chocolate"}
	for idx, val:= range food{
		fmt.Printf("%v\t%v\n",idx,val)
	}
}
