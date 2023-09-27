package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup //pointers usually
var mut sync.Mutex    //pointers usally

func main() {
	fmt.Println("This is for goroutines in golang")
	// go greeter("hello")
	// greeter("World")

	websitelist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://github.com",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1) //
	}
	fmt.Println(signals)
	wg.Wait() // wait till the threads are back

}

// func greeter(s string) {
// 	// for i := 0; i < 6; i++ {
// 	// 	time.Sleep(3 * time.Millisecond)
// 	// 	fmt.Println(s)

// 	// }
// }

func getStatusCode(endpoint string) {
	defer wg.Done() //send singal done
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("Oops in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)

	}
}
