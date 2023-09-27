package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang -lco.in")
	myChannel := make(chan int, 2)

	wg := &sync.WaitGroup{}

	// myChannel <- 5

	// fmt.Println(<-myChannel)

	wg.Add(2)
	//receive only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		//fmt.Println(<-myChannel)
		//fmt.Println(<-myChannel)
		val, isChanelOpen := <-myChannel

		fmt.Println(isChanelOpen)
		fmt.Println(val)
		wg.Done()
	}(myChannel, wg)
	//send only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myChannel <- 5
		myChannel <- 6
		close(myChannel)
		wg.Done()
	}(myChannel, wg)

	wg.Wait()

}

//we can only dump value in to channel when there is listener
