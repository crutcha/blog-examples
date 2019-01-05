package main

import "testing"

func TestDijkstraPathTable(t *testing.T) {
	graph := NewGraph()

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

	graph.String()
}
