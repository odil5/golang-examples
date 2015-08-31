package main


import (
	"fmt"
	"net/http"
	"encoding/json"
	"handlers"
)


type Author struct {
	Id int `json:”id”`
	Name string `json:”name”`
}

// Объект для работы с жанрами
type Genre struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

// Объект для работы с книгами
type Book struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Genre string `json:"genre"`
	Author string `json:"author"`
}

type BookV2 struct {
    Id     int    `json:"id"`
    Name   string `json:"name"`
    Genre  Genre `json:"genre"`
    Author Author `json:"author"`
}


func author_handler(w http.ResponseWriter, r *http.Request) {

	author_marshalled, err := json.Marshal(Author{Id: 1, Name: "А.С Пушкин"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(w, string(author_marshalled))
}

func genres_handler(w http.ResponseWriter, r *http.Request) {

	genres_marshalled, err := json.Marshal(Genre{Id: 1, Name: "Сказки"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(w, string(genres_marshalled))
}

func books_handler(w http.ResponseWriter, r *http.Request) {
	
	tmp1 := Book{Id: 1, Name: "Сказки", Genre: "что тут", Author: "demo"}
	

	books_marshalled, err := json.Marshal(tmp1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(w, string(books_marshalled))

}

func main() {


	fmt.Print("Server running on port 8080")

	http.HandleFunc("/author", author_handler)
	http.HandleFunc("/genres", genres_handler)
	http.HandleFunc("/books", books_handler)

	http.ListenAndServe(":8080", nil)
}