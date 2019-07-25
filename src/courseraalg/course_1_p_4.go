package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

/* We dont really need a full adjacency list.
This algorihm works on the edges, and the result depends on the number of edges
so it's not necessary to maintain a list of the correct lists of other nodes each vertex touches
We use a simple hashmap of ints so that deletion is easy
*/
func getVE(fname string) (map[int]bool, [][2]int) {

	b, _ := ioutil.ReadFile(fname)

	lines := strings.Split(string(b), "\n")

	vertices := make(map[int]bool)
	edges := make([][2]int, 0)

	for _, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 {
			continue
		}

		line := strings.Fields(l)

		whatrow := 0
		for jindex, j := range line {
			i, _ := strconv.Atoi(j)
			if jindex == 0 {
				whatrow = i
			} else {

				// we don't want to doublecount edges
				if whatrow < i {
					edges = append(edges, [...]int{whatrow, i})
				}
			}
		}
		vertices[whatrow] = true
	}

	return vertices, edges
}

// stolen from https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang/37335777#37335777
// note, we assume order doesn't matter, there seems to be no need to keep the edges in order
func remove(s [][2]int, i int) [][2]int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

//we will collapse v2 into v1:
// 1. anything v2 points to, v1 should point to that thing instead
// 2. anything tht pointed TO v2 should point to v1 instead
// 3. remove self loops
// 4. delete v2 from the vertex hash
func merge_vertices(vertices map[int]bool, edges [][2]int, v1 int, v2 int) [][2]int {

	newedges := make([][2]int, 0)

	for eindex, e := range edges {
		// step 1
		if e[0] == v2 {
			edges[eindex][0] = v1
		}
		// step 2
		if e[1] == v2 {
			edges[eindex][1] = v1
		}
	}

	// step 3
	for _, e := range edges {
		if e[0] != e[1] {
			newedges = append(newedges, e)
		}
	}

	// step 4
	delete(vertices, v2)

	return newedges
}

func contraction(vertices map[int]bool, edges [][2]int) (int, map[int]bool, [][2]int) {

	//fmt.Println("conraction", len(vertices), len(edges))

	E := len(edges)
	if len(vertices) == 2 {
		return E, vertices, edges
	}

	// choose an edge at random to obliterate

	eindex := rand.Intn(E) // pseudo-random number in [0,n)

	v1 := edges[eindex][0]
	v2 := edges[eindex][1]

	// removce the edge
	edges = remove(edges, eindex)

	// merge the vertices
	edges = merge_vertices(vertices, edges, v1, v2)

	return contraction(vertices, edges)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	min := 20000
	mincut := 0
	for i := 0; i < 1000; i++ {
		vertices, edges := getVE("course_1_p_4.txt")
		mincut, vertices, edges = contraction(vertices, edges)
		if mincut < min {
			fmt.Println("\n\n new min found", min, "->", mincut)
			min = mincut
			fmt.Println(vertices)
			fmt.Println(edges)
			fmt.Println(min)
		}
	}

}
