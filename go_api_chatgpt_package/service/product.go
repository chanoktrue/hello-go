package service

import "github.com/chanoktrue/go_api_chatgpt_package/model"

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

func GetProduct(code string) (model.Product, bool) {
	for _, product := range products {
		if product.ProductCode == code {
			return product, true
		}
	}

	return model.Product{}, false
}

func CreateProduct(product model.Product) model.Product {
	products = append(products, product)

	return product
}

func UpdateProduct(code string, input model.Product) (model.Product, bool) {
	input.ProductCode = code

	for i, product := range products {
		if product.ProductCode == code {
			products[i] = input

			return products[i], true
		}
	}

	return model.Product{}, false
}

func DeleteProduct(code string) bool {
	for i, proudct := range products {
		if proudct.ProductCode == code {
			products = append(products[:i], products[i+1:]...)

			return true
		}
	}

	return false
}
