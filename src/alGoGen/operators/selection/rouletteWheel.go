package selection

import (
	"alGoGen/shared"
	"math/rand"
	"time"
)

type RouletteWheel struct {
}

func (s *RouletteWheel) Execute(individuals shared.Individuals, size int) (shared.Individuals, error) {

	var selectedIndividuals shared.Individuals

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	totalFitness := float64(0)
	for _, v := range individuals {
		totalFitness += (*v).Fitness()
	}

	for i := 0; i < size; i++ {
		value := r.Float64() * totalFitness
		j := 0
		sum := float64(0)
		for j < len(individuals) && sum < value {
			j++
			sum += (*individuals[j]).Fitness()
		}
		selectedIndividuals = append(selectedIndividuals, individuals[j])
	}

	return selectedIndividuals, nil
}
