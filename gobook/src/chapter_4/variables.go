package main

import "fmt"

// this is a comment

var wow = "Multiple Functions have Access to this variable!"

const konstante = 42

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
	fahrenheitToCelsius(13)
	fettToMeters(13)
}

func printWoW() {
	fmt.Println(wow)
}

func fahrenheitToCelsius(value float64) {
	value = ((value - 32) * 5 / 9)
	fmt.Println(value)
}

func fettToMeters(value float64) {
	value = (value * 0.3048)
	fmt.Println(value)
}
