package main

import (
	"alGoGen/shared"
	"math/rand"
)

type MutationError struct {
	message string
}
func (err MutationError) Error() string {
	return err.message
}

type Mutation struct {
}

func (m *Mutation) Execute(individual *shared.Individual) (*shared.Individual, error) {
	word, ok := (*individual).(*Word)
	if !ok {
		return nil, MutationError{"Invalid individual type!"}
	}
	genotype := []byte(word.GetGenotype())
	newByte := TargetWord[rand.Intn(len(TargetWord))]
	genotype[rand.Intn(len(genotype))] = newByte
	word.genotype = string(genotype)

	var idv shared.Individual
	idv = word

	return &idv, nil
}
