package operators

import "alGoGen/shared"

type Statistics interface {
	Compute(shared.Individuals)
}
