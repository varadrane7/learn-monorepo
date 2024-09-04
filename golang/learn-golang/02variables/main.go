package main

import "fmt"

// public declaration
const LoginToken string = "abcdefgh"

func main() {
	var username string = "checker"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	isLoggeedIn := false
	fmt.Println(isLoggeedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggeedIn)

	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("Variable is of type: %T \n", smallValue)

	var smallFLoat float32 = 255.45545454445
	fmt.Println(smallFLoat)
	fmt.Printf("Variable is of type: %T \n", smallFLoat)

	// default values
	var anotherVariable int
	fmt.Println(anotherVariable)

	// implicit type
	var website = "github.com"
	fmt.Println(website)

	// no var style
	numberOfUsers := 30000
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
