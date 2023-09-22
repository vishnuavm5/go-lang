package main

import (
	"fmt"
)

func main() {

	//there is no inheritence in structs in golang
	fmt.Println("This is structs in golang")
	vishnu := User{"Vishnu", "vishnuavm5@gmail.com", true, 24}
	fmt.Println(vishnu)
	fmt.Printf("vishnu details are:%+v\n", vishnu)
	fmt.Printf("Name is %v and email is %v", vishnu.Name, vishnu.Email)

}

type User struct {
	Name   string
	Email  string
	status bool
	Age    int
}
