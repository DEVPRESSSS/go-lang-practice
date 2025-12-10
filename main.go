package main

import "fmt"

//This is comment
func main() {

	//Multilple variable
	fmt.Println("***Multiple variable with initial value***")
	var a, b, c int = 1, 2, 3
	fmt.Println(a, b, c)

	//With string and bool
	fmt.Println("**************With string and bool*****************")

	var age, name, gay = 22, "Jerald", false
	fmt.Println(age, name, gay)

	fmt.Println("******************************************")

	r ,d := 34, "Panpan"
	fmt.Println(r,d)

}
