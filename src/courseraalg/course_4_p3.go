package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

const NODES = 33708
const INF = 9999999999999.9

type Vertex struct {
	x float64
	y float64
}

func getnodes(fname string) [NODES]Vertex {

	b, _ := ioutil.ReadFile(fname)
	lines := strings.Split(string(b), "\n")

	var nodes [NODES]Vertex

	for lindex, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 {
			continue
		}

		line := strings.Fields(l)

		x, _ := strconv.ParseFloat(line[1], 64)
		y, _ := strconv.ParseFloat(line[2], 64)

		nodes[lindex] = Vertex{x, y}
	}
	return nodes
}

func main() {
	nodes := getnodes("course_4_p3.txt")

	visited := make(map[int]bool)
	visited[0] = true

	tour_cost := 0.0
	cur_index := 0

	for i := 0; i < NODES-1; i++ { // -1 here because we start with 1 visited so we can only move to n-1 nodes
		min_d := INF
		min_index := -1
		for j := 0; j < NODES; j++ {
			_, exists := visited[j]
			if !exists {
				xdiff := nodes[cur_index].x - nodes[j].x
				ydiff := nodes[cur_index].y - nodes[j].y
				sqd := xdiff*xdiff + ydiff*ydiff
				if sqd < min_d {
					min_d = sqd
					min_index = j
				}
			}
		}
		tour_cost += math.Sqrt(min_d)
		visited[min_index] = true
		cur_index = min_index
	}
	fmt.Println(tour_cost)
	fmt.Println(cur_index) // final index before return

	// now add back to 0
	xdiff := nodes[cur_index].x - nodes[0].x
	ydiff := nodes[cur_index].y - nodes[0].y
	tour_cost += math.Sqrt(xdiff*xdiff + ydiff*ydiff)
	fmt.Println(tour_cost)
}
