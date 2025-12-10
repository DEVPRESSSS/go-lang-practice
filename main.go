package main

import "fmt"

//This is comment
func main() {

	//Create slice with [] datatype {values}
	first_slice := []string{"Hello"}
	fmt.Println(first_slice)

	//Get the length and capacity of the slice
	fmt.Println(len(first_slice))
	fmt.Println(cap(first_slice))

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)
}
