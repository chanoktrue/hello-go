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

	data, err := json.Marshal(Emp{
		ID:    101,
		Name:  "test",
		Tel:   "00000",
		Email: "email@email.com",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(data)
	fmt.Println(string(data))
}
