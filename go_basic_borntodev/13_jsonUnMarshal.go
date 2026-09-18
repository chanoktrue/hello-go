package main

import (
	"encoding/json"
	"fmt"
)

type Emp struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Tel   string `json:"tel"`
	Email string `json:"email"`
}

func main() {
	data := []byte(`{
		"id": 101,
		"name": "test",
		"tel": "00000",
		"email": "email@email.com"
	}`)

	var emp Emp

	err := json.Unmarshal(data, &emp)

	if err != nil {
		panic(err)
	}

	fmt.Println(emp)
	fmt.Println(emp.Name)
}
