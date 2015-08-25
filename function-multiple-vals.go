package main

import (
	"fmt"	
)

func someFunction()(int,int) {
	return 3, 5
}

func main() {

	a,b := someFunction()
	fmt.Println(a)
	fmt.Println(b)

}
