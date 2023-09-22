package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Switch and case in golang")

	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("Dice value is 2 move to spot 2")
	case 3:
		fmt.Println("Dice value is 3 move to spot 3")
	case 4:
		fmt.Println("Dice value is 4 move to spot 4")
	case 5:
		fmt.Println("Dice value is 5 move to spot 5")
	case 6:
		fmt.Println("Dice value is 6 roll again!")
	default:
		fmt.Println("what is what!")
	}

}
