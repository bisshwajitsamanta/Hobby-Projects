package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan string)
	var wg sync.WaitGroup
	start := time.Now()
	wg.Add(2)
	go buy(ch, &wg)
	go sell(ch, &wg)
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func buy(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- "Buy Furniture from the Channel"
}
func sell(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Waiting for Data")
	val := <-ch
	fmt.Println("Got Data from Channel - ", val)
	close(ch)
}
