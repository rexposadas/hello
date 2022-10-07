package main

import "fmt"

type Hello struct {
	Value string
}

func main() {

	h := Hello{
		Value: "hello",
	}

	fmt.Println(h.Value)

}
