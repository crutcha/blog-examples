package main

import (
	"fmt"
	"testing"
)

func TestGraph(t *testing.T) {
	graph := NewGraph()
	firstNode := NewNode("testnode")
	secondNode := NewNode("secondnode")

	graph.AddNode(firstNode)
	graph.AddNode(secondNode)

	graph.AddEdge(firstNode, secondNode, 10)

	//fmt.Println(graph.nodes)
	//fmt.Println(firstNode)

	for k, v := range graph.nodes {
		fmt.Println("Key: ", k)
		for _, e := range v {
			fmt.Println("Edge: ", e.dest.name)
			fmt.Println("Edge Weight: ", e.weight)
		}
		fmt.Println(v)
	}
}
