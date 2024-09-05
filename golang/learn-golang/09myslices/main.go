package main

import (
	"fmt"
)

func main() {

	// fmt.Println("Welcome to Slices")

	// var fruitList = []string{"Apple", "Tomato", "Peach"}
	// fmt.Printf("Type of fruitlist is %T\n", fruitList)

	// fruitList = append(fruitList, "Mango", "Banana")
	// fmt.Println(fruitList)

	// fruitList = fruitList[:3]
	// fmt.Println(fruitList)

	// highScores := make([]int, 4)
	// highScores[0] = 234
	// highScores[1] = 945
	// highScores[2] = 465
	// highScores[3] = 867
	// //highScores[4] = 777

	// highScores = append(highScores, 555, 666, 777, 321)
	// fmt.Println(highScores)

	// //sorting
	// sort.Ints(highScores)
	// fmt.Println(highScores)
	// fmt.Println(sort.IntsAreSorted(highScores))

	// remove value from slice based on index

	courses := []string{"Python", "Javascript", "C#", "C++", "Ruby"}
	fmt.Println(courses)

	index := 2

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
