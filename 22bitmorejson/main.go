package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` //alias
	Price    int    `json:"price"`
	Platform string
	Password string   `json:"-"`              //dash doesnt show anything
	Tags     []string `json:"tags,omitempty"` //removes nulls in data
}

func main() {
	fmt.Println("json in golang")
	//EncodeJson()
	decodeJson()
}

func EncodeJson() {

	lcocourses := []course{
		{"ReactJS Bootcamp", 299, "lco.in", "sauwhh", []string{"webdev", "js"}},
		{"MERN Bootcamp", 199, "lco.in", "ujsj", []string{"fullstack", "js"}},
		{"Angular Bootcamp", 399, "lco.in", "saupsiksjswhh", nil},
	}

	//package this data as json

	finalJson, err := json.MarshalIndent(lcocourses, "", "\t") //it takes interface name,prefix and indent

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func decodeJson() {

	jsonDataFromWeb := []byte(`
	{
		"coursename":"Reactjs Bootcamp",
		"price":299,
		"website":"lco.dev",
		"tags":["webdev","js"]
	}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse) //to print interface

	} else {
		fmt.Println("JSON is not valid")
	}

	//some cases where you just want to add data to key value

	var myOnlineData map[string]interface {
	}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData) //to print interface

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is :%T\n", k, v, v)
	}

}
