package operators

import (
	"alGoGen/shared"
)

type Crossover interface {
	Execute(*shared.Individual, *shared.Individual, int) ([]*shared.Individual, bool)
}
