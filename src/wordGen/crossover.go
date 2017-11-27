package main

import (
	"alGoGen/shared"
	"math/rand"
)

type CrossoverError struct {
	message string
}
func (err CrossoverError) Error() string {
	return err.message
}

type Crossover struct {
}

func (c *Crossover)Execute(a *shared.Individual, b *shared.Individual) (shared.Individuals, error) {
	var newGenotype []byte
	var newIndividual Word
	var individuals []*shared.Individual

	wordA, okA := (*a).(*Word)
	wordB, okB := (*b).(*Word)

	if !okA || !okB {
		return nil, CrossoverError{"Invalid individual type!"}
	}

	gA := []byte(wordA.GetGenotype())
	gB := []byte(wordB.GetGenotype())
	for i := 0; i < len(gA); i++ {
		p := rand.Float32()
		if p < 0.5 {
			newGenotype = append(newGenotype, gA[i])
		} else {
			newGenotype = append(newGenotype, gB[i])
		}
	}

	var idv shared.Individual
	idv = &newIndividual

	newIndividual.genotype = string(newGenotype)
	individuals = append(individuals, &idv)
	return individuals, nil
}

func (c *Crossover) NumberOfOffspring() int {
	return 1
}
