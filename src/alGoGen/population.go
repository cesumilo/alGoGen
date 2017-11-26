package alGoGen

import (
	"math"
	"math/rand"
	"time"
	"alGoGen/operators"
	"alGoGen/shared"
	"log"
)

type Population struct {
	individuals []*shared.Individual
	fitnessValues []float64
	config PopulationSettings
	randomGenerator *rand.Rand
}

type PopulationSettings struct {
	settings *Settings

	stoppingCriteria *operators.StoppingCriterion
	createIndividualOperator *operators.CreateIndividual
	fitnessOperator *operators.Fitness
	crossoverOperator *operators.Crossover
	parentSelectionOperator *operators.ParentSelection
	populationSelectionOperator *operators.PopulationSelection
	mutationOperator *operators.Mutation
}

func (p *Population) Init(settings PopulationSettings) {
	defer p.errorHandler()

	p.config = settings
	p.individuals = make([]*shared.Individual, settings.settings.populationSize, settings.settings.populationSize)
	p.randomGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < settings.settings.populationSize; i++ {
		idv, err := (*settings.createIndividualOperator).Execute(i)
		if err != nil {
			panic(err)
		}
		p.individuals[i] = idv
	}
}

func (p *Population) Run() {
	defer p.errorHandler()

	i := 0
	for (*p.config.stoppingCriteria).Execute(i, p.fitnessValues) {

		fitnessValues, fitnessOk := (*p.config.fitnessOperator).Execute(p.individuals)
		if fitnessOk != nil {
			panic(fitnessOk)
		}

		totalNumOfOffspring := int(float32(p.config.settings.populationSize) * p.config.settings.offspringProportion)
		totalNumOfOffspring = int(math.Min(float64(totalNumOfOffspring), float64(p.config.settings.populationSize)))
		totalNumOfSelectedMod := totalNumOfOffspring % p.config.settings.generatedOffspringNumber
		totalNumOfSelected := totalNumOfOffspring / p.config.settings.generatedOffspringNumber
		if totalNumOfSelectedMod != 0 {
			totalNumOfSelected = totalNumOfSelected + totalNumOfSelectedMod * p.config.settings.generatedOffspringNumber
		}

		selectedIndividuals, selectOk := (*p.config.parentSelectionOperator).Execute(p.individuals, fitnessValues, totalNumOfSelected)
		if selectOk != nil {
			panic(selectOk)
		}

		selectedIdx := 0
		offspring := make([]*shared.Individual, totalNumOfOffspring, totalNumOfOffspring)
		for totalNumOfOffspring > 0 {
			idv1, idv2 := selectedIndividuals[selectedIdx], selectedIndividuals[selectedIdx + 1]

			createdOffspring, offspringOk := (*p.config.crossoverOperator).Execute(idv1, idv2, p.config.settings.generatedOffspringNumber)
			if offspringOk != nil {
				panic(offspringOk)
			}
			offspring = append(offspring, createdOffspring...)
			totalNumOfOffspring -= p.config.settings.generatedOffspringNumber
		}

		selectedNextIndividuals, selectNOk := (*p.config.populationSelectionOperator).Execute(p.individuals, fitnessValues, len(p.individuals) - len(offspring))
		if selectNOk != nil {
			panic(selectNOk)
		}

		newPopulation := append(offspring, selectedNextIndividuals...)
		mutatedPopulation := make([]*shared.Individual, len(newPopulation), len(newPopulation))
		for i, v := range newPopulation {
			var newIndividual *shared.Individual

			if p.randomGenerator.Float32() <= p.config.settings.mutationProbability {
				newIdv, mutationOk := (*p.config.mutationOperator).Execute(v, i)
				if mutationOk != nil {
					panic(mutationOk)
				}
				newIndividual = newIdv
			} else {
				newIndividual = v
			}

			mutatedPopulation = append(mutatedPopulation, newIndividual)
		}

		p.individuals = mutatedPopulation
		i++
	}
}

func (p *Population) errorHandler() {
	if r := recover(); r != nil {
		log.Print(r)
	}
}
