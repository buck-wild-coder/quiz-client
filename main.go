package main

import (
	"fmt"
)

func main() {
	cache := Cache{
		ch: make(chan Quiz, 10),
	}

	go cache.buffering()
	for {
		answer := cache.askQuestion()
		input := read()
		if input == answer {
			fmt.Printf("Correct Answer\n")
		} else {
			fmt.Printf("Wrong answer\n")
		}
	}
}
