package main

import (
	"github.com/rodaine/table"
	"time"

	"github.com/fatih/color"
)

type Tracker struct {
	name  string
	time  time.Duration
	count int
}

type PrettyPrinter struct {
	Headers []string
	Tracker []Tracker
}

var (
	operators      = []int8{-1, -2, -3, -4}
	numbers        = []int8{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10, 25, 50, 75, 100} // 177100 combinations; 5322360 permutations; 228904058880 equations
	testNumbers    = []int8{10, 10, 25, 50, 75, 100}                                                          // 1 combination; 360 permutations; 15482880 equations
	testNumbersv2  = []int8{1, 2, 10, 10, 25, 50, 75, 100}                                                    // 28 combinations; 10440 permutations; 449003520 equations
	permTrie       = NewTrie()                                                                                // Avoids duplicate permutations
	equationsCount int
	permutationMap = make(map[int][][]int8) // Stores postfix equations for each permutation
)

func main() {
	pp := PrettyPrinter{Headers: []string{"Name", "Time", "Count"}}
	start := time.Now()
	k := 6

	// 1. Combinations //
	cResult := combinations(testNumbersv2, k)
	combinationTime := time.Since(start)
	// =============== //

	// 2. Permutations //
	for _, combination := range cResult {
		permutations(combination)
	}
	permutationsTime := time.Since(start) - combinationTime
	// =============== //

	// 3. Postfix //
	for i, permutation := range permTrie.getPaths() {
		result := postfix(permutation)
		permutationMap[i] = result

		equationsCount += len(result)
	}
	postfixTime := time.Since(start) - permutationsTime
	// =============== //

	//// 1. Combinations //
	//cResult := combinationGenerator(testNumbersv2, k)
	//combinationTime := time.Since(start)
	//// =============== //
	//
	//// 2. Permutations //
	//permutationGenerator(cResult)
	//permutationsTime := time.Since(start) - combinationTime
	//// =============== //
	//
	//// 3. Postfix //
	//postfixGenerator(permTrie.getPaths())
	//postfixTime := time.Since(start) - permutationsTime
	//// =============== //

	pp.Tracker = append(pp.Tracker, Tracker{"Combinations", combinationTime, len(cResult)})
	pp.Tracker = append(pp.Tracker, Tracker{"Permutations", permutationsTime, permTrie.totalPaths()})
	pp.Tracker = append(pp.Tracker, Tracker{"Postfix", postfixTime, equationsCount})
	pp.Tracker = append(pp.Tracker, Tracker{"Total", time.Since(start), 0})

	prettyPrint(pp)
}

func prettyPrint(pp PrettyPrinter) {
	headerFmt := color.New(color.FgGreen, color.Bold, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	headers := make([]interface{}, len(pp.Headers))
	for i, h := range pp.Headers {
		headers[i] = h
	}

	tbl := table.New(headers...)

	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt).WithPadding(5)

	for _, t := range pp.Tracker {
		tbl.AddRow(t.name, t.time, t.count)
	}

	tbl.Print()
}

func combinationGenerator(nums []int8, k int) [][]int8 {
	return combinations(nums, k)
}

func permutationGenerator(combinations [][]int8) {
	for _, combination := range combinations {
		permutations(combination)
	}
}

func postfixGenerator(permutations [][]int8) {
	for i, permutation := range permutations {
		result := postfix(permutation)
		permutationMap[i] = result

		equationsCount += len(result)
	}
}
