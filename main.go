package main

import "fmt"

//This is comment
func main() {

	//With keyword datatype
	fmt.Println("***Multiple variable with initial value***")
	var a, b, c int = 1, 2, 3
	fmt.Println(a, b, c)

	//With string and bool
	fmt.Println("**************With string and bool*****************")

	var age, name, gay = 22, "Jerald", false
	fmt.Println(age, name, gay)

	fmt.Println("******************************************")

	r, d := 34, "Panpan"
	fmt.Println(r, d)

	//Variable declartion in a block

	var (
		cat1    string = "Panpan"
		cat1Age        = 3
		sound   string = "Meow"
	)
	fmt.Println(cat1, cat1Age, sound)

	fmt.Println("*****************PI Value************************")

	//This is the typed constant
	const piValue float32 = 3.14

	//Untyped constants
	const pivalue2 = 3.14

	fmt.Println(piValue)
	fmt.Println(pivalue2)
}
