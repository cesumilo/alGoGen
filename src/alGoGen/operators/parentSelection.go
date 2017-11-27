package operators

import (
	"alGoGen/shared"
)

type ParentSelection interface {
	Execute(shared.Individuals, int) (shared.Individuals, error)
}
