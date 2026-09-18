package main

func main() {

	var arr [5]int = [5]int{1, 2, 3, 4, 5}

	println("array length:", len(arr))

	for _, v := range arr {
		println(v)
	}
}
