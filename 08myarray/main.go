package main

import "fmt"

func main(){
	fmt.Println("Welcome to array in golangs")


	var fruitList [5]string

	fruitList[0]="Apple"
	fruitList[1]="Orange"
	fruitList[2]="Banana"
	fruitList[3]="Strawberry"
	fruitList[4]="Pineapple"


	fmt.Println("Fruit List is ",fruitList)
	fmt.Println("Fruit List is ",len(fruitList))
	var vegList =[5]string{"tomato","potato","beans"}
	fmt.Println("Veg list",vegList)
}
