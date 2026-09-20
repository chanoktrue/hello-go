package main

import "fmt"

func main() {
	name := []string{"Alice", "Bob", "Charlie"}
	name = append(name, "David")

	fmt.Println("name: ", name)
	fmt.Println("name: ", name[1:2])
	fmt.Println("name: ", name[0:1])
	fmt.Println("name: ", name[1:3])

	name = append(name[:2], name[2+1:]...)
	fmt.Println("---> ", name)
}
