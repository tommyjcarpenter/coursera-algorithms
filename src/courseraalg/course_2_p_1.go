package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

const NODES = 875714

var t = 0

type Node struct {
	edge_to []int
	leader  int
	finish  int
}

func getnodes(fname string) (map[int]Node, map[int]Node) {

	b, _ := ioutil.ReadFile(fname)
	lines := strings.Split(string(b), "\n")

	nodes := make(map[int]Node, 0)
	nodes_rev := make(map[int]Node, 0)

	for i := 0; i < NODES; i++ {
		emptyarr := make([]int, 0)
		nodes[i] = Node{emptyarr, -1, -1}
		nodes_rev[i] = Node{emptyarr, -1, -1}
	}

	for _, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 {
			continue
		}

		line := strings.Fields(l)

		start, _ := strconv.Atoi(line[0])
		end, _ := strconv.Atoi(line[1])
		// the file doesn't use 0 indexing
		start--
		end--

		if start != end {

			// nodes
			n := nodes[start]
			n.edge_to = append(n.edge_to, end)
			nodes[start] = n
			//nodes rev
			nr := nodes_rev[end]
			nr.edge_to = append(nr.edge_to, start)
			nodes_rev[end] = nr
		}

	}
	return nodes, nodes_rev
}

func dfs(n map[int]Node, cur_index int, explored map[int]bool, sort_by_finish []int, leader_marker int, buildsortorder bool) {

	explored[cur_index] = true
	for _, i := range n[cur_index].edge_to {
		if !explored[i] {
			dfs(n, i, explored, sort_by_finish, leader_marker, buildsortorder)
		}
	}

	// update finishing time of n
	node := n[cur_index]
	node.finish = t
	node.leader = leader_marker
	n[cur_index] = node

	// store in reverse index. Only do this on first pass else we will be changing the array we are iterating over!!
	if buildsortorder {
		sort_by_finish[(NODES-1)-t] = cur_index
	}

	t++ // update global finishing time

}

func kosaraju(nodes map[int]Node, nodes_rev map[int]Node) {

	// loop 1
	explored := make(map[int]bool)

	t = 0
	for i := 0; i < NODES; i++ {
		explored[i] = false
	}

	sort_by_finish := make([]int, NODES)
	for i := NODES - 1; i >= 0; i-- {
		if !explored[i] {
			dfs(nodes_rev, i, explored, sort_by_finish, 666, true) // don't care about leader marker here
		}
	}

	// loop 2
	t = 0
	for i := 0; i < NODES; i++ {
		explored[i] = false
	}
	for i := 0; i < NODES; i++ {
		realindex := sort_by_finish[i]
		if !explored[realindex] {
			dfs(nodes, realindex, explored, sort_by_finish, realindex, false) // set leader marker to this index and for all current recurive calls called by that function itself
		}
	}
	// sanity check
	for i := 0; i < NODES; i++ {
		if !explored[i] || nodes[i].leader == -1 {
			fmt.Println("Node that should have been explored is dilapidated:   ", i, explored[i], nodes[i])
			os.Exit(1)
		}
	}
}

func main() {
	nodes, nodes_rev := getnodes("course_2_p_1.txt")
	kosaraju(nodes, nodes_rev)

	// build freq dict
	fd := make(map[int]int)
	for i := 0; i < NODES; i++ {
		l := nodes[i].leader
		if val, ok := fd[l]; ok {
			fd[l] = val + 1
		} else {
			fd[l] = 1
		}
	}
	// now reverse it, stolen from https://stackoverflow.com/questions/18695346/how-to-sort-a-mapstringint-by-its-values
	n := map[int][]int{}
	var a []int
	// build up a reverse map mapping value : all the orignial keys
	for k, v := range fd {
		n[v] = append(n[v], k)
	}
	// get just the values
	for k := range n {
		a = append(a, k)
	}
	// get the sorted values
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	for _, k := range a {
		// print the keys, sorted by value
		for _, s := range n[k] {
			fmt.Printf("%d, %d\n", s, k)
		}
	}
}
