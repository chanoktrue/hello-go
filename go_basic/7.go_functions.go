package main

import "fmt"

// function with parameter
func printMessage(fname string) {
	fmt.Println("Hello, " + fname + "!")
}

// return value
func sum(a int, b int) int {
	result := a + b
	return result
}

// recursive function
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// function with multiple return values
func swap(x, y string) (string, string) {
	return y, x
}

// returning a function
func testcount(x int) int {
	if x == 11 {
		return 0
	}
	fmt.Println(x)
	return testcount(x + 1)
}

// varidic function
func sumAll(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// main function
func main() {
	printMessage("Alice")
	printMessage("Bob")
	printMessage("Charlie")
	fmt.Println(swap("Alice", "Bob"))
	result := sum(5, 3)
	fmt.Println("Result sum :", result)
	fmt.Println(result)
	fmt.Println("-------")
	testcount(1)

	fmt.Println("-------")
	fmt.Println("Sum of 1, 2, 3, 4, 5 is:", sumAll(1, 2, 3, 4, 5))

}
