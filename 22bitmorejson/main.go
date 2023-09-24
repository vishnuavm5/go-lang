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
	EncodeJson()
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
