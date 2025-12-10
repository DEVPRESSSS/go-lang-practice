package main

import (
	"fmt"
)

func addition(x int, c int) int {

	return x + c

}

func subtraction(x int, c int) (result int) {
	result = x - c

	return

}

func multiplication(x int, c int) (result int, message string) {

	result = x * c
	message = "This is return type string"
	return
}

// This is comment
func main() {

	fmt.Println(addition(1, 2))
	fmt.Println(subtraction(10, 4))
	fmt.Println(multiplication(3, 3))
}
