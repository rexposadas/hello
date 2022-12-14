package main

import (
	"fmt"
)

func main() {
	list := []string{"h", "e", "l", "l", "o"}

	message := make(chan string)

	// populate the channel with our message on rune at a time
	go func() {
		for _, r := range list {
			message <- r
		}
		close(message)
	}()

	// consume the items in the channel and print them
	for m := range message {
		fmt.Print(m)
	}
}
