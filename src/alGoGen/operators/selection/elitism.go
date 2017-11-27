package selection

import (
	"alGoGen/shared"
	"sort"
)

type Elitism struct {
}

func (s *Elitism) Execute(individuals shared.Individuals, size int) (shared.Individuals, error) {

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
