package main

import "fmt"

type emp struct {
	id    string
	name  string
	phone string
}

func main() {

	emps := []emp{}

	fmt.Println("emps: ", emps)

	emp1 := emp{
		id:    "1",
		name:  "John Doe",
		phone: "123-456-7890",
	}

	emp2 := emp{
		id:    "2",
		name:  "Jane Smith",
		phone: "987-654-3210",
	}

	emps = append(emps, emp1, emp2)

	fmt.Println("emps: ", emps)
}
