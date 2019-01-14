package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDijkstraPathTable(t *testing.T) {
	assert := assert.New(t)
	graph := NewGraph()

	graph.AddNode("A")
	graph.AddNode("B")
	graph.AddNode("C")
	graph.AddNode("D")
	graph.AddNode("E")

	graph.AddEdge("A", "B", 6)
	graph.AddEdge("A", "D", 1)
	graph.AddEdge("B", "D", 2)
	graph.AddEdge("B", "E", 2)
	graph.AddEdge("C", "B", 5)
	graph.AddEdge("C", "E", 5)
	graph.AddEdge("E", "D", 1)

	graph.String()
	result := graph.DijkstraPaths("A")
	result.String()

	assert.Equal(result.paths["B"].cost, 3)
	assert.Equal(result.paths["B"].prev, "D")
	assert.Equal(result.paths["C"].cost, 7)
	assert.Equal(result.paths["C"].prev, "E")
	assert.Equal(result.paths["E"].cost, 2)
	assert.Equal(result.paths["E"].prev, "D")
}

func TestDijkstraShortestPath(t *testing.T) {
	assert := assert.New(t)
	graph := NewGraph()

	graph.AddNode("A")
	graph.AddNode("B")
	graph.AddNode("C")
	graph.AddNode("D")
	graph.AddNode("E")

	graph.AddEdge("A", "B", 6)
	graph.AddEdge("A", "D", 1)
	graph.AddEdge("B", "D", 2)
	graph.AddEdge("B", "E", 2)
	graph.AddEdge("C", "B", 5)
	graph.AddEdge("C", "E", 5)
	graph.AddEdge("E", "D", 1)

	sp := graph.ShortestPath("A", "B")
	sp2 := graph.ShortestPath("A", "C")
	assert.Equal(sp, []string{"A", "D", "B"})
	assert.Equal(sp2, []string{"A", "D", "E", "C"})
}
