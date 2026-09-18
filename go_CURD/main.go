package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Hello API")
}

func main() {
	http.HandleFunc("/hello", hello)

	fmt.Println("Server runnin at http://localhost:8000")

	http.ListenAndServe(":8000", nil)
}
