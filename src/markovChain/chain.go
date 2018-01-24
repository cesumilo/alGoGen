package markovChain

import (
	"gonum.org/v1/gonum/mat"
)

type node struct {
	index int
	links []string
}

type Chain struct {
	states map[string]*node
	trans mat.Matrix
}

func (c *Chain) Init(states []string) {
	for i, s := range states {
		c.states[s].index = i
	}
	c.trans = mat.NewDense(len(states), len(states), nil)
}

func (c *Chain) AddLink(s1 string, s2 string) bool {
	_, ok1 := c.states[s1]
	_, ok2 := c.states[s2]

	if !ok1 || !ok2 {
		return false
	}

	c.states[s1].links = append(c.states[s1].links, s2)
	return true
}

func (c *Chain) RemoveLink(s1 string, s2 string) bool {
	_, ok1 := c.states[s1]
	_, ok2 := c.states[s2]

	if !ok1 || !ok2 {
		return false
	}

	for i, s := range c.states[s1].links {
		if s == s2 {
			begin := c.states[s1].links[0:i]

			if i+1 < len(c.states[s1].links) {
				end := c.states[s1].links[i+1:len(c.states[s1].links)]
				begin = append(begin, end...)
			}

			c.states[s1].links = begin
			return true
		}
	}

	return false
}