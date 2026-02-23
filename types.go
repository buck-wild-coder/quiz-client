package main

type Quiz struct {
	Results []results `json:"results"`
}

type results struct {
	Question          string   `json:"question"`
	Correct_answer    string   `json:"correct_answer"`
	Incorrect_answers []string `json:"incorrect_answers"`
}
