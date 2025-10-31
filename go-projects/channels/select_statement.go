package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	var wg sync.WaitGroup
	wg.Add(1)
	go One(ch1, &wg)
	go Two(ch2, &wg)

	select {
	case data1 := <-ch1:
		println(data1)
	case data2 := <-ch2:
		println(data2)
	default:
		fmt.Println("Executed Default Block")
	}
	wg.Wait()
}

func Two(ch2 chan string, wg *sync.WaitGroup) {
	ch2 <- "Data Received from Channel Two"
	defer wg.Done()
}

func One(ch1 chan string, wg *sync.WaitGroup) {
	ch1 <- "Data Received from Channel One"
	defer wg.Done()
}
