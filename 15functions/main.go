package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is ", result)

	proResult, mymessage := proAdder(1, 2, 3, 4, 5, 6, 7)
	fmt.Println("ProResult is", proResult, mymessage)

}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}
	return total, "hello world"
}

func greeter() {
	fmt.Println("Namaste from golang")
}
