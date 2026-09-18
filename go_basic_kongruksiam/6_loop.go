package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for k, v := range []string{"Alice", "Bob", "Charlie"} {
		fmt.Println(k, v)
	}
}
