package operators

import (
	"alGoGen/shared"
)

type ParentSelection interface {
	Execute([]*shared.Individual, []float64, int) ([]*shared.Individual, error)
}
