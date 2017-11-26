package selection

import (
	"alGoGen/shared"
	"fmt"
	"time"
	"math/rand"
	"sort"
)

type tournamentError struct {
	message string
}
func (err tournamentError) Error() string {
	return err.message
}

type Tournament struct {
	k int
}

func (s* Tournament) Init(k int) {
	s.k = k
}

func (s *Tournament) Execute(individuals shared.Individuals, scores []float64, size int) (shared.Individuals, error) {

	var selectedIndividuals shared.Individuals

	if len(individuals) != len(scores) {
		return nil, tournamentError{fmt.Sprintf("Individuals and scores arrays should have the same size: %i != %i", len(individuals), len(scores))}
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < size; i++ {
		var pool shared.Individuals
		for j := 0; j < s.k; j++ {
			idx := r.Intn(len(individuals))
			pool = append(pool, individuals[idx])
		}
		sort.Sort(pool)
		selectedIndividuals = append(selectedIndividuals, pool[s.k - 1])
	}

	return selectedIndividuals, nil
}
