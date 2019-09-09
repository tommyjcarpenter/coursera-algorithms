package heap

import (
	"container/heap"
)

// used in course_3_p3q1

// An Vertex is something we manage in a Priority queue.
type TreeVertex struct {
	Index    int // The Index of the item in the heap THIS IS USED INTERNALLY IN HEAP OPERATIONS
	Priority int // The Priority of the item in the queue.
	Left     *TreeVertex
	Right    *TreeVertex
	// The Index is needed by update and is maintained by the heap.Interface methods.
}

// A PQMID implements heap.Interface and holds TreeVertexss.
type PQMID []*TreeVertex

func (pq PQMID) Len() int { return len(pq) }

func (pq PQMID) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PQMID) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j

}

// never call this directly!
func (pq *PQMID) Push(x interface{}) {
	n := len(*pq)
	item := x.(*TreeVertex)
	item.Index = n
	*pq = append(*pq, item)
}

// never call this directly!
func (pq *PQMID) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// heap.Pop calls this. Don't ave clients use Pop directly!
// The confusion here is that the interface requires Pop to be implemented, but you DONT want clients calling the above Pop directly because it doesnt fix the heap.
// Rather you want clients calling the heap method, heap.Pop
func (pq *PQMID) ExtractMin() interface{} {
	return heap.Pop(pq)
}

// heap.Push calls this.
func (pq *PQMID) AddElement(x interface{}) {
	heap.Push(pq, x)
}

func (pq *PQMID) HInit() {
	heap.Init(pq)
}

// update modifies the Priority and value of an TreeVertexs in the queue.
func (pq *PQMID) Update(item *TreeVertex, Priority int) {
	item.Priority = Priority
	heap.Fix(pq, item.Index)
}
