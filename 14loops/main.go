package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	day := []string{"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}
	fmt.Println(day)

	// for d := 0; d < len(day); d++ {
	// 	fmt.Println(day[d])
	// }

	// for i := range day { //in i the index is stored not values
	// 	fmt.Println(day[i])
	// }

	for index, day := range day {
		fmt.Printf("index is %v, value is %v\n", index, day)
	}

	for _, day := range day {
		fmt.Printf("value is %v\n", day)
	}

	rogueValue := 1

	for rogueValue < 10 {
		if rogueValue == 2 {
			goto medo
		}
		if rogueValue == 5 {
			rogueValue++
			continue
		}
		fmt.Println("Value is :", rogueValue)
		rogueValue++
	}
medo:
	fmt.Println("I am in medo")
}
