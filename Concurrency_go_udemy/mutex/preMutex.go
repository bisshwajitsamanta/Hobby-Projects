package main

import (
	"fmt"
	"sync"
)

var (
	msg string
	wg  sync.WaitGroup
)

func updateMessage(s string) {
	msg = s
	defer wg.Done()
}
func main() {
	msg = "Hello, World!"
	wg.Add(3)
	go updateMessage("Hello from goroutine 1")
	go updateMessage("Hello from goroutine 2")
	go updateMessage("Hello from goroutine 3")
	wg.Wait()
	fmt.Println(msg)
}
