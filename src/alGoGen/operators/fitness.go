package operators

import (
	"alGoGen/shared"
)

type Fitness interface {
	Execute([]*shared.Individual) ([]float64, error)
}
