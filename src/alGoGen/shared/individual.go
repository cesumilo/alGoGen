package shared

type Genotype interface {}

type Individual interface {
	Init(id int)
	Id() int
	GetGenotype() Genotype
}
