package main

import (
	"fmt"
	"math"
)

type PathVector struct {
	dest string
	prev string
	cost float64
}

type PathTable struct {
	paths map[string]PathVector
}

func (pt *PathTable) String() {
	fmt.Println("not quite yet")
}

func (g *Graph) DijkstraPaths(src string) PathTable {
	visitMap := make(map[string]bool)
	pathMap := make(map[string]PathVector)

	// Used to track whether or not a specific node has been
	// visited yet or not.
	for node, _ := range g.nodes {
		visitMap[node] = false

		// Set initial cost to infinity unless the current
		// node is our own
		if node == src {
			pathMap[node] = PathVector{
				dest: node,
				cost: 0,
			}
		} else {
			pathMap[node] = PathVector{
				dest: node,
				cost: math.Inf(1),
			}
		}
	}

	return PathTable{paths: pathMap}
}
