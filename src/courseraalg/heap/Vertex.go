package heap

// An Vertex is something we manage in a Priority queue.
type Vertex struct {
	Index    int // The Index of the item in the heap THIS IS USED INTERNALLY IN HEAP OPERATIONS
	Nodeid   int // The satelite data of the original node id, which we care about
	Priority int // The Priority of the item in the queue.
	// The Index is needed by update and is maintained by the heap.Interface methods.
}

type Vertexf struct {
	Index    int // The Index of the item in the heap THIS IS USED INTERNALLY IN HEAP OPERATIONS
	Nodeid   int // The satelite data of the original node id, which we care about
	Weight   int
	Length   int
	Priority float64 // The Priority of the item in the queue.
	// The Index is needed by update and is maintained by the heap.Interface methods.
}
