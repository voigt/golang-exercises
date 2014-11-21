package main

import "fmt"

// this is a comment

func main() {
	//For
	for i := 0; i < 10; i++ {
		fmt.Printf("Count %d\n", i)
	}
	//if
	for i := 1; i <= 5; i++ {
		//switch
		switch i {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		case 4:
			fmt.Println("Four")
		case 5:
			fmt.Println("Five")
		default:
			fmt.Println("Unknown Number")
		}
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}

}
