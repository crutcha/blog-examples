package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraph(t *testing.T) {
	assert := assert.New(t)

	graph := NewGraph()
	graph.AddNode("testnode")
	graph.AddNode("secondnode")

	graph.AddEdge("testnode", "secondnode", 10)

	assert.Equal(1, len(graph.nodes["testnode"]))
	graph.String()
}
