package operators

import (
	"alGoGen/shared"
)

type PopulationSelection interface {
	Execute([]*shared.Individual, []float64, int) ([]*shared.Individual, error)
}
