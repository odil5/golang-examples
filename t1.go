package main

import "fmt"

func math1(a,b int) int{
	return a+b
}

func swap(x, y string, z int) (string, string, int) {

	return y, x, z
}

var a1, b1, c1 bool
var a2, b2, c2 int = 2, 4, 8
const  a3,a4,a5 = 556, 441, 996

func main() {

	f1 := "Demo 1 Message"
	
	fmt.Println(f1)

	f3 := math1(20000,300000000)

	fmt.Println(f3);
	
	// Use function
	a, b , d:= swap("hello", "world" , 3)
	fmt.Println(a, b, d)

	// Use variables 
	a1:= true
	b1:= true
	fmt.Println(a1, b1, c1);

	a2 := 88 
	fmt.Println(a2);

	//type conversions
	i := 42
	f := float64(i)
	fmt.Println(f);	

	// constants example
	fmt.Println(a3,a4);
}


