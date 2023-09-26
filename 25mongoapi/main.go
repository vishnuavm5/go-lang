package main

import (
	"fmt"
	"mongoapi/router"

	"net/http"
)

func main() {
	fmt.Println("Mongo API")
	r := router.Router()
	fmt.Println("Server is getting started...")
	http.ListenAndServe(":4000", r)
	fmt.Println("Listening at port 4000 ...")
}
