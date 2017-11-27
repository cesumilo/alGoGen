package operators

import (
	"alGoGen/shared"
)

type Mutation interface {
	Execute(*shared.Individual) (*shared.Individual, error)
}
