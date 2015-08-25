package main

import (
	"fmt"	
)

func someFunction(numbers ...int){
	fmt.Println(numbers, "")
	for _,num := range numbers{
		fmt.Println(num)
	}
}

func main() {

	someFunction(3,4)

	someFunction(8,9,6,31,2,4)

}
