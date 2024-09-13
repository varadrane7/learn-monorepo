package main

import "fmt"

func main() {
	fmt.Println("If else in golang")

	loginCount := 43
	var result string

	if loginCount < 42 {
		result = "Regular User"
	} else if loginCount > 42 {
		result = "More than regular user"
	} else {
		result = "You are a GOD"
	}
	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is NOT less than 10")
	}
}
