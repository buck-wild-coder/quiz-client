package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Cache struct {
	ch chan Quiz
}

func (c *Cache) fetch() (Quiz, error) {
	resp, err := http.Get("https://opentdb.com/api.php?amount=1")
	if err != nil {
		return Quiz{}, err
	}
	defer resp.Body.Close()

	var quiz Quiz
	binary, err := io.ReadAll(resp.Body)
	if err != nil {
		return Quiz{}, err
	}
	json.Unmarshal(binary, &quiz)
	return quiz, nil
}

func (c *Cache) buffering() {
	for {
		quiz, err := c.fetch()
		if err != nil {
			log.Print("FUcked", err)
		}
		fmt.Print(len(c.ch))
		c.ch <- quiz
	}
}
