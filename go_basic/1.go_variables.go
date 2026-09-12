package main

import "fmt"

var name, age = "Alice", 20 // Using var with multiple variables

func main() {

	// var name string = "Alice"
	// var age int = 20

	// name := "Alice" // Using short variable declaration
	// age := 20       // Using short variable declaration

	// var name, age = "Alice", 20 // Using var with multiple variables

	fmt.Println("Name: ", name, "Age: ", age)

	fmt.Printf("Type: %T\n", name)
	fmt.Printf("Type: %T\n", age)
}
