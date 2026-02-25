package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func (c *Cache) askQuestion() string {
	for {
		data := <-c.ch
		fmt.Print(len(data.Results))
		answer, err := printFormat(data)
		if err != nil {
			log.Print(err)
			continue
		}
		return answer
	}
}

func printFormat(decoded Quiz) (string, error) {
	if len(decoded.Results) == 0 {
		return "", errors.New("Server did not responded, trying again...")
	}
	choices := decoded.Results[0].Incorrect_answers
	answer := decoded.Results[0].Correct_answer
	choices = append(choices, answer)

	fmt.Println(decoded.Results[0].Question)
	for _, item := range choices {
		fmt.Println(item)
	}
	return strings.ToLower(answer), nil
}

func read() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	return strings.ToLower(line)
}
