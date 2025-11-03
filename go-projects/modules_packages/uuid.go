package main

import "github.com/pborman/uuid"

func main() {
	uid := uuid.NewRandom()
	println(uid.String())
}
