package main

import "fmt"

func main() {

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages:", languages)
	fmt.Println("JS is short for:", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all languages:", languages)

	// loops
	for key, value := range languages {
		fmt.Printf("Key is %v, value is %v\n", key, value)
	}
}
