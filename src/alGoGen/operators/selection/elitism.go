package selection

import (
	"fmt"
	"alGoGen/shared"
	"sort"
)

type elitismError struct {
	message string
}
func (err elitismError) Error() string {
	return err.message
}

type Elitism struct {
	fitness shared.SortFitness
}

func (s *Elitism) Execute(individuals []*shared.Individual, scores []float64, size int) ([]*shared.Individual, error) {

	if len(individuals) != len(scores) {
		return nil, elitismError{fmt.Sprintf("Individuals and scores arrays should have the same size: %i != %i", len(individuals), len(scores))}
	}

	s.fitness = make(shared.SortFitness, len(individuals), len(individuals))
	for i, v := range scores {
		s.fitness[i] = &shared.SortedFitness{i, v}
	}
	sort.Sort(s.fitness)

	var selectedIndividuals []*shared.Individual
	i := 0
	for i < size {
		j := len(s.fitness) - 1
		for j >= 0  && i < size {
			selectedIndividuals = append(selectedIndividuals, individuals[s.fitness[j].Idx])
			i++
			j--
		}
	}

	return selectedIndividuals, nil
}
