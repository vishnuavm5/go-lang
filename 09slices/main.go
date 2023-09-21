package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("This is on slices")
	var fruitList = []string{"Apple", "Banana", "Peach"}
	fmt.Printf("Type of fruitList %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Tomato")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	highScores := make([]int, 4)
	highScores[0] = 244
	highScores[1] = 944
	highScores[2] = 744
	highScores[3] = 544
	//highScores[4]=000  this doesnt work

	highScores = append(highScores, 666, 888, 999)
	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores)) // returns boolean

	//remove a value from slice
	var courses = []string{"react", "javascript", "swift", "python"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
