package main

import "fmt"

var product = make(map[string]float64)

func main() {
	fmt.Println("product: ", product)

	//adding key-value pairs to the map
	product["Apple"] = 1.99
	product["Banana"] = 0.99
	product["Orange"] = 1.49
	fmt.Println("product: ", product)

	//deleting key-value pair from the map
	delete(product, "Banana")
	fmt.Println("product: ", product)

	//updating value of a key in the map
	product["Apple"] = 2.49
	fmt.Println("product: ", product)

	name := map[string]string{
		"firstName": "John",
		"lastName":  "Doe",
	}
	fmt.Println("name: ", name)

}
