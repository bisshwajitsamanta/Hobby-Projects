package main

import (
	"fmt"
	"sync"
)

func PrintSomething(s string, wg *sync.WaitGroup) {
	fmt.Println(s)
	defer wg.Done()
}
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go PrintSomething("Hello from a goroutine!", &wg)
	wg.Wait()
}
