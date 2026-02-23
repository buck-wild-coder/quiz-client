package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func fetch() (Quiz, error) {
	resp, err := http.Get("https://opentdb.com/api.php?amount=1")
	if err != nil {
		return Quiz{}, err
	}
	defer resp.Body.Close()

	var quiz Quiz
	binary, err := io.ReadAll(resp.Body)
	if err != nil {
		return quiz, err
	}
	json.Unmarshal(binary, &quiz)
	return quiz, nil
}
