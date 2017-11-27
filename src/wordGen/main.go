package main

import (
	"alGoGen"
	"alGoGen/operators/selection"
	"log"
	"fmt"
)

const WordLength = 10
const PopulationSize = 200
const OffspringProportion = 0.8
const MutationProbability = 0.01
const K = 10
const TargetWord = "HelloWorld"

func main() {
	var pop alGoGen.Population
	var parentSelector selection.RouletteWheel
	var popSelector selection.Tournament

	popSelector.Init(K)

	settings := alGoGen.Settings{PopulationSize, OffspringProportion, MutationProbability}
	config := alGoGen.PopulationSettings{
		&settings,
		&WordCriteria{},
		&CreateWord{WordLength},
		&Fitness{},
		&Crossover{},
		&parentSelector,
		&popSelector,
		&Mutation{}}

	pop.Init(config)
	pop.Run()
	bestIndividual := pop.GetBestIndividual()
	word, ok := (*bestIndividual).(*Word)
	if !ok {
		log.Fatal("Invalid best individual type!")
	}
	fmt.Printf("Best Genotype (%f): ", word.fitness)
	fmt.Println(word.GetGenotype())
}
