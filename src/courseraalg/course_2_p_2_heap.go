package main

import (
	"courseraalg/heap"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getE(fname string) ([200][][2]int, heap.PriorityQueue) {

	b, _ := ioutil.ReadFile(fname)

	lines := strings.Split(string(b), "\n")

	var edges_outgoing [200][][2]int // edges from is only for when we select a node to move into the set x
	for i := 0; i < 200; i++ {
		edges_outgoing[i] = make([][2]int, 0)
	}

	var best_from_s [200]int // used to create the intial heap
	best_from_s[0] = 0
	for i := 1; i < 200; i++ {
		best_from_s[i] = 1000000
	}

	for lindex, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 {
			continue
		}
		line := strings.Fields(l)
		for jindex, j := range line {

			if jindex > 0 { // first column is a repeat of the row index
				s := strings.Split(j, ",")
				end, _ := strconv.Atoi(s[0])
				length, _ := strconv.Atoi(s[1])
				end--

				a := [2]int{end, length}
				edges_outgoing[lindex] = append(edges_outgoing[lindex], a)

				if lindex == 0 {
					if length < best_from_s[end] {
						best_from_s[end] = length
					}
				}

			}
		}
	}
	pq := make(heap.PriorityQueue, 200)
	for i := 0; i < 200; i++ {
		pq[i] = &heap.Vertex{i, i, best_from_s[i]}
	}
	pq.HInit()

	return edges_outgoing, pq

}

func main() {
	//edges_outgoing, edges_incoming, pq := getE("course_2_p_2.txt")
	edges_outgoing, pq := getE("course_2_p_2.txt")
	var A [200]int

	for pq.Len() > 0 {
		v := pq.HPop().(*heap.Vertex)
		A[v.Nodeid] = v.Priority

		// check whether v -> E is a better path to E for all E leaving V
		for _, o := range edges_outgoing[v.Nodeid] {

			end := o[0]
			length := o[1]

			// find the item; TODO, IS THERE A BETTER WAY TO DO THIS???
			for i := range pq {
				if pq[i].Nodeid == end { // we found the node in the heap, now check it
					if A[v.Nodeid]+length < pq[i].Priority { // we have a better path
						pq.Update(pq[i], A[v.Nodeid]+length)
					}
				}
			}
		}
	}
	//fmt.Println(A)
	fmt.Println(A[6], ",", A[36], ",", A[58], ",", A[81], ",", A[98], ",", A[114], ",", A[132], ",", A[164], ",", A[187], ",", A[196])
}
