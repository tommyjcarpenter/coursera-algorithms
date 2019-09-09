package main

import (
	"courseraalg/heap"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const NODES = 1000

// A Tree is a binary tree with integer values.
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func getE(fname string) heap.PQMID {

	pq := make(heap.PQMID, NODES)

	b, _ := ioutil.ReadFile(fname)

	lines := strings.Split(string(b), "\n")

	for lindex, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 {
			continue
		}
		s := strings.Fields(l)

		weight, _ := strconv.Atoi(s[0])

		pq[lindex] = &heap.TreeVertex{lindex, weight, nil, nil}
	}

	pq.HInit()

	return pq
}

func BFS(T *heap.TreeVertex, depth int, mindepth int, maxdepth int) (int, int) {

	rmin := mindepth
	rmax := maxdepth

	if T.Left == nil && T.Right == nil {
		if depth < mindepth {
			rmin = depth
		}
		if depth > maxdepth {
			rmax = depth
		}
	} else {
		if T.Left != nil {
			mleft, mright := BFS(T.Left, depth+1, mindepth, maxdepth)
			if mleft < rmin {
				rmin = mleft
			}
			if mright > rmax {
				rmax = mright
			}

		}

		if T.Right != nil {
			mleft, mright := BFS(T.Right, depth+1, mindepth, maxdepth)
			if mleft < rmin {
				rmin = mleft
			}
			if mright > rmax {
				rmax = mright
			}

		}
	}

	return rmin, rmax
}

func main() {
	//adj_to, edges_incoming, pq := getE("course_2_p_2.txt")
	pq := getE("course_3_p3q1.txt")

	counter := NODES

	for pq.Len() > 1 {
		n1 := pq.ExtractMin().(*heap.TreeVertex)
		n2 := pq.ExtractMin().(*heap.TreeVertex)

		newe := &heap.TreeVertex{counter, n1.Priority + n2.Priority, n1, n2}

		//fmt.Println(n1, n2, newe)

		pq.AddElement(newe)
		counter++
	}

	T := pq.Pop().(*heap.TreeVertex)
	fmt.Println(BFS(T, 0, 999999, 1))

}
