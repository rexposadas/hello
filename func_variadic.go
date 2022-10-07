package main

import "fmt"

func helloVar(say ...rune) {
	for _, s := range say {
		fmt.Print(string(s))
	}
	fmt.Println()
}

func main() {

	say := []rune{'h', 'e', 'l', 'l', 'o'}
	helloVar(say...)
}
