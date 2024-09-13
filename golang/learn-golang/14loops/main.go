package main

import "fmt"

func main() {
	fmt.Println("Loops in golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Satuday"}
	fmt.Println(days)

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for _, day := range days {
		fmt.Printf("Value is %v\n", day)
	}

	rogueValue := 1

	for rogueValue < 10 {

		if rogueValue == 7 {
			goto lco
		}

		if rogueValue == 5 {
			rogueValue++
			continue
		}
		fmt.Println("Value is:", rogueValue)
		rogueValue++
	}

lco:
	fmt.Println("Jumping at LCO")
}
