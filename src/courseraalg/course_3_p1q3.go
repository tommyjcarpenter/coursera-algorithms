package main

import (
	"courseraalg/heap"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const NODES = 500

func getE(fname string) ([NODES][][2]int, heap.PriorityQueue) {

	b, _ := ioutil.ReadFile(fname)

	lines := strings.Split(string(b), "\n")

	var adj_to [NODES][][2]int
	for i := 0; i < NODES; i++ {
		adj_to[i] = make([][2]int, 0)
	}

	var smallest_adj [NODES]int         // used to create the intial heap
	smallest_adj[0] = -9999999999999999 // make sure node 0 gets popped
	for i := 1; i < NODES; i++ {
		smallest_adj[i] = 100000000
	}

	for _, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 {
			continue
		}
		s := strings.Fields(l)

		// this graph is undirected
		endpoint1, _ := strconv.Atoi(s[0])
		endpoint1-- // 0 Indexing
		endpoint2, _ := strconv.Atoi(s[1])
		endpoint2--
		length, _ := strconv.Atoi(s[2])

		//undirected
		adj_to[endpoint1] = append(adj_to[endpoint1], [2]int{endpoint2, length})
		adj_to[endpoint2] = append(adj_to[endpoint2], [2]int{endpoint1, length})
		if endpoint1 == 0 { // we will start from node 0; all other paths are infinity
			if length < smallest_adj[endpoint2] {
				smallest_adj[endpoint2] = length
			}
		}
	}

	// we don't need to add node 0 in
	pq := make(heap.PriorityQueue, NODES)
	for i := 0; i < NODES; i++ {
		pq[i] = &heap.Vertex{i, i, smallest_adj[i]}
	}
	pq.HInit()

	return adj_to, pq

}

func main() {
	//adj_to, edges_incoming, pq := getE("course_2_p_2.txt")
	adj_to, pq := getE("course_3_p1q3.txt")

	// boolean array of whether we checked each node, we start from 0
	var already_checked [NODES]bool
	already_checked[0] = true
	for i := 1; i < NODES; i++ {
		already_checked[i] = false
	}

	// array of final edges
	final_mst_edge_lengths := make([]int, 0)

	//pop node 0
	v := pq.ExtractMin().(*heap.Vertex)
	fmt.Println(v)

	for pq.Len() > 0 {

		v := pq.ExtractMin().(*heap.Vertex)                                 // pop the node
		already_checked[v.Nodeid] = true                                    // mark as checked
		final_mst_edge_lengths = append(final_mst_edge_lengths, v.Priority) // append edge

		// check whether v -> w is a smaller edge for all w not yet checked (that are connected to v)
		for _, o := range adj_to[v.Nodeid] {
			end := o[0]
			length := o[1]

			for i := range pq { // find the item; TODO, IS THERE A BETTER WAY TO DO THIS???
				if pq[i].Nodeid == end { // we found the node in the heap, now check it
					if length < pq[i].Priority { // node i in the heap can now be reached by length
						pq.Update(pq[i], length)
					}
				}
			}
		}
	}
	s := 0
	for _, i := range final_mst_edge_lengths {
		s += i
	}
	fmt.Println(s)
}
