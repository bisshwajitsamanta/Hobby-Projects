package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)

	go SendValue(ch1)
	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Select timeout")
	}
}

func SendValue(ch1 chan int) {
	//time.Sleep(3 * time.Second)
	ch1 <- 10
}
