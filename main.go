package main

import (
	"fmt"
)

func main() {
	cache := Cache{
		ch:    make(chan Quiz, 10),
		errCh: make(chan error),
	}

	go cache.fetch()
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
