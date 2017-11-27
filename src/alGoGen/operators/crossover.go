package operators

import (
	"alGoGen/shared"
)

type Crossover interface {
	Execute(*shared.Individual, *shared.Individual) (shared.Individuals, error)
	NumberOfOffspring() int
}
