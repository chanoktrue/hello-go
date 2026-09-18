package main

import "fmt"

func zeroValue(ivalue int) {
	ivalue = 0
}

func zeroPointer(ipointer *int) {
	*ipointer = 0
}

func main() {
	i := 1
	fmt.Println("i: ", i)

	zeroValue(i)
	fmt.Println("i from function zeroValue: ", i)

	zeroPointer(&i)
	fmt.Println("i value function zeroPointer: ", i)
	fmt.Println("i address function zeroPointer: ", &i)
}

/*
&i // ขอ address ของตัวแปร i
*i // อ่านหรือแก้ค่าที่ address นั้น
*/
