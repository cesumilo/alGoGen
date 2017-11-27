package alGoGen

import (
	"encoding/json"
	"os"
)

type Settings struct {
	PopulationSize int
	OffspringProportion float32
	MutationProbability float32
}

func LoadSettings(filename string) *Settings {
	file, openErr := os.Open(filename)

	if openErr != nil {
		panic(openErr)
	}

	decoder := json.NewDecoder(file)
	settings := Settings{}
	decodeErr := decoder.Decode(&settings)

	if decodeErr != nil {
		panic(decodeErr)
	}

	return &settings
}
