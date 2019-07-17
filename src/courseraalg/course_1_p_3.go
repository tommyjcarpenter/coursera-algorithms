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

func swap(a []int, i int, j int) {
	var temp int
	temp = a[i]
	a[i] = a[j]
	a[j] = temp
}

func partition_left(a []int, l int, rinclusive int) int {
	p := a[l]
	i := l + 1
	for j := l + 1; j <= rinclusive; j++ {
		if a[j] < p {
			swap(a, i, j)
			i++
		}
	}
	final_position := i - 1
	swap(a, l, final_position)
	return final_position
}

func partition_right(a []int, l int, rinclusive int) int {
	swap(a, l, rinclusive)
	return partition_left(a, l, rinclusive)
}

func partition_median_three(a []int, l, rinclusive int) int {
	one := a[l]
	last := a[rinclusive]

	// compute "middle" index
	length := rinclusive - l
	// if a is of length 20 and we are examining 5 to 15 inclusive, want mid 10
	// [0,1,2,3,4,[5],6,7,8,9,(10),11,12,13,14,[15],16,17,18,19,20]
	// 15 - 5 - 10, 10/2 =5, 5+5 = 10
	// if instead we are examining 5 to 14 inclusive, we want the 9
	// [0,1,2,3,4,[5],6,7,8,(9),10,11,12,13,[14],15,16,17,18,19,20]
	// 14 - 5 = 9, 9/2 = 4, 5+4 = 9
	// if length is odd, the above is already good.
	midindex := l + (length / 2)
	mid := a[midindex]

	// find the median
	if (one < mid && mid < last) || (last < mid && mid < one) {
		// the middle element is the median
		swap(a, l, midindex)
	} else if (mid < last && last < one) || (one < last && last < mid) {
		// the last element is the median
		swap(a, l, rinclusive)
	} // else the first element is the median, no need to do anything
	return partition_left(a, l, rinclusive)
}

func quicksort(a []int, l int, rinclusive int, partmethod int) int {
	if rinclusive <= l {
		return 0
	}
	fixed_index := -1
	if partmethod == 1 {
		fixed_index = partition_left(a, l, rinclusive)
	} else if partmethod == 2 {
		fixed_index = partition_right(a, l, rinclusive)
	} else if partmethod == 3 {
		fixed_index = partition_median_three(a, l, rinclusive)
	}
	count := rinclusive - l
	count += quicksort(a, l, fixed_index-1, partmethod)
	count += quicksort(a, fixed_index+1, rinclusive, partmethod)
	return count
}

func checkans(a []int) bool {
	good := true
	for i := 0; i < 10000; i++ {
		if a[i] != i+1 {
			good = false
		}
	}
	return good
}

func main() {
	nums, _ := readFile("course_1_p_3.txt")
	count := quicksort(nums, 0, len(nums)-1, 1)
	fmt.Println(checkans(nums), count)

	nums, _ = readFile("course_1_p_3.txt")
	count = quicksort(nums, 0, len(nums)-1, 2)
	fmt.Println(checkans(nums), count)

	nums, _ = readFile("course_1_p_3.txt")
	count = quicksort(nums, 0, len(nums)-1, 3)
	fmt.Println(checkans(nums), count)
}
