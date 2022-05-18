package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels")

	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}
	// myCh <- 5           //coming in to channel
	// fmt.Println(<-myCh) //going out from channel
	wg.Add(2)
	// Send only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		fmt.Println(<-myCh)
		fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)
	// Recieve only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 5
		myCh <- 6
		close(myCh)
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
