package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj52jsn"

func main() {
	fmt.Println("Url in golang")
	fmt.Println(myurl)

	//parsing
	result, _ := url.Parse(myurl)

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	qparams := result.Query()

	fmt.Printf("the type of query parms %T\n", qparams)
	fmt.Println(qparams["coursename"])
	for _, value := range qparams {
		fmt.Println("Param is :", value)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=vishnu",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)

}
