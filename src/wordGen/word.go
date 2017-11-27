package main

import (
	"math/rand"
)

type Word struct {
	genotype string
	fitness float64
	id int
}

const LetterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = LetterBytes[rand.Intn(len(LetterBytes))]
	}
	return string(b)
}

func (w *Word) Init(id int, length int) {
	w.id = id
	w.genotype = randStringBytes(length)
}

func (w *Word) Id() int {
	return w.id
}

func (w *Word) Fitness() float64 {
	return w.fitness
}

func (w *Word) SetFitness(fitness float64) {
	w.fitness = fitness
}

func (w *Word) GetGenotype() string {
	return w.genotype
}
