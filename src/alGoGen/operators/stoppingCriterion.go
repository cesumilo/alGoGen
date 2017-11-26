package operators

type StoppingCriterion interface {
	Execute(int, []float64) bool
}
