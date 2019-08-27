package main

import (
	"courseraalg/heap"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const JOBS = 10000

func getjobs(fname string) heap.MaxPQ {

	pq := make(heap.MaxPQ, JOBS)

	b, _ := ioutil.ReadFile(fname)

	lines := strings.Split(string(b), "\n")

	for lindex, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 {
			continue
		}
		s := strings.Fields(l)

		weight, _ := strconv.Atoi(s[0])
		length, _ := strconv.Atoi(s[1])

		pq[lindex] = &heap.Vertexf{lindex, lindex, weight, length, float64(weight) - float64(length)}
	}

	pq.HInit()

	return pq

}

func main() {
	pq := getjobs("course_3_p1q1.txt")

	sum := 0
	curr_completion_time := 0

	for pq.Len() > 0 {

		v := pq.ExtractMax().(*heap.Vertexf)
		curr_completion_time += v.Length
		sum += curr_completion_time * v.Weight
	}

	fmt.Println(sum)
}
