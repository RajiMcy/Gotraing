package main

import (
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":7000", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello ,world")

}
