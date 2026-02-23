package main

import (
	"fmt"
)

func main() {
	for {
		answer := askQuestion()
		input := read()
		if input == answer {
			fmt.Printf("Correct Answer\n")
		} else {
			fmt.Printf("Wrong answer\n")
		}
	}
}
