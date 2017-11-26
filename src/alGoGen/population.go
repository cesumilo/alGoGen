package alGoGen

import (
	"math"
	"math/rand"
	"time"
	"alGoGen/operators"
	"alGoGen/shared"
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
	p.config = settings
	p.individuals = make([]*shared.Individual, settings.settings.populationSize, settings.settings.populationSize)
	p.randomGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < settings.settings.populationSize; i++ {
		idv, ok := (*settings.createIndividualOperator).Execute(i)
		if (!ok) {
			// TODO: error handling
		}

		p.individuals[i] = idv
	}
}

func (p *Population) Run() {
	i := 0

	for (*p.config.stoppingCriteria).Execute(i, p.fitnessValues) {

		fitnessValues, fitnessOk := (*p.config.fitnessOperator).Execute(p.individuals)
		if (!fitnessOk) {
			// TODO: error handling
		}

		totalNumOfOffspring := int(float32(p.config.settings.populationSize) * p.config.settings.offspringProportion)
		totalNumOfOffspring = int(math.Min(float64(totalNumOfOffspring), float64(p.config.settings.populationSize)))
		totalNumOfSelectedMod := totalNumOfOffspring % p.config.settings.generatedOffspringNumber
		totalNumOfSelected := totalNumOfOffspring / p.config.settings.generatedOffspringNumber
		if totalNumOfSelectedMod != 0 {
			totalNumOfSelected = totalNumOfSelected + totalNumOfSelectedMod * p.config.settings.generatedOffspringNumber
		}

		selectedIndividuals, selectOk := (*p.config.parentSelectionOperator).Execute(p.individuals, fitnessValues, totalNumOfSelected)
		if (!selectOk) {
			// TODO: error handling
		}

		selectedIdx := 0
		offsprings := make([]*shared.Individual, totalNumOfOffspring, totalNumOfOffspring)
		for totalNumOfOffspring > 0 {
			idv1, idv2 := selectedIndividuals[selectedIdx], selectedIndividuals[selectedIdx + 1]

			createdOffsprings, offspringOk := (*p.config.crossoverOperator).Execute(idv1, idv2, p.config.settings.generatedOffspringNumber)
			if (!offspringOk) {
				// TODO: error handling
			}
			offsprings = append(offsprings, createdOffsprings...)
			totalNumOfOffspring -= p.config.settings.generatedOffspringNumber
		}

		selectedNextIndividuals, selectNOk := (*p.config.populationSelectionOperator).Execute(p.individuals, fitnessValues, len(p.individuals) - len(offsprings))
		if (!selectNOk) {
			// TODO: error handling
		}

		newPopulation := append(offsprings, selectedNextIndividuals...)
		mutatedPopulation := make([]*shared.Individual, len(newPopulation), len(newPopulation))
		for i, v := range newPopulation {
			var newIndividual *shared.Individual

			if p.randomGenerator.Float32() <= p.config.settings.mutationProbability {
				newIdv, mutationOk := (*p.config.mutationOperator).Execute(v, i)
				if (!mutationOk) {
					// TODO: error handling
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
