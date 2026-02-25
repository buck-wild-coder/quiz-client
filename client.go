package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type Cache struct {
	ch      chan Quiz
	timeout int
}

func (c *Cache) fetch() (Quiz, error) {
	resp, err := http.Get("https://opentdb.com/api.php?amount=1")
	if err != nil {
		time.Sleep(200 * time.Millisecond)
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
			log.Print(err)
			c.shutdown()
			continue
		}
		if len(quiz.Results) == 0 {
			continue
		}
		c.timeout = 0
		c.ch <- quiz
	}
}

func (c *Cache) shutdown() {
	c.timeout += 1
	if c.timeout > 10 {
		log.Fatal("Timeout. Server is not responding...")
	}
}
