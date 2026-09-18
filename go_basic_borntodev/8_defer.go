package main

import "fmt"

func add(a int, b int) {
	result := a + b
	fmt.Println("Result:", result)
}

func loop() {
	for i := 0; i < 5; i++ {
		fmt.Println("Loop:", i)
	}
}

func deferLoop() {
	for i := 0; i < 5; i++ {
		defer fmt.Println("Defer Loop:", i)
	}
}

func main() {
	fmt.Println("Hello, World!")
	defer fmt.Println("End")

	defer add(5, 3)
	defer add(10, 20)
	defer add(15, 25)

	loop()
	deferLoop()

}
