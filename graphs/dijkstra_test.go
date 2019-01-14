package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDijkstraPathTable(t *testing.T) {
	assert := assert.New(t)
	graph := NewGraph()

	/*
		graph.AddNode("1.1.1.1")
		graph.AddNode("2.2.2.2")
		graph.AddNode("3.3.3.3")
		graph.AddNode("4.4.4.4")
		graph.AddNode("5.5.5.5")
		graph.AddNode("6.6.6.6")

		graph.AddEdge("1.1.1.1", "2.2.2.2", 10)
		graph.AddEdge("1.1.1.1", "3.3.3.3", 10)
		graph.AddEdge("2.2.2.2", "4.4.4.4", 100)
		graph.AddEdge("2.2.2.2", "5.5.5.5", 1)
		graph.AddEdge("3.3.3.3", "5.5.5.5", 10)
		graph.AddEdge("4.4.4.4", "6.6.6.6", 10)
		graph.AddEdge("5.5.5.5", "6.6.6.6", 10)
	*/

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
