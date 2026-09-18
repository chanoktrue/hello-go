package main

import "fmt"

var productName [4]string
var productPrice [4]float64

func main() {
	productName[2] = "Apple"
	productPrice[2] = 1.99

	price := [4]float64{0.99, 1.49, 1.99, 2.49}

	fmt.Println("productName: ", productName)
	fmt.Println("productPrice: ", productPrice)
	fmt.Println("price: ", price)
}
