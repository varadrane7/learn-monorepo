package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays in Golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("Fruit list is", fruitList)
	fmt.Println("List length is", len(fruitList))

	var vegList = [4]string{"Potato", "Beans", "Mushroom", "Basil"}

	fmt.Println("Veg list is:", vegList)
	fmt.Println("Veg list length is:", len(vegList))
}
