package main

import (
	"fmt"
	"sync"
)

var (
	message   string
	waitGroup sync.WaitGroup
)

func PrintSomething(s string, m *sync.Mutex) {
	m.Lock()
	message = s
	m.Unlock()
	defer waitGroup.Done()
}

func main() {
	message = "Hello from a goroutine with Mutex!"
	var mutex sync.Mutex
	waitGroup.Add(3)
	go PrintSomething("Hello Universe!", &mutex)
	go PrintSomething("Hello Cosmos!", &mutex)
	go PrintSomething("Hello Milky way!", &mutex)
	waitGroup.Wait()
	fmt.Println(message)
}
