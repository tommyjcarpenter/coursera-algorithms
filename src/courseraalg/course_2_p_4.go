package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

const arrsize = 1000000
const T = 10000 // one sided range size

func get(fname string) [1000000]int {
	var arr [1000000]int
	b, _ := ioutil.ReadFile(fname)
	lines := strings.Split(string(b), "\n")

	for lindex, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 {
			continue
		}

		line := strings.Fields(l)
		i, _ := strconv.Atoi(line[0])

		arr[lindex] = i
	}
	return arr
}

func key(i int) int {
	return i / T
}

func main() {
	start := time.Now()
	arr := get("course_2_p_4.txt")

	hash := make(map[int][]int)

	for _, i := range arr {
		key := key(i)
		hash[key] = append(hash[key], i)
	}

	solns := make(map[int]bool)
	for _, i := range hash { // for each bucket
		for _, x := range i { // for each item in this bucket
			// CHECK THE BIZZARO WORLD
			otherkey := key(-i[0]) // this should be all the ints within 10k of the opposite sign
			for b := otherkey - 1; b <= otherkey+1; b++ {
				for _, y := range hash[b] {
					s := x + y
					if x != y && s >= -T && s <= T {
						solns[s] = true
					}
				}
			}

		}
	}
	howmany := len(solns)
	fmt.Println(howmany)

	t := time.Now()
	fmt.Println(t.Sub(start))
}
