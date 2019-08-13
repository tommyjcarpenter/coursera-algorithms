package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getarray(fname string) [10000]int {
	var a [10000]int

	b, _ := ioutil.ReadFile(fname)
	lines := strings.Split(string(b), "\n")

	for lindex, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 {
			continue
		}

		line := strings.Fields(l)
		i, _ := strconv.Atoi(line[0])
		a[lindex] = i

	}
	return a
}

// is there an operators.lessthan(a,b)?
func min(a int, b int) bool {
	return a < b
}

func max(a int, b int) bool {
	return a > b
}

func heap_insert(heap []int, item int, operator func(int, int) bool) []int {
	heap = append(heap, item)
	i := len(heap) - 1
	for {
		if i == 0 {
			break
		}
		// if we have 3 nodes, 1 > 2,3, if i == 1, meaning we are at node 2, 1/2 =1
		// if we have 3 nodes, 1 > 2,3, if i == 2, meaning we are at node 3, 2/2-1 =0
		p := i / 2    // floor is taken care by integer addressing
		if i%2 == 0 { // even, parent == i/2 , but since 0 indexing, -1
			p--
		}
		if operator(heap[p], heap[i]) { // operator is not violated
			break
		} else {
			heap[p], heap[i] = heap[i], heap[p]
			i = p
		}
	}
	return heap
}

func extract_root(heap []int, operator func(int, int) bool) ([]int, int) {
	return_item := heap[0]
	lastindex := len(heap) - 1
	heap[0] = heap[lastindex]
	heap = heap[:lastindex]
	newlastindex := len(heap) - 1

	// bubble down
	i := 0
	for {
		l := 2*i + 1 // last +1 because of 0 indexing
		r := l + 1
		if l > newlastindex { // no childen
			break
		} else { // at least one child
			if r > newlastindex { // no right child, only need to check left
				if operator(heap[i], heap[l]) {
					break
				} else {
					heap[i], heap[l] = heap[l], heap[i]
					i = l
				}
			} else { // has two children
				if operator(heap[i], heap[l]) && operator(heap[i], heap[r]) { // neither child violates
					break
				} else { // have two children with at least one violation, will do a swap either way
					if operator(heap[l], heap[r]) { // in this case, we want the operator in line with operator, not !operator! if operator is min, and left is smaller, we want to swap with left
						heap[i], heap[l] = heap[l], heap[i]
						i = l
					} else { // right smaller
						heap[i], heap[r] = heap[r], heap[i]
						i = r
					}
				}
			}
		}
	}
	return heap, return_item
}

func min_heap_insert(heap []int, item int) []int {
	return heap_insert(heap, item, min)
}

func max_heap_insert(heap []int, item int) []int {
	return heap_insert(heap, item, max)
}

func min_extract_root(heap []int) ([]int, int) {
	return extract_root(heap, min)
}

func max_extract_root(heap []int) ([]int, int) {
	return extract_root(heap, max)
}

func main() {

	a := getarray("course_2_p_3.txt")

	lowheap := make([]int, 0)
	highheap := make([]int, 0)
	e := -1
	median := -1
	mediansum := 0

	for i := 0; i < 10000; i++ {
		if i == 0 {
			lowheap = max_heap_insert(lowheap, a[0])
		} else if a[i] <= lowheap[0] { // should go into low heap, root is max
			lowheap = max_heap_insert(lowheap, a[i])
			// resolve imbalancing
			if len(lowheap)-len(highheap) == 2 {
				lowheap, e = max_extract_root(lowheap)
				highheap = min_heap_insert(highheap, e)
			}
		} else {
			highheap = min_heap_insert(highheap, a[i])
			if len(highheap)-len(lowheap) == 2 {
				highheap, e = min_extract_root(highheap)
				lowheap = max_heap_insert(lowheap, e)
			}
		}
		if len(lowheap)-len(highheap) > 1 || len(highheap)-len(lowheap) > 1 {
			fmt.Println("asdfasdf")
		}
		if len(lowheap) == len(highheap) {
			median = lowheap[0] // question wants k/2 when even, so if k = 10, 5, in 0 indexing 4 so 0123[4], 56789
		} else if len(lowheap) < len(highheap) {
			median = highheap[0]
		} else {
			median = lowheap[0]
		}
		mediansum += median
	}
	fmt.Println(len(lowheap), len(highheap), median)
	fmt.Println(mediansum)

}
