package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")
	languagues := make(map[string]string)
	languagues["JS"] = "javascript"
	languagues["RB"] = "Ruby"
	languagues["PY"] = "Python"

	fmt.Println("List of all languages", languagues)
	fmt.Println("JS shorts for", languagues["JS"])
	delete(languagues, "RB")
	fmt.Println("List of all languages", languagues)

	//loops

	for key, value := range languagues {
		fmt.Printf("For key %v,value is %v\n", key, value)
	}
}
