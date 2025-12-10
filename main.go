package main

import "fmt"

//This is comment
func main() {

	//Two types of declaring an array
	//length is defined here
	var array1 = [3]string{"Jerald", "Montemor", "BSCSJIRO"}
	fmt.Println(array1)
	//Second way
	var array2 = [...]string{"Jerald", "Montemor", "BSCSJIRO"}
	fmt.Println(array2)

	array3 := [4]string{"Jerald", "Montemor", "BSCSJIRO"}
	fmt.Println((array3))

	//Change the index 1 "Montemor" to "Panpan"
	array3[1] = "Panpan"
	fmt.Println(array3)

	//Array initialization
	var cars = [3]string{"Volvo", "Honda"}
	var cars1 = [3]string{}
	var cars2 = [3]string{"Volvo", "Honda", "Test"}

	fmt.Println(cars, cars1, cars2)

	//Iniliaze only specific index value

	agents := [5]string{1: "SANTIAGO5", 4: "SANTIAGO6"}
	fmt.Println(agents)

}
