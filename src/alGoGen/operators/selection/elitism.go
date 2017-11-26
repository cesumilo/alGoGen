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

type sortedFitness struct {
	idx int
	score float64
}

type SortFitness []*sortedFitness
func (a SortFitness) Len() int {
	return len(a)
}
func (a SortFitness) Swap(i, j int) {
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp
}
func (a SortFitness) Less(i, j int) bool {
	return a[i].score < a[j].score
}

type Elitism struct {
	fitness SortFitness
}

func (s *Elitism) Execute(individuals []*shared.Individual, scores []float64, size int) ([]*shared.Individual, error) {

	if len(individuals) != len(scores) {
		return nil, elitismError{fmt.Sprintf("Individuals and scores arrays should have the same size: %i != %i", len(individuals), len(scores))}
	}

	s.fitness = make(SortFitness, len(individuals), len(individuals))
	for i, v := range scores {
		s.fitness[i] = &sortedFitness{i, v}
	}
	sort.Sort(s.fitness)

	var selectedIndividuals []*shared.Individual
	i := 0
	for i < size {
		j := 0
		for j < len(s.fitness) && i < size {
			selectedIndividuals = append(selectedIndividuals, individuals[s.fitness[j].idx])
			i++
			j++
		}
	}

	return selectedIndividuals, nil
}
