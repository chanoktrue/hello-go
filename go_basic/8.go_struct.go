package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {

	person1 := Person{Name: "Alice", Age: 30}
	person2 := Person{Name: "Bob", Age: 25}

	people := []Person{person1, person2}

	for _, person := range people {
		fmt.Println(person.Name, "is", person.Age, "years old.")
	}

}
