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
		j := 0
		for j < len(individuals)  && i < size {
			selectedIndividuals = append(selectedIndividuals, individuals[j])
			i++
			j++
		}
	}

	return selectedIndividuals, nil
}
