package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")
	secNa := r.URL.Query().Get("secondname")
	inputTypes := "Калкулятор <br> " +
		"<form>" +
		"a=<input type='text' name='a'><br>" +
		"b=<input type='text' name='b'> <br> " +
		"<input type='submit'> nimadur " +
		"</form>"
	header := "<html>"
	footer := "</html>"

	a := []int{r.URL.Query().Get("a")}
	b := r.URL.Query().Get("b")
	c := a + b
	result := "Salom  <b>" + name + "</b> " + secNa + " <br> a+b=" + c + " <br>" + inputTypes
	io.WriteString(w, header+result+footer)

}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/ta/", hello)

	http.ListenAndServe(":8000", mux)

}
