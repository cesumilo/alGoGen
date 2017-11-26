package operators

import (
	"alGoGen/shared"
)

type Mutation interface {
	Execute(*shared.Individual, int) (*shared.Individual, bool)
}
