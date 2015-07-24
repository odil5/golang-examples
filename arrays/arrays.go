package main

import "fmt"


func main() {

	// array with zero values
	var a1[10]int
	fmt.Println("array:", a1)

	a1[0] = 5
	a1[1] = 6
	fmt.Println("get value:", a1[0])

	// get Length our array
	fmt.Println("get value:", len(a1))

	// declare array and initialize an array
	a2 := [5]int{3, 5, 8, 4, 5}
    fmt.Println("declared:", a2)
    fmt.Println("get 2 element:", a2[2])

	var twoDArr [2][3]int
    twoDArr[0][1] = 2
	twoDArr[0][2] = 3

    fmt.Println("2d: ", twoDArr)

}


