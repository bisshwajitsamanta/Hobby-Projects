package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(3)
	jobs := make(chan int)
	results := make(chan int)

	go func() {
		for j := range jobs {
			fmt.Println("received job", j)
			result := j * 2
			results <- result
		}
		defer wg.Done()
		close(results)
	}()

	go func() {
		for j := 1; j <= 3; j++ {
			jobs <- j
			fmt.Println("sent job", j)
		}
		defer wg.Done()
		close(jobs)

	}()

	go func() {
		for r := range results {
			fmt.Println("received result", r)
		}
		wg.Done()

	}()

	wg.Wait()
}
