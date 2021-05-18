package main

import "fmt"

const pi = 3.1415
const (
	first  = 1
	second = "second"
)

const (
	expr1 = iota + 6
	expr2 = 2 << iota
)

const (
	expr3 = iota + 5
	expr4
)

func main() {
	// iota & constant expressions
	fmt.Println("iota & constant expressions")

	fmt.Println(first, second)

	fmt.Println("\nIota example")
	fmt.Println(expr1, expr2)

	fmt.Println(expr3, expr4)
}
