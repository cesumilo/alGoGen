package shared

type SortedFitness struct {
	Idx int
	Score float64
}

type SortFitness []*SortedFitness

func (a SortFitness) Len() int {
	return len(a)
}

func (a SortFitness) Swap(i, j int) {
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp
}

func (a SortFitness) Less(i, j int) bool {
	return a[i].Score < a[j].Score
}
