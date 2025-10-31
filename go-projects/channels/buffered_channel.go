package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan string, 3)

	wg.Add(2)
	go sendBufferedChannel(ch, &wg)
	wg.Wait()

}

func sendBufferedChannel(ch chan string, wg *sync.WaitGroup) {

	defer wg.Done()
	ch <- "Apple"
	ch <- "Banana"
	ch <- "Mango"
	go receiveBufferedChannel(ch, wg)
	close(ch)
}

func receiveBufferedChannel(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Waiting for Data")
	for i := range ch {
		fmt.Println("Received Data:", i)
	}

}
