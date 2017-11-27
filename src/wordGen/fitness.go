package main

import (
	"alGoGen/shared"
	"math"
)

type FitnessError struct {
	message string
}
func (err FitnessError) Error() string {
	return err.message
}

type Fitness struct {
}

func (f *Fitness) Execute(individuals shared.Individuals) error {
	for _, v := range individuals {
		word, ok := (*v).(*Word)
		if !ok {
			return FitnessError{"Invalid individual type!"}
		}
		genotype := word.GetGenotype()
		sum := 0
		for i := 0; i < len(genotype); i++ {
			sum += int(math.Abs(float64(int(TargetWord[i]) - int(genotype[i]))))
		}
		word.SetFitness(float64(sum))
	}
	return nil
}
