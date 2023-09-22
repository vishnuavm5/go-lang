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
	fmt.Printf("Name is %v and email is %v\n", vishnu.Name, vishnu.Email)
	vishnu.GetStatus()
	vishnu.NewMail()

}

type User struct {
	Name   string
	Email  string
	status bool
	Age    int
}

// method syntax
func (u User) GetStatus() {
	fmt.Println("Is user active:", u.status)

}

func (u User) NewMail() { //this doesnt change the original value its only the copy of the value
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is ", u.Email)
}
