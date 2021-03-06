package main

import "fmt"

func main() {

	customerNumber := make([]int, 3)
	// 3 is length & capacity
	// // length - number of elements referred to by the slice
	// // capacity - number of elements in the underlying array
	customerNumber[0] = 7
	customerNumber[1] = 10
	customerNumber[2] = 15

	fmt.Println(customerNumber[0])
	fmt.Println(customerNumber[1])
	fmt.Println(customerNumber[2])

	greeting := make([]string, 3, 3)
	// 3 is length - number of elements referred to by the slice 切片所引用的元素数
	// 5 is capacity - number of elements in the underlying array 基础数组中的元素数
	// you could also do it like this

	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "dias!"

	fmt.Println(greeting[2])
	fmt.Println(greeting, len(greeting))

}
