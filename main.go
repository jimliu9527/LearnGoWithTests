package main

import (
	"learn-go-with-tests/greet"
	"net/http"
)

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	greet.Greet(w, "world")
}

func main() {
	http.ListenAndServe(":5002", http.HandlerFunc(MyGreeterHandler))

}
