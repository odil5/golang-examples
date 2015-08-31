package main

import (
   
    "fmt"
    "math/rand"
    "time"
)

func getRandomWord() string{

    answers := []string{
        "aaa",
        "bbb",
        "ccc",
        "ddd",
        "eee",
    }
    rand.Seed(time.Now().UTC().UnixNano())
    a := rand.Intn(5)

    return answers[a]
}

func main() {

    fmt.Println("\n" ,getRandomWord())
    
}