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

func GetProduct() []model.Product {
	return products
}
