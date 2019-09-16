package main

import (
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

func getV(fname string) [NODES + 1]int {

	var V [NODES + 1]int
	V[0] = 0

	b, _ := ioutil.ReadFile(fname)

	lines := strings.Split(string(b), "\n")

	for lindex, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 {
			continue
		}
		s := strings.Fields(l)

		weight, _ := strconv.Atoi(s[0])

		V[lindex+1] = weight
	}
	return V
}

func max(a int, b int) int {
	// return (a is bigger), a or (a is less, b)
	if a > b {
		return a
	}
	return b
}

func main() {
	//adj_to, edges_incoming, pq := getE("course_2_p_2.txt")
	V := getV("course_3_p3q3.txt")

	var A [NODES + 1]int
	var Inc [NODES + 1]bool
	for i := 0; i < NODES+1; i++ {
		Inc[i] = false
	}

	A[0] = V[0] // already set to 0
	A[1] = V[1]

	for i := 2; i < NODES+1; i++ {
		A[i] = max(A[i-2]+V[i], A[i-1])
	}

	i := NODES
	for i >= 2 {
		if A[i-1] > A[i-2]+V[i] { // when i = 2, if a[1] > a[0]+v[2], meaning if we chose a[1], we cannot choose a[2] (or a[0]), therefore inc[2] is false
			Inc[i] = false
			i--
		} else {
			Inc[i] = true
			i -= 2
		}
	}

	// the backtracking alg doesn't cover this edge case;
	// intuition here is that if we didn't include the 2nd element, or stated alternativelhy if we DID include the third, then we would also include
	// the first because the solution would be strictly better
	if Inc[3] {
		Inc[1] = true
	}

	//1, 2, 3, 4, 17, 117, 517, and 997
	var ans = [...]int{1, 2, 3, 4, 17, 117, 517, 997}
	for _, i := range ans {
		if Inc[i] {
			fmt.Print(1)
		} else {
			fmt.Print(0)
		}
	}

}
