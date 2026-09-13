package main

import "fmt"

func main() {
	country := map[string]string{}

	country["USA"] = "Washington D.C."
	country["Canada"] = "Ottawa"
	country["Mexico"] = "Mexico City"

	for key, value := range country {
		fmt.Println(key, ":", value)
	}
}
