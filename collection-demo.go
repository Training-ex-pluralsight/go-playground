package main

import "fmt"

func main() {
	fmt.Println("Arrays")
	var array1 [3]int
	array1[0] = 1
	array1[1] = 2
	array1[2] = 3

	fmt.Println(array1)

	array2 := [3]int{0, 1, 2}
	fmt.Println(array2)

	fmt.Println("\nSlices")
	array3 := [3]int{0, 1, 2}

	slice := array3[:2]
	fmt.Println(slice)

	// new example
	slice2 := []int{1, 2, 3}
	fmt.Println(slice2)
	slice2 = append(slice2, 5, 6, 7, 8)
	fmt.Println(slice2)

	fmt.Println("\nMaps")
	map1 := map[string]int{"foo": 42}
	fmt.Println(map1)
	fmt.Println(map1["foo"])

	map1["foo"] = 24
	fmt.Println(map1["foo"])

	delete(map1, "foo")
	fmt.Println(map1)

	fmt.Println("\nStructs")
	type user struct {
		ID        int
		FirstName string
		LastName  string
	} // remember that this specific struct is scoped to function. They can be defined in package level too

	var u user
	fmt.Println("\tStruct initialized to zero-values of its parameters")
	fmt.Println(u)

	u.ID = 1234
	u.FirstName = "Fury"
	u.LastName = "Wojcik"
	fmt.Println("\tStructs with data")
	fmt.Println(u)

	u2 := user{ID: 4567,
		LastName:  "Ahern",
		FirstName: "David"} //don't forget to NOT put that curly brace
	fmt.Println(u2)

	u3 := user{ID: 4567,
		LastName:  "Something",
		FirstName: "Nemo",
	}
	fmt.Println(u3)
}
