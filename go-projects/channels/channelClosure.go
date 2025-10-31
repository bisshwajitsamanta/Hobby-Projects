package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	n := 10
	wg.Add(n)
	for i := range n {
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
}
