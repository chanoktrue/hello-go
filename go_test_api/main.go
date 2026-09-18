package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Product struct {
	ProductCode string `json:"productcode"`
	ProductName string `json:"productname"`
	Price       int    `json:"price"`
}

var products = []Product{
	{ProductCode: "P001", ProductName: "Speaker", Price: 1500},
	{ProductCode: "P002", ProductName: "Microphone", Price: 2500},
}

// GET /products
func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// POST /products
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

// PUT /products/{code}
func updateProduct(w http.ResponseWriter, r *http.Request) {
	code := strings.TrimPrefix(r.URL.Path, "/products/")

	var newProduct Product

	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	for i, product := range products {
		if product.ProductCode == code {
			newProduct.ProductCode = code
			products[i] = newProduct

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(newProduct)
			return
		}
	}

	http.Error(w, "Product not found", http.StatusNotFound)
}

// DELETE /products/{code}
func deleteProduct(w http.ResponseWriter, r *http.Request) {
	code := strings.TrimPrefix(r.URL.Path, "/products/")

	for i, product := range products {
		if product.ProductCode == code {
			products = append(products[:i], products[i+1:]...)

			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Product not found", http.StatusNotFound)
}

func productsHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		getProducts(w, r)

	case http.MethodPost:
		createProduct(w, r)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func productHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodPut:
		updateProduct(w, r)

	case http.MethodDelete:
		deleteProduct(w, r)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {

	http.HandleFunc("/products", productsHandler)
	http.HandleFunc("/products/", productHandler)

	fmt.Println("Server running at http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
