package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race condition - lco.in")
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{} //also check RWMutex

	var score = []int{0}
	wg.Add(3)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {

		fmt.Println("one Routine")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()

	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Two Routine")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Three Routine")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)

}

//command to run the program in race condition < go run --race . >
