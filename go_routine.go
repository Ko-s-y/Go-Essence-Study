package main

import (
	"fmt"
	"time"
)

func main() {
	message := "hi"

	go func() {
		sendMessage(message)
	}()
	message = "ho"

	time.Sleep(time.Second)
	println(message)
	time.Sleep(time.Second)
}

func sendMessage(msg string) {
	fmt.Println(msg)
}
