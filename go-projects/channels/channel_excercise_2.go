package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	// Create first goroutine here
	go printNumbers(&wg)
	// Create second goroutine here
	go printLetters(&wg)

	// Wait for both goroutines to finish here
	wg.Wait()
}

func printNumbers(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range 5 {
		fmt.Println(i)
	}

}

func printLetters(wg *sync.WaitGroup) {
	defer wg.Done()
	for ch := 'a'; ch < 'a'+5; ch++ {
		fmt.Println(string(ch))
	}

}
