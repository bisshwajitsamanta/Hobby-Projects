package main

import "testing"

func Test_updateMessage(t *testing.T) {
	msg = "Hello, World!"
	wg.Add(1)
	go updateMessage("GoodBye, cruel World!")
	wg.Wait()
	if msg != "GoodBye, cruel World!" {
		t.Errorf("Expected message to be 'GoodBye, cruel World!', but got '%s'", msg)
	}
}
