package main

import (
    "fmt"
    "os"
)

func main() {

    file, err := os.Open("simple_text.txt")
    if err != nil {
        fmt.Println("Error 1")
        return
    }
    defer file.Close()

    // get the file size
    stat, err := file.Stat()
    if err != nil {
        return
    }

    // read the file
    byteStrings := make([]byte, stat.Size())
    _, err = file.Read(byteStrings)
    if err != nil {
        return
    }

    text_f := string(byteStrings)
    fmt.Println(text_f)
}