package selection

import (
	"alGoGen/shared"
	"fmt"
	"math/rand"
	"time"
)

type rouletteWheelError struct {
	message string
}
func (err rouletteWheelError) Error() string {
	return err.message
}

type RouletteWheel struct {
}

func (s *RouletteWheel) Execute(individuals shared.Individuals, scores []float64, size int) (shared.Individuals, error) {

	var selectedIndividuals shared.Individuals

	if len(individuals) != len(scores) {
		return nil, rouletteWheelError{fmt.Sprintf("Individuals and scores arrays should have the same size: %i != %i", len(individuals), len(scores))}
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	totalFitness := float64(0)
	for _, v := range scores {
		totalFitness += v
	}

	for i := 0; i < size; i++ {
		value := r.Float64() * totalFitness
		j := 0
		sum := float64(0)
		for j < len(scores) && sum < value {
			j++
			sum += scores[j]
		}
		selectedIndividuals = append(selectedIndividuals, individuals[j])
	}

	return selectedIndividuals, nil
}
