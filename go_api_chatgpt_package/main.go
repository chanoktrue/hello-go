package main

import (
	"fmt"
	"net/http"

	"github.com/chanoktrue/go_api_chatgpt_package/handler"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /products", handler.GetProducts)

	fmt.Println("Server runnin on http://localhost:8080")

	http.ListenAndServe(":8080", mux)

}
