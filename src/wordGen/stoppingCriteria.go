package main

import (
	"alGoGen/shared"
)

type WordCriteria struct {
}

func (s *WordCriteria) Execute(epoch int, individuals shared.Individuals) bool {
	for _, v := range individuals {
		fitness := (*v).Fitness()
		if fitness == 0 {
			return true
		}
	}
	return false
}
