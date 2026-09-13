package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var productList []Product
var productMu sync.RWMutex

func init() {
	productJSON := `[
		{"id": 1, "name": "AA", "price": 10},
		{"id": 2, "name": "BB", "price": 20},
		{"id": 3, "name": "CC", "price": 30}
	]`

	err := json.Unmarshal([]byte(productJSON), &productList)
	if err != nil {
		log.Fatal(err)
	}
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/product")
	if path != "" && !strings.HasPrefix(path, "/") {
		http.NotFound(w, r)
		return
	}

	if path == "/" {
		path = ""
	}

	if path == "" && r.Method == http.MethodGet {
		productMu.RLock()
		defer productMu.RUnlock()
		writeJSON(w, http.StatusOK, productList)
		return
	}

	if path == "" && r.Method == http.MethodPost {
		var newProduct Product
		if err := decodeProduct(r, &newProduct); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		productMu.Lock()
		newProduct.ID = nextProductID()
		productList = append(productList, newProduct)
		productMu.Unlock()

		writeJSON(w, http.StatusCreated, newProduct)
		return
	}

	id, err := strconv.Atoi(strings.TrimPrefix(path, "/"))
	if err != nil || id <= 0 || strings.Contains(strings.TrimPrefix(path, "/"), "/") {
		http.NotFound(w, r)
		return
	}

	productMu.Lock()
	defer productMu.Unlock()
	for index := range productList {
		if productList[index].ID != id {
			continue
		}

		switch r.Method {
		case http.MethodGet:
			writeJSON(w, http.StatusOK, productList[index])
		case http.MethodPut:
			var updatedProduct Product
			if err := decodeProduct(r, &updatedProduct); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			updatedProduct.ID = id
			productList[index] = updatedProduct
			writeJSON(w, http.StatusOK, updatedProduct)
		case http.MethodDelete:
			productList = append(productList[:index], productList[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
		return
	}

	http.NotFound(w, r)
}

func nextProductID() int {
	nextID := 1
	for _, product := range productList {
		if product.ID >= nextID {
			nextID = product.ID + 1
		}
	}
	return nextID
}

func decodeProduct(r *http.Request, product *Product) error {
	decoder := json.NewDecoder(io.LimitReader(r.Body, 1<<20))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(product); err != nil {
		return err
	}
	if strings.TrimSpace(product.Name) == "" || product.Price < 0 {
		return &validationError{message: "name is required and price must be non-negative"}
	}
	return nil
}

type validationError struct {
	message string
}

func (e *validationError) Error() string {
	return e.message
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(value); err != nil {
		log.Printf("encode product response: %v", err)
	}
}

/*
	API endpoints:
	GET    /product       list products
	GET    /product/{id}  get one product
	POST   /product       create product
	PUT    /product/{id}  replace product
	DELETE /product/{id}  delete product
*/

func middlewareHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Before handler middle start")
		handler.ServeHTTP(w, r)
		fmt.Println("middleware finised")
	})
}

func main() {
	productHandlerWithMiddleware := middlewareHandler(http.HandlerFunc(productHandler))
	http.Handle("/", productHandlerWithMiddleware)
	http.Handle("/product", productHandlerWithMiddleware)
	http.Handle("/product/", productHandlerWithMiddleware)

	log.Println("HTTP server listening on http://127.0.0.1:8000/product")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
