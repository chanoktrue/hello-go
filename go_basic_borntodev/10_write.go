package main

import "os"

func main() {
	data1 := []byte("Hello, World!")
	err := os.WriteFile("go_basic_borntodev/data.txt", data1, 0644)
	if err != nil {
		panic(err)
	}
	f, ferr := os.Create("name")
	if ferr != nil {
		panic(ferr)
	}
	defer f.Close()
}
