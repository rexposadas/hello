package main

import "fmt"

func main() {
	h := struct {
		Value string
	}{
		Value: "hello",
	}

	fmt.Println(h.Value)
}
