package main

// https://pkg.go.dev/container/heap#example-package-PriorityQueue

import "container/heap"

// Priority computation: neighborNode.rank = cost + neighbor.PathEstimatedCost(to)
// See https://github.com/beefsack/go-astar/blob/master/astar.go
// Seems to be literally the distance

// An Item is something we manage in a priority queue.
type Node struct {
	distance int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index      int // The index of the item in the heap.
	coordinate Coordinate
	blocker    bool
	visited    bool
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest distance, so we use lower than here.
	return pq[i].distance < pq[j].distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Node)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Node, distance int, coordinate Coordinate, blocker bool) {
	item.distance = distance
	item.coordinate = coordinate
	item.blocker = blocker
	heap.Fix(pq, item.index)
}
