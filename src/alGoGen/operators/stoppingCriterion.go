package operators

import (
	"alGoGen/shared"
)

type StoppingCriterion interface {
	Execute(int, shared.Individuals) bool
}
