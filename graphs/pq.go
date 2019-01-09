// priority queue interface to be used by graph structs
package main

type EdgePQ []Edge

func (pq EdgePQ) Len() int {
	return len(pq)
}

func (pq EdgePQ) Less(i, j int) bool {
	// min heap, return lowest
	return pq[i].weight < pq[j].weight
}

func (pq EdgePQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *EdgePQ) Push(x interface{}) {
	item := x.(Edge)
	*pq = append(*pq, item)
}

func (pq *EdgePQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
