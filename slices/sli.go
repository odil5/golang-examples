package main

import "fmt"


func main() {

	str1 := make([]string, 5)
	fmt.Println("emp:", str1)


	str1[0] = "demo1"
	str1[1] = "demo2"
	str1[2] = "demo3"
	fmt.Println("set:", str1)
	fmt.Println("get:", str1[2])


	c := make([]string, len(str1))
    copy(c, str1)
    fmt.Println("cpy:", c)


}


