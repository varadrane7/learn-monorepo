package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers")

	// var ptr *int
	// fmt.Println("Value of pointer is", ptr)

	myNumber := 23

	ptr := &myNumber

	fmt.Println("Value of actual pointer is", ptr)
	fmt.Println("Value stored at pointer is", *ptr)

	*ptr *= 2
	fmt.Println("New value is", myNumber)

}
