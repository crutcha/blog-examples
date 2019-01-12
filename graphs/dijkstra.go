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
	fmt.Println("not quite yet")
}

func (g *Graph) DijkstraPaths(src string) PathTable {
	fmt.Println("DIJKSTRA TEST")
	visitMap := make(map[string]bool)
	pathMap := make(map[string]PathVector)

	pq := make(EdgePQ, 0)
	heap.Push(&pq, &Edge{dest: src, weight: 0})
	fmt.Println("PQ: ", pq)

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
			fmt.Println("Current shortest: ", currentShortest)
			fmt.Println("Old Shortest: ", pathMap[neighbor.dest].cost)
			if currentShortest < previousShortest || previousShortest == -1 {
				fmt.Println("processing update")
				update := pathMap[neighbor.dest]
				fmt.Println(update)
				update.cost = currentShortest
				update.prev = current.dest
				pathMap[neighbor.dest] = update
				fmt.Println(pathMap[neighbor.dest])
			}

			// place into PQ
			heap.Push(&pq, neighbor)
		}

		fmt.Println("Updated PathTable")
		fmt.Println(pathMap)
		fmt.Println("------")
	}

	return PathTable{paths: pathMap}
}
