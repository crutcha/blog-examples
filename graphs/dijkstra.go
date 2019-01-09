package main

import (
	"container/heap"
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

	pq := make(EdgePQ, 0)

	for node, _ := range g.nodes {
		// Used to track whether or not a specific node has been
		// visited yet or not.
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

	fmt.Println("wut")
	// start PQ with source node weight of 0
	heap.Push(&pq, Edge{dest: "5.5.5.5", weight: 10})
	heap.Push(&pq, Edge{dest: "1.1.1.1", weight: 0})
	fmt.Println("NOW: ", pq)
	heap.Push(&pq, Edge{dest: "2.2.2.2", weight: 1})
	fmt.Println("NOW: ", pq)
	popped := heap.Pop(&pq)
	fmt.Println("POPPPED: ", popped)
	fmt.Println("NOW: ", pq)

	// Start traversing breadth-first through all nodes in map
	// starting with our own

	return PathTable{paths: pathMap}
}
