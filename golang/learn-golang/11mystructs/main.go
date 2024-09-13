package main

import "fmt"

func main() {
	// No inheritance in golang

	ghostNinja := User{"Ghost", "ghost@go.dev", true, 19}
	fmt.Println(ghostNinja)
	fmt.Printf("Details are: %+v\n", ghostNinja)
	fmt.Printf("Name is %v, Age is %v", ghostNinja.Name, ghostNinja.Age)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
