package operators

import (
	"alGoGen/shared"
)

type Fitness interface {
	Execute(shared.Individuals) (error)
}
