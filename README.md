# alGoGen
A framework coded in Go for Genetic Algorithm.

## Motivation
// TODO

## Pseudo-code
1. Generate population of individuals
2. While a stopping criteria is not satisfied:
    1. Compute fitness of each individual
    2. Apply selection operator on population
    3. Select parent individuals
    4. Apply crossover operator on parents
    5. Apply mutation operator on offspring

## Generics
### Genetic Algorithm
- Population
    - Generates individuals
    - Evolves population using operators
- Individual
    - Fitness value
    - Identifier
    - Genotype
- Operators:
    - Fitness
        - Returns list of fitness value of individuals
    - Selection
        - Selects enough individual to create next population 
    - Crossover
        - Generates offspring using selected individuals
    - Mutation
        - Mutates individuals
        - Mutation probability

### Input/Output
- Logging
- Graphs
- Save/Load

## Configuration
// TODO
### config.json
```json
{
    "Users": ["UserA","UserB"],
    "Groups": ["GroupA"]
}
```
### Example: Go loading program
```go
package main

import (
    "encoding/json"
    "os"
    "fmt"
)

type Configuration struct {
    Users    []string
    Groups   []string
}

func main() {
    file, _ := os.Open("conf.json")
    decoder := json.NewDecoder(file)
    configuration := Configuration{}
    err := decoder.Decode(&configuration)
    if err != nil {
        fmt.Println("error:", err)
    }
    fmt.Println(configuration.Users) // output: [UserA, UserB]
}
```

## Set development environment up
```
$ . ./dev/setup.sh
```
