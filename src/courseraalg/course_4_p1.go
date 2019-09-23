package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const NODES = 1000
const INF = 9999999999999

func getnodes(fname string) [NODES][NODES]int {

	b, _ := ioutil.ReadFile(fname)
	lines := strings.Split(string(b), "\n")

	var nodes [NODES][NODES]int

	for i := 0; i < NODES; i++ {
		for j := 0; j < NODES; j++ {
			if i != j {
				nodes[i][j] = INF
			} else {
				nodes[i][j] = 0
			}
		}
	}

	for _, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 {
			continue
		}

		line := strings.Fields(l)

		start, _ := strconv.Atoi(line[0])
		end, _ := strconv.Atoi(line[1])
		cost, _ := strconv.Atoi(line[2])
		// the file doesn't use 0 indexing
		start--
		end--

		nodes[start][end] = cost

	}
	return nodes
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

	file := ""
	f := os.Args[1]
	if f == "1" {
		file = "course_4_p1g1.txt"
	} else if f == "2" {
		file = "course_4_p1g2.txt"
	} else {
		file = "course_4_p1g3.txt"
	}

	nodes := getnodes(file)

	//we only need to keep around two
	// we will altnerate between index 0 and 1 to symbolically represent "k-1" and "k"
	var A [NODES][NODES][2]int
	last_index := 0
	cur_index := 1

	// init
	for i := 0; i < NODES; i++ {
		for j := 0; j < NODES; j++ {
			A[i][j][0] = nodes[i][j]
		}
	}

	for k := 0; k < NODES; k++ {
		for i := 0; i < NODES; i++ {
			for j := 0; j < NODES; j++ {
				A[i][j][cur_index] = min(A[i][j][last_index], A[i][k][last_index]+A[k][j][last_index])
			}
		}
		// swap what index we use
		if last_index == 0 {
			last_index = 1
			cur_index = 0
		} else {
			last_index = 0
			cur_index = 1
		}

	}

	min := INF
	for i := 0; i < NODES; i++ {
		for j := 0; j < NODES; j++ {

			if i == j && A[i][i][0] < 0 {
				fmt.Println("neg cycle!")
				os.Exit(1)
			}
			if i != j && A[i][j][0] < min {
				min = A[i][j][0]
			}
		}
	}
	fmt.Println(min)
}
