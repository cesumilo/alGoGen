package operators

import (
	"alGoGen/shared"
)

type PopulationSelection interface {
	Execute(shared.Individuals, int) (shared.Individuals, error)
}
