package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type Cache struct {
	ch    chan Quiz
	errCh chan error
}

func (c *Cache) fetch() {
	resp, err := http.Get("https://opentdb.com/api.php?amount=1")
	if err != nil {
		c.ch <- Quiz{}
		c.errCh <- err
	}
	defer resp.Body.Close()

	var quiz Quiz
	binary, err := io.ReadAll(resp.Body)
	if err != nil {
		c.ch <- Quiz{}
		c.errCh <- err
	}
	json.Unmarshal(binary, &quiz)
	c.ch <- quiz
	c.errCh <- nil
	c.fetch()
}
