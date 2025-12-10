package main

import "fmt"

//This is comment
func main() {

	var (
		name     = "Jerald"
		lastname = "Montemor"
		age      = 22
	)
	//Can print the values using %v
	fmt.Printf("This is the value of the name %v: \n This is the value of lastname: %v : \n this is the value of the age: %v", name, lastname, age)
	fmt.Printf("Types: %T,%T,%T", name, lastname, age)
}
