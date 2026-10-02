package service

import (
	"errors"

	"github.com/chanoktrue/go_api_chatgpt_package/model"
)

// Errors
var ErrorProductCodeExists = errors.New("product code already exists")
var ErrorProductNotFound = errors.New("product not found")

// Data
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

func GetProducts() []model.Product {
	return products
}

func GetProduct(code string) (model.Product, error) {
	for _, product := range products {
		if product.ProductCode == code {
			return product, nil
		}
	}

	return model.Product{}, ErrorProductNotFound
}

func CreateProduct(product model.Product) (model.Product, error) {

	if product.ProductCode == "" {
		return model.Product{}, errors.New("product code is required")
	}

	if product.ProductName == "" {
		return model.Product{}, errors.New("product name is required")
	}

	if product.Price <= 0 {
		return model.Product{}, errors.New("price must be graeter than 0")
	}

	products = append(products, product)

	return product, nil
}

func UpdateProduct(code string, input model.Product) (model.Product, error) {
	input.ProductCode = code

	for i, product := range products {
		if product.ProductCode == code {
			products[i] = input

			return products[i], nil
		}
	}

	return model.Product{}, ErrorProductNotFound
}

func DeleteProduct(code string) error {
	for i, proudct := range products {
		if proudct.ProductCode == code {
			products = append(products[:i], products[i+1:]...)

			return nil
		}
	}

	return ErrorProductNotFound
}
