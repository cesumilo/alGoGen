package operators

import (
	"alGoGen/shared"
)

type PopulationSelection interface {
	Execute(shared.Individuals, []float64, int) (shared.Individuals, error)
}
