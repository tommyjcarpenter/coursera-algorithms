package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// stolen from https://stackoverflow.com/questions/9862443/golang-is-there-a-better-way-read-a-file-of-integers-into-an-array
func readFile(fname string) (nums []int, err error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")
	// Assign cap to avoid resize on every append.
	nums = make([]int, 0, len(lines))

	for _, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 {
			continue
		}
		// Atoi better suits the job when we know exactly what we're dealing
		// with. Scanf is the more general option.
		n, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}

	return nums, nil
}

func MergeSort(a []int) ([]int, int) {
	n := len(a)
	sorted := make([]int, n)

	if n <= 1 {
		return a, 0
	}

	left_sorted, left_inversions := MergeSort(a[:n/2])
	right_sorted, right_inversions := MergeSort(a[n/2:])

	l := 0
	r := 0
	num_split := 0
	for k := 0; k < n; k++ {
		// no elements left in left array, copy from right, but no inversions since all bigger than left (empty)
		if l >= len(left_sorted) {
			sorted[k] = right_sorted[r]
			r++
			// no elements in right array, copy from left, not an inversion
		} else if r >= len(right_sorted) {
			sorted[k] = left_sorted[l]
			l++
		} else {
			// left is smallar, not an inversion
			if left_sorted[l] <= right_sorted[r] {
				sorted[k] = left_sorted[l]
				l++
				// this is the inversion case
			} else {
				sorted[k] = right_sorted[r]
				r++
				// since we are copying from right, it is smaller than all remaining elements in left, which means its an inversion against all of them
				num_split += len(left_sorted) - l
			}
		}
	}

	return sorted, left_inversions + right_inversions + num_split
}

func main() {
	nums, _ := readFile("course_1_p_2.txt")
	_, count := MergeSort(nums)
	fmt.Println(count)
}
