package main

import (
	"fmt"
)

type user struct {
	name string
	age int
	company string
}

func main() {
	
	fmt.Println(user{"Nikol", 25, "Apple"})

	var1 := user{name:"Nikol", age:25, company:"Apple"}
	fmt.Println(var1.age)

	var22 := &var1
	var22.name = "Vanya"
	fmt.Println(var1.name)
}

