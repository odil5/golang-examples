package main

import "fmt"

var c, python, java bool
var a,b,d int

func main() {

	f1 := "Demo 1 Message"
	f2 := "Demo 2 Message"
	fmt.Println(f1)

	d := demoMyFunc(1,1,1);

	fmt.Println(d)
	
	fmt.Println(f2)
}



func demoMyFunc(a, b, d int) int {
	return a+b+d
}

