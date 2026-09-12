package main

import "fmt"

func main() {
	names := []string{"Alice", "Bob", "Charlie"}

	fmt.Println(names[1:2])

	fmt.Println("Length:", len(names))

	names = append(names, "David", "Eve")

	println("length: ", len(names))
	println(" ")

	println("capacity: ", cap(names))

	names[2] = "Charlie Updated"

	for key, name := range names {
		println(key, name)
	}

}
