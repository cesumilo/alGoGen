package operators

import (
	"alGoGen/shared"
)

type Fitness interface {
	Execute(shared.Individuals) ([]float64, error)
}
