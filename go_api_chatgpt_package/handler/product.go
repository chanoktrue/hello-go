package handler

import (
	"encoding/json"
	"net/http"

	"github.com/chanoktrue/go_api_chatgpt_package/model"
	"github.com/chanoktrue/go_api_chatgpt_package/service"
)

var products = []model.Product{
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

func GetProducts(w http.ResponseWriter, r *http.Request) {
	products := service.GetProducts()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")

	product, found := service.GetProduct(code)
	if !found {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode((product))
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product model.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	product = service.CreateProduct(product)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")

	var input model.Product

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	product, found := service.UpdateProduct(code, input)

	if !found {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")

	deleted := service.DeleteProduct(code)

	if !deleted {
		http.Error(w, "Product not found", http.StatusNotFound)

		return
	}

	w.WriteHeader(http.StatusNoContent)
}
