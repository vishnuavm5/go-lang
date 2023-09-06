package main

import "fmt"

//making the first letter capital makes the variable or constant public

const LoginToken string="vsusuwsnuwshwns"

func main(){
	var username string="Vishnu"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n",username)

	var value uint8=56
	fmt.Println(value)
	fmt.Printf("Variable is of type: %T\n",value)

	var isTrue=false
	fmt.Println(isTrue);
	fmt.Printf("Variable is of type: %T\n",isTrue)

	var smallFloat float32=56.937378383
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T\n",smallFloat)

	var anotherVariable int

	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T\n",anotherVariable)


	var sampleString string
	fmt.Printf(sampleString)
	fmt.Printf("Variable is of type: %T\n",sampleString)

	//implicit types

	var name="Vishnu"
	fmt.Println(name)

	//no var style
	numberOfUser :=300000
	fmt.Println(numberOfUser);

	fmt.Println(LoginToken);
	fmt.Printf("Variable is of type: %T\n",LoginToken)


}