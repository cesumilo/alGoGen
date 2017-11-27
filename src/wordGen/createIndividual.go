package main

import "alGoGen/shared"

type CreateWord struct {
	wordLength int
}

func (c *CreateWord) Execute(id int) (*shared.Individual, error) {
	var word Word
	word.Init(id, c.wordLength)
	var idv shared.Individual
	idv = &word
	return &idv, nil
}
