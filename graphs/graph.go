package main

import "errors"

type Graph struct {
	nodes map[*Node][]*Edge
}

type Node struct {
	visited bool
	name    string
}

/*
Since this will be an adjacency list implementation, the source will
already be known since it will be the key in graph.nodes, we only need
the destination
*/
type Edge struct {
	dest   *Node
	weight int
}

func NewGraph() Graph {
	nodes := make(map[*Node][]*Edge)
	return Graph{
		nodes: nodes,
	}
}

func NewNode(name string) *Node {
	node := Node{
		name:    name,
		visited: false,
	}
	return &node
}

func (g *Graph) AddNode(node *Node) {
	var edges []*Edge
	g.nodes[node] = edges
}

func (g *Graph) AddEdge(src *Node, dst *Node, weight int) error {
	var error error

	// double check we actually have that node
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

	return error
}
