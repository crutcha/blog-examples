package main

import (
	"container/heap"
	"fmt"
)

type PathVector struct {
	prev string
	cost int
}

type PathTable struct {
	paths map[string]PathVector
}

func (pt *PathTable) String() {
	for key, value := range pt.paths {
		fmt.Printf("Destination: %s: ", key)
		fmt.Printf("Cost: %d ", value.cost)
		fmt.Printf("Previous: %s\n", value.prev)
	}
}

func (g *Graph) DijkstraPaths(src string) PathTable {
	visitMap := make(map[string]bool)
	pathMap := make(map[string]PathVector)

	pq := make(EdgePQ, 0)
	heap.Push(&pq, &Edge{dest: src, weight: 0})

	for node, _ := range g.nodes {
		// Used to track whether or not a specific node has been
		// visited yet or not.
		visitMap[node] = false

		// Set initial cost to infinity unless the current
		// node is our own
		if node == src {
			pathMap[node] = PathVector{
				cost: 0,
			}
		} else {
			pathMap[node] = PathVector{
				cost: -1,
			}
		}
	}

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*Edge)
		currentEdges := g.nodes[current.dest]

		if visitMap[current.dest] {
			continue
		} else {
			visitMap[current.dest] = true
		}

		for _, neighbor := range currentEdges {
			currentShortest := pathMap[current.dest].cost + neighbor.weight
			previousShortest := pathMap[neighbor.dest].cost
			if currentShortest < previousShortest || previousShortest == -1 {
				update := pathMap[neighbor.dest]
				update.cost = currentShortest
				update.prev = current.dest
				pathMap[neighbor.dest] = update
			}

			// place into PQ
			heap.Push(&pq, neighbor)
		}

	}

	return PathTable{paths: pathMap}
}
