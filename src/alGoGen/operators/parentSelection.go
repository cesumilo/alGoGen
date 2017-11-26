package operators

import (
	"alGoGen/shared"
)

type ParentSelection interface {
	Execute(shared.Individuals, []float64, int) (shared.Individuals, error)
}
