package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb in golang")
	//performGetRequest()
	//performPostJsonRequest()
	performPostFormRequest()

}

func performGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("status code :", response.StatusCode)

	fmt.Println("content length is:", response.ContentLength)

	var responseString strings.Builder

	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content) //this writes to responseString and we can use it .string() method to get the data as string
	fmt.Println("Byte count is:", byteCount)
	fmt.Println(responseString.String())

	fmt.Println(string(content))
}
func performPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	//fake payload

	requestBody := strings.NewReader(`
	{
		"coursename":"lets go with golang",
		"price":0,
		"platform":"learncodeOnline.in"
	}
	`)
	response, _ := http.Post(myurl, "application/json", requestBody)
	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
	defer response.Body.Close()
}

func performPostFormRequest() {
	const myurl = "http://localhost:8000/postform"
	//formdata

	data := url.Values{}
	data.Add("firstname", "Vishnu")
	data.Add("lastname", "Arredhula")
	data.Add("email", "vishnuavm5@gmail.com")

	response, _ := http.PostForm(myurl, data)

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
	defer response.Body.Close()

}
