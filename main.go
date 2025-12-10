package main

import (
	"fmt"
)

// This is comment
func main() {

	//Access the values using index of the slice
	prices := []float32{100, 400.00, 30.00}
	prices[1] = 20.00
	fmt.Println((prices[1]))

	//Append function
	fmt.Println("********Using append function to insert new value********")
	prices = append(prices, 500, 900)
	fmt.Println(prices)

	//Append one slice to another one
	fmt.Println("********Using append function to append another slice********")
	my_slice_1 := []int{1, 2, 3, 4, 5}
	my_slice_2 := []int{6, 7, 8, 9, 10}
	my_slice_1 = append(my_slice_1, my_slice_2...)
	fmt.Println(my_slice_1)

}
