package main

import (
	"fmt"
)

func main() {

	
	anonVar := func(name string) string {
	   return "Hello, " + name
	}

	anotherFunction(anonVar("wewe"))

}


func anotherFunction(f func(string) string) {
	result := f("David")
	fmt.Println(result) // Prints "Hiya, David"
}