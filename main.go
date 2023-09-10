package main

import (
	"sync"
	"time"
)

type Tracker struct {
	name    string
	time    time.Duration
	count   int
	enabled bool
}

type PrettyPrinter struct {
	Headers []string
	Tracker []Tracker
}

var (
	operators      = []int8{-1, -2, -3, -4}
	numbers        = []int8{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10, 25, 50, 75, 100} // 177100 combinations; 5322360 permutations; 228904058880 equations
	testNumbers    = []int8{10, 10, 25, 50, 75, 100}                                                       // 1 combination; 360 permutations; 15482880 equations
	testNumbersV2  = []int8{1, 2, 10, 10, 25, 50, 75, 100}                                                 // 28 combinations; 10440 permutations; 449003520 equations
	permTrie       = NewTrie()                                                                             // Avoids duplicate permutations
	equationsCount int64                                                                                   // Tallies the total number of postfix equations
	testMap        sync.Map                                                                                // Stores postfix equations for each permutation
	wg             sync.WaitGroup
)

func main() {
	var enabled = map[string]bool{
		"combination": true,
		"permutation": true,
		"postfix":     true,
	}

	application(
		enabled,
		testNumbersV2,
	)
}

func application(enabled map[string]bool, numbers []int8) {
	var cResult [][]int8

	var combinationTime time.Duration
	var permutationsTime time.Duration
	var postfixTime time.Duration

	pp := PrettyPrinter{Headers: []string{"Name", "Time", "Count", "Enabled"}}

	start := time.Now()
	k := 6

	if enabled["combination"] {
		cResult = combinationRunner(numbers, k)
		combinationTime = time.Since(start)
	}

	if enabled["permutation"] {
		permutationRunner(cResult)
		permutationsTime = time.Since(start) - combinationTime
	}

	if enabled["postfix"] {
		postfixRunner(permTrie.getPaths())
		postfixTime = time.Since(start) - permutationsTime
	}

	pp.Tracker = append(pp.Tracker, Tracker{"Combinations", combinationTime, len(cResult), enabled["combination"]})
	pp.Tracker = append(pp.Tracker, Tracker{"Permutations", permutationsTime, permTrie.totalPaths(), enabled["permutation"]})
	pp.Tracker = append(pp.Tracker, Tracker{"Postfix", postfixTime, int(equationsCount), enabled["postfix"]})
	pp.Tracker = append(pp.Tracker, Tracker{"Total", time.Since(start), 0, true})

	prettyPrint(pp)
}
