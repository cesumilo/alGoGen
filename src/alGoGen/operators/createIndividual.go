package operators

import (
	"alGoGen/shared"
)

type CreateIndividual interface {
	Execute(int) (*shared.Individual, error)
}
