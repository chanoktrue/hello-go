package main

import "fmt"

type Product struct {
	ProductCode string `json:"productCode"`
	ProductName string `json:"productName"`
	Price       int    `json:"price"`
}

var products = []Product{
	{ProductCode: "P001", ProductName: "Speaker", Price: 2500},
	{ProductCode: "P002", ProductName: "Microphone", Price: 1500},
	{ProductCode: "P003", ProductName: "Amplifier", Price: 5000},
}

func main() {
	p1 := products[:1]
	p2 := products[2:]
	products = append(products[:1], products[1+1:]...)
	fmt.Println("p1 ", p1)
	fmt.Println("p2 ", p2)
	fmt.Println("product: ", products)
}
