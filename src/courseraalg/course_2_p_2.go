package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getE(fname string) map[int][][2]int {

	b, _ := ioutil.ReadFile(fname)

	lines := strings.Split(string(b), "\n")

	edges := make(map[int][][2]int, 0)

	for lindex, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 {
			continue
		}
		edges[lindex] = make([][2]int, 0)

		line := strings.Fields(l)
		for jindex, j := range line {
			if jindex > 0 { // first column is a repeat of the row index
				s := strings.Split(j, ",")
				end, _ := strconv.Atoi(s[0])
				length, _ := strconv.Atoi(s[1])
				a := [2]int{end - 1, length} // -1 to 0index

				elist := edges[lindex]
				elist = append(elist, a)
				edges[lindex] = elist
			}
		}
	}

	return edges
}

func dik(edges map[int][][2]int) [200]int {
	var covered [200]bool
	var A [200]int
	covered[0] = true
	A[0] = 0
	for i := 1; i < 200; i++ {
		A[i] = 1000000
		covered[i] = false
	}

	for iteration := 1; iteration < 200; iteration++ {
		lowestscore := 1000000
		target_end := -1

		for potential := 0; potential < 200; potential++ {
			if covered[potential] {
				for _, o := range edges[potential] {
					end := o[0]
					length := o[1]
					if !covered[end] && A[potential]+length < lowestscore {
						lowestscore = A[potential] + length
						target_end = end
					}
				}
			}
		}
		covered[target_end] = true
		A[target_end] = lowestscore
	}
	return A
}

func main() {
	edges := getE("course_2_p_2.txt")
	paths := dik(edges)
	//7,37,59,82,99,115,133,165,188,197
	fmt.Println(paths[6], ",", paths[36], ",", paths[58], ",", paths[81], ",", paths[98], ",", paths[114], ",", paths[132], ",", paths[164], ",", paths[187], ",", paths[196])
}
