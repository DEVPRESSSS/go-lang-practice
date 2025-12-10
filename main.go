package main

import "fmt"

//This is comment
func main() {
	//Create a slice from an array
	myarray := [6]int{1, 2, 3, 4, 5, 6}
	myslice := myarray[0:3]
	fmt.Println(myslice)
	fmt.Println(cap(myslice))

	//Using make function
	fmt.Println("********Using make()**********")
	myslice1 := make([]int, 4, 6)
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))

}
