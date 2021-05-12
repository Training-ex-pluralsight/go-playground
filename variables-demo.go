package main

import "fmt"

func main() {

	fmt.Println("Primitive data types")
	var i = 43
	fmt.Println(i)

	var f float32 = 3.14
	fmt.Println(f)

	firstName := "Fury"
	fmt.Println(firstName)

	fmt.Println("\nBuilt-in complex number")
	c := complex(3, 4)
	fmt.Println(c)

	r, im := real(c), imag(c)
	fmt.Println(r, im)

	// pointer vars
	fmt.Println("\nPointer variables")

	var name *string = new(string)
	*name = "Fury"
	fmt.Println(*name)

	//address of operator demo
	nextName := "Dave"
	fmt.Println(nextName)

	ptr := &nextName
	fmt.Println(ptr, *ptr)

	nextName = "Clara"
	fmt.Println(ptr, *ptr)

	// constants
	fmt.Println("\nConstants")

	const num1 = 3 // implicitly typed constant
	fmt.Println(num1 + 3)

	fmt.Println(num1 + 1.2)

	const num2 int = 5 // explicitly typed constant
	fmt.Println(num2 + 3)

	fmt.Println(float32(num2) + 1.2)

}
