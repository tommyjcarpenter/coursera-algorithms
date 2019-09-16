package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

const CAPACITY = 2000000
const ITEMS = 2000

var v [ITEMS]int
var w [ITEMS]int

type soln struct {
	cur_item      int
	remaining_cap int
}

var cache = make(map[soln]int)

func get_v_w(fname string) {

	b, _ := ioutil.ReadFile(fname)
	lines := strings.Split(string(b), "\n")

	for lindex, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 {
			continue
		}
		s := strings.Fields(l)

		v[lindex], _ = strconv.Atoi(s[0])
		w[lindex], _ = strconv.Atoi(s[1])
	}
}

func max(a int, b int) int {
	// return (a is bigger), a or (a is less, b)
	if a > b {
		return a
	}
	return b
}

func knapsack(cur_item int, remaining_cap int) int {

	s := soln{cur_item, remaining_cap}

	// cache hit
	cval, exists := cache[s]
	if exists {
		return cval
	}

	// base case
	if cur_item < 0 {
		return 0
	}

	dont_take := knapsack(cur_item-1, remaining_cap)

	// not feasible to choose current item
	if w[cur_item] > remaining_cap {
		return dont_take
	}

	take := v[cur_item] + knapsack(cur_item-1, remaining_cap-w[cur_item])

	best := max(dont_take, take)

	cache[s] = best

	return best

}

func main() {
	start := time.Now()
	get_v_w("course_3_p4q2.txt")
	solution := knapsack(ITEMS-1, CAPACITY)
	fmt.Println(solution)
	t := time.Now()
	fmt.Println(t.Sub(start))
}
