package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var N int = -1

type Clause struct {
	negatex bool
	x       int
	negatey bool
	y       int
}

func getclauses(fname string) []Clause {

	b, _ := ioutil.ReadFile(fname)
	lines := strings.Split(string(b), "\n")

	clauses := make([]Clause, N)

	for lindex, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 {
			continue
		}

		line := strings.Fields(l)

		x, _ := strconv.Atoi(line[0])
		y, _ := strconv.Atoi(line[1])
		negx := false
		negy := false
		if x < 0 {
			negx = true
			x = -x
		}
		if y < 0 {
			negy = true
			y = -y
		}
		x -= 1
		y -= 1
		clauses[lindex] = Clause{negx, x, negy, y}

	}
	return clauses
}

func remove(s []Clause, i int) []Clause {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

func reduceclauses(clauses []Clause) []Clause {
	/*
		From: https://www.coursera.org/learn/algorithms-npcomplete/discussions/weeks/4/threads/Cl-n9enMEeavDRL9DbMJZA
		   We can effectivly reduce number of clauses in each task. If any variable has only one representation (negated or not) in all clauses, than we can remove all this clauses from the task. Because we can easely make all such clause to be valid, by setting this variable to true or false.

		   Number of clauses after reduction:

		   2sat1.txt - 6
		   2sat2.txt - 57
		   2sat3.txt - 295
		   2sat4.txt - 11
		   2sat5.txt - 101
		   2sat6.txt - 26
	*/
	fmt.Println("Pruning", len(clauses))

	canreduce := make(map[int]bool)
	states := make(map[int]bool)

	for i := 0; i < N; i++ {
		canreduce[i] = true
	}

	for c := 0; c < len(clauses); c++ {
		xstate, xok := states[clauses[c].x]
		if !xok {
			states[clauses[c].x] = clauses[c].negatex
		} else if clauses[c].negatex != xstate {
			canreduce[clauses[c].x] = false
		}

		ystate, yok := states[clauses[c].y]
		if !yok {
			states[clauses[c].y] = clauses[c].negatey
		} else if clauses[c].negatey != ystate {
			canreduce[clauses[c].y] = false
		}

	}

	somethingchanged := true
	for somethingchanged {
		fmt.Println("reducing", len(clauses))
		somethingchanged = false
		for c := 0; c < len(clauses); c++ {
			xok, _ := canreduce[clauses[c].x]
			yok, _ := canreduce[clauses[c].y]
			if xok || yok {
				clauses = remove(clauses, c)
				somethingchanged = true
			}
		}
	}

	return clauses

}

func check_soln(A []bool, clauses []Clause) []int {
	failures := make([]int, 0)
	for c := 0; c < len(clauses); c++ {
		C := clauses[c]
		x := A[C.x]
		if C.negatex {
			x = !x
		}
		y := A[C.y]
		if C.negatey {
			y = !y
		}

		if !(x || y) { // just take the first constraint that's failing
			failures = append(failures, c)
		}
	}
	return failures
}

func run(clauses []Clause) {
	// log_2(100000) ~= 17
	for start := 0; start < int(math.Floor(math.Log2(float64(N)))); start++ {
		fmt.Println("iteration", start)

		// init guess to random answer
		A := make([]bool, N)
		for i := 0; i < N; i++ {
			A[i] = false
			if rand.Intn(2) == 1 {
				A[i] = true
			}
		}
		// 2*n*n is just too big for these even if C is very very small, when we fail
		//400,000 * 400,000 = 60,000,000,000
		for i := 0; i < 2*N; i++ {
			failures := check_soln(A, clauses)

			if len(failures) == 0 {
				fmt.Println("Solution found!")
				os.Exit(0)
			} else {
				c := failures[rand.Intn(len(failures))] // pick a random bad constraint
				C := clauses[c]
				if rand.Intn(2) == 1 { // randomly pick which var
					A[C.x] = !A[C.x]
				} else {
					A[C.y] = !A[C.y]
				}
			}

		}
	}
	fmt.Println("NO SOLUTION FOUND")
}

func main() {

	rand.Seed(time.Now().UnixNano())
	file := ""
	f := os.Args[1]
	switch f {
	case "1":
		{
			file = "course_4_p4q1.txt"
			N = 100000
		}
	case "2":
		{
			file = "course_4_p4q2.txt"
			N = 200000
		}
	case "3":
		{
			file = "course_4_p4q3.txt"
			N = 400000
		}
	case "4":
		{
			file = "course_4_p4q4.txt"
			N = 600000
		}
	case "5":
		{
			file = "course_4_p4q5.txt"
			N = 800000
		}
	case "6":
		{
			file = "course_4_p4q6.txt"
			N = 1000000
		}
	}

	clauses := getclauses(file)

	for r := 0; r < 100; r++ {
		clauses = reduceclauses(clauses)
	}

	fmt.Println(len(clauses))

	run(clauses)

}
