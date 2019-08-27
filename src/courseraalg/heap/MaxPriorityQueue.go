package heap

import (
	"container/heap"
)

// A MaxPQ implements heap.Interface and holds Vertexfs.
type MaxPQ []*Vertexf

/*
heap requires
len
less
swap
push
pop

they writ this as
type Interface interface {
        sort.Interface
        Push(x interface{}) // add x as element Len()
        Pop() interface{}   // remove and return element Len() - 1.
}

but sort.Interface is

type Interface interface {
        // Len is the number of elements in the collection.
        Len() int
        // Less reports whether the element with
        // index i should sort before the element with index j.
        Less(i, j int) bool
        // Swap swaps the elements with indexes i and j.
        Swap(i, j int)
}
*/
func (pq MaxPQ) Len() int { return len(pq) }

func (pq MaxPQ) Less(i, j int) bool {
	if pq[i].Priority > pq[j].Priority {
		return true
	} else if pq[i].Priority == pq[j].Priority {
		return pq[i].Weight > pq[j].Weight
	} else {
		return false
	}
}

func (pq MaxPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j

}

func (pq *MaxPQ) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Vertexf)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *MaxPQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// this is so that you can call pq.Pop, rather than the client code doing heap.Pop and having to import heap
func (pq *MaxPQ) ExtractMax() interface{} {
	// heap.Pop eventually calls the pop above.
	// do NOT want end clients calling the above Pop directly
	return heap.Pop(pq)
}

func (pq *MaxPQ) HInit() {
	heap.Init(pq)
}

// update modifies the Priority and value of an Vertexf in the queue.
func (pq *MaxPQ) Update(item *Vertexf, Priority float64) {
	item.Priority = Priority
	heap.Fix(pq, item.Index)
}
