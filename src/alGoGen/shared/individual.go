package shared

type Genotype interface {}

type Individual interface {
	Init(id int)
	Id() int
	Fitness() float64
}

type Individuals []*Individual

func (a Individuals) Len() int {
	return len(a)
}

func (a Individuals) Swap(i, j int) {
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp
}

func (a Individuals) Less(i, j int) bool {
	return (*a[i]).Fitness() < (*a[j]).Fitness()
}
