package main

import (
	"errors"
	"fmt"
)

type Graph struct {
	nodes map[string][]*Edge
}

/*
Since this will be an edge list implementation, the source will
already be known since it will be the key in graph.nodes, we only
need the destination.
*/
type Edge struct {
	dest   string
	weight int
}

func NewGraph() Graph {
	nodes := make(map[string][]*Edge)
	return Graph{
		nodes: nodes,
	}
}

func (g *Graph) AddNode(name string) {
	// TODO: do we care to check if node exists already and return
	// an error instead?
	var edges []*Edge
	g.nodes[name] = edges
}

func (g *Graph) AddEdge(src string, dst string, weight int) error {
	var error error

	// TODO: loop through edges and check if it's a duplicate and ignore?

	// double check we actually have that node for both source and
	// dest before processing
	if _, ok := g.nodes[src]; ok {
		newEdge := &Edge{
			dest:   dst,
			weight: weight,
		}
		g.nodes[src] = append(g.nodes[src], newEdge)
	} else {
		error = errors.New("Source node does not exist in graph.")
		return error
	}

	if _, ok := g.nodes[dst]; ok {
		newEdge := &Edge{
			dest:   src,
			weight: weight,
		}
		g.nodes[dst] = append(g.nodes[dst], newEdge)
	} else {
		error = errors.New("Destination node does not exists in graph.")
		return error
	}

	// this seems redundant and stupid?
	return error
}

func (g *Graph) String() {
	for node, edges := range g.nodes {
		fmt.Printf("%-15s", node)

		fmt.Printf("[ ")
		for _, edge := range edges {
			fmt.Printf("(%s, ", edge.dest)
			fmt.Printf("%d) ", edge.weight)
		}
		fmt.Printf("]\n")
	}
}
