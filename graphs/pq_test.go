package main

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPriorityQueue(t *testing.T) {
	fmt.Println("testing")
	assert := assert.New(t)
	pq := make(EdgePQ, 0)

	heap.Push(&pq, Edge{dest: "5.5.5.5", weight: 10})
	heap.Push(&pq, Edge{dest: "1.1.1.1", weight: 0})
	heap.Push(&pq, Edge{dest: "4.4.4.4", weight: 4})
	heap.Push(&pq, Edge{dest: "3.3.3.3", weight: 3})

	popped := heap.Pop(&pq).(Edge)
	assert.Equal(popped.weight, 0)
	assert.Equal(popped.dest, "1.1.1.1")

	secondPop := heap.Pop(&pq).(Edge)
	assert.Equal(secondPop.weight, 3)
	assert.Equal(secondPop.dest, "3.3.3.3")

	heap.Push(&pq, Edge{dest: "2.2.2.2", weight: 2})
	newPriority := heap.Pop(&pq).(Edge)
	assert.Equal(newPriority.weight, 2)
	assert.Equal(newPriority.dest, "2.2.2.2")
}
