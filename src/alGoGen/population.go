package alGoGen

import (
	"math"
	"math/rand"
	"time"
	"alGoGen/operators"
	"alGoGen/shared"
	"sort"
)

type PopulationError struct {
	message string
}
func (err PopulationError) Error() string {
	return err.message
}

type Population struct {
	individuals shared.Individuals
	config PopulationSettings
	randomGenerator *rand.Rand
}

type PopulationSettings struct {
	Settings *Settings

	StoppingCriteria operators.StoppingCriterion
	CreateIndividualOperator operators.CreateIndividual
	FitnessOperator operators.Fitness
	CrossoverOperator operators.Crossover
	ParentSelectionOperator operators.ParentSelection
	PopulationSelectionOperator operators.PopulationSelection
	MutationOperator operators.Mutation

	FitnessStatisticsOperator operators.Statistics
	ParentStatisticsOperator operators.Statistics
	OffspringStatisticsOperator operators.Statistics
	SelectionStatisticsOperator operators.Statistics
}

func (p *Population) Init(settings PopulationSettings) {

	switch {
	case settings.StoppingCriteria == nil:
		panic(PopulationError{"Invalid operator: stoppingCriteria"})
	case settings.CreateIndividualOperator == nil:
		panic(PopulationError{"Invalid operator: createIndividual"})
	case settings.FitnessOperator == nil:
		panic(PopulationError{"Invalid operator: fitness"})
	case settings.CrossoverOperator == nil:
		panic(PopulationError{"Invalid operator: crossover"})
	case settings.ParentSelectionOperator == nil:
		panic(PopulationError{"Invalid operator: parentSelection"})
	case settings.PopulationSelectionOperator == nil:
		panic(PopulationError{"Invalid operator: populationSelection"})
	case settings.MutationOperator == nil:
		panic(PopulationError{"Invalid operator: mutation"})
	}

	p.config = settings
	p.randomGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < settings.Settings.PopulationSize; i++ {
		idv, err := settings.CreateIndividualOperator.Execute(i)
		if err != nil {
			panic(err)
		}
		p.individuals = append(p.individuals, idv)
	}
}

func (p *Population) Run() {

	i := 0

	fitnessOk := p.config.FitnessOperator.Execute(p.individuals)
	if fitnessOk != nil {
		panic(fitnessOk)
	}

	if p.config.FitnessStatisticsOperator != nil {
		p.config.FitnessStatisticsOperator.Compute(p.individuals)
	}

	for !p.config.StoppingCriteria.Execute(i, p.individuals) {

		totalNumOfOffspring := int(float32(p.config.Settings.PopulationSize) * p.config.Settings.OffspringProportion)
		totalNumOfOffspring = int(math.Min(float64(totalNumOfOffspring), float64(p.config.Settings.PopulationSize)))
		totalNumOfSelectedMod := totalNumOfOffspring % p.config.CrossoverOperator.NumberOfOffspring()
		totalNumOfSelected := totalNumOfOffspring / p.config.CrossoverOperator.NumberOfOffspring()
		if totalNumOfSelectedMod != 0 {
			totalNumOfSelected = totalNumOfSelected + totalNumOfSelectedMod
		}

		totalNumOfSelected *= 2

		selectedIndividuals, selectOk := p.config.ParentSelectionOperator.Execute(p.individuals, totalNumOfSelected)
		if selectOk != nil {
			panic(selectOk)
		}

		if p.config.ParentStatisticsOperator != nil {
			p.config.ParentStatisticsOperator.Compute(selectedIndividuals)
		}

		selectedIdx := 0
		var offspring shared.Individuals
		for totalNumOfOffspring > 0 {
			idv1, idv2 := selectedIndividuals[selectedIdx], selectedIndividuals[selectedIdx + 1]

			createdOffspring, offspringOk := p.config.CrossoverOperator.Execute(idv1, idv2)
			if offspringOk != nil {
				panic(offspringOk)
			}
			offspring = append(offspring, createdOffspring...)
			totalNumOfOffspring -= p.config.CrossoverOperator.NumberOfOffspring()
			selectedIdx += 2
		}

		if totalNumOfOffspring < 0 {
			offspring = offspring[:len(offspring) + totalNumOfOffspring]
		}

		if p.config.OffspringStatisticsOperator != nil {
			p.config.OffspringStatisticsOperator.Compute(offspring)
		}

		selectedNextIndividuals, selectNOk := p.config.PopulationSelectionOperator.Execute(p.individuals, len(p.individuals) - len(offspring))
		if selectNOk != nil {
			panic(selectNOk)
		}

		if p.config.SelectionStatisticsOperator != nil {
			p.config.SelectionStatisticsOperator.Compute(selectedNextIndividuals)
		}

		newPopulation := append(offspring, selectedNextIndividuals...)
		var mutatedPopulation []*shared.Individual
		for _, v := range newPopulation {
			var newIndividual *shared.Individual

			if p.randomGenerator.Float32() <= p.config.Settings.MutationProbability {
				newIdv, mutationOk := p.config.MutationOperator.Execute(v)
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

		fitnessOk := p.config.FitnessOperator.Execute(p.individuals)
		if fitnessOk != nil {
			panic(fitnessOk)
		}

		if p.config.FitnessStatisticsOperator != nil {
			p.config.FitnessStatisticsOperator.Compute(p.individuals)
		}

		i++
	}
}

func (p *Population) GetIndividuals() shared.Individuals {
	return p.individuals
}

func (p *Population) GetBestIndividual() *shared.Individual {
	sort.Sort(p.individuals)
	return p.individuals[0]
}
