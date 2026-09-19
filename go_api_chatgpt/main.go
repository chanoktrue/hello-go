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

func getProduct(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")

	// fmt.Println("code =")

	for _, product := range products {
		if product.ProductCode == code {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode((product))
			return
		}
	}

	http.Error(w, "Product not found", http.StatusNotFound)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")

	fmt.Println("code =", code)

	var input Product

	fmt.Println("input1 =", input)

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid JONS", http.StatusBadRequest)
		return
	}

	fmt.Println("input2 =", input)

	for i, product := range products {
		if product.ProductCode == code {
			products[i] = input

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(products[i])
			return
		}
	}

	http.Error(w, "Product not found", http.StatusNotFound)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")

	for i, product := range products {
		if product.ProductCode == code {
			products = append(products[:i], products[i+1:]...)

			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Product not found", http.StatusNotFound)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /products", getProducts)
	mux.HandleFunc("GET /products/{code}", getProduct)
	mux.HandleFunc("POST /products", createProduct)
	mux.HandleFunc("PUT /products/{code}", updateProduct)
	mux.HandleFunc("DELETE /products/{code}", deleteProduct)

	fmt.Println("Server runnin on http://localhost:8080")

	http.ListenAndServe(":8080", mux)

}
