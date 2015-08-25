package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type User struct {
		Id     int
		Name   string
		DateBirth int
		Access []string

	}

	user := User{
		Id:     1,
		Name:   "Kolo",
		DateBirth: 1990,
		Access: []string{"personal_page", "change_pass", "settings"},
	}

	b, err := json.Marshal(user)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
    res := &Response2{}
    json.Unmarshal([]byte(str), &res)
    fmt.Println(res)
    fmt.Println(res)
}