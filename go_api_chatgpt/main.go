package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ProductCode string `json:"productCode"`
	ProductName string `json:"productName"`
	Price       int    `json:"price"`
}

var products = []Product{
	{
		ProductCode: "P001",
		ProductName: "Speaker",
		Price:       2500,
	},
	{
		ProductCode: "P002",
		ProductName: "Microphone",
		Price:       1500,
	},
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	var product Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	products = append(products, product)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(product)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /products", getProducts)
	mux.HandleFunc("POST /products", createProduct)

	fmt.Println("Server runnin on http://localhost:8080")

	http.ListenAndServe(":8080", mux)

}
