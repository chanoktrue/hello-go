package main

import "fmt"

type MyInterface interface {
	calc_sum()
}

type MyStruct struct {
	a, b int
}

func (m MyStruct) calc_sum() int {
	sum := m.a + m.b
	return sum
}

func calc_sum2(m MyStruct) int {
	sum := m.a + m.b
	return sum
}

func main() {
	x := MyStruct{a: 10, b: 20}
	result := x.calc_sum()
	fmt.Println("Sum:", result)
}
