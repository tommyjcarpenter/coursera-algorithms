package heap

import (
	"container/heap"
)

// A PriorityQueue implements heap.Interface and holds Vertexs.
type PriorityQueue []*Vertex

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j

}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Vertex)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) ExtractMin() interface{} {
	return heap.Pop(pq)
}

func (pq *PriorityQueue) HInit() {
	heap.Init(pq)
}

// update modifies the Priority and value of an Vertex in the queue.
func (pq *PriorityQueue) Update(item *Vertex, Priority int) {
	item.Priority = Priority
	heap.Fix(pq, item.Index)
}
