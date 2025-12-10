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

// This is comment
func main() {

	fmt.Println(addition(1, 2))
	fmt.Println(subtraction(10, 4))
}
