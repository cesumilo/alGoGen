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
}

func (s *Elitism) Execute(individuals shared.Individuals, scores []float64, size int) (shared.Individuals, error) {

	if len(individuals) != len(scores) {
		return nil, elitismError{fmt.Sprintf("Individuals and scores arrays should have the same size: %i != %i", len(individuals), len(scores))}
	}

	sort.Sort(individuals)

	var selectedIndividuals shared.Individuals
	i := 0
	for i < size {
		j := len(individuals) - 1
		for j >= 0  && i < size {
			selectedIndividuals = append(selectedIndividuals, individuals[j])
			i++
			j--
		}
	}

	return selectedIndividuals, nil
}
