package main

import "fmt"

// this is a comment

var wow = "Multiple Functions have Access to this variable!"

func main() {
	var x string = "Hello World"
	var y string
	y = "bla!"
	fmt.Println(x)
	fmt.Println(y + " + second")

	fmt.Println("true =", !(x == y))

	z := 1
	z += z
	fmt.Println(z)

	printWoW()
}

func printWoW() {
	fmt.Println(wow)
}
