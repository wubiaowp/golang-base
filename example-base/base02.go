package main

import (
	"time"
	"fmt"
)

func say(report string) {

	for i := 0; i < 5; i++ {
		time.Sleep(10 * time.Second)
		fmt.Print(report)
	}
}

func main() {
	go say("hello")
	say("world")
}
