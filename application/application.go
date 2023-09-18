package application

import (
	"sync"
	"time"
	algos "version_Go/algorithms"
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
	combinationTrie           = algos.NewTrie()
	equationsCount            int64    // Tallies the total number of postfix equations
	permutationCount          int64    // Tallies the total number of permutations
	combinationPermutationMap sync.Map // Stores permutations by combination
	permutationPostfixMap     sync.Map // Stores postfix equations by permutation
	solutionsMap              sync.Map // Running sum of each three-digit solution found
	wg                        sync.WaitGroup
	testMap                   = make(map[*[]int8][][]int8)
)

func Application(enabled map[string]bool, numbers []int8) {
	var combinationTime time.Duration
	var permutationsTime time.Duration
	var postfixTime time.Duration

	algos.CombinationTrie = combinationTrie

	for i := 101; i < 1000; i++ {
		solutionsMap.Store(i, 0)
	}

	pp := PrettyPrinter{Headers: []string{"Name", "Time", "Count", "Enabled"}}

	start := time.Now()
	k := 6

	if enabled["combination"] {
		combinationRunner(numbers, k)
		combinationTime = time.Since(start)
	}

	if enabled["permutation"] {
		permutationRunner(combinationTrie.GetPaths())
		permutationsTime = time.Since(start) - combinationTime
	}

	if enabled["postfix"] {
		combinationPermutationMap.Range(func(key, value interface{}) bool {
			testMap[key.(*[]int8)] = value.([][]int8)
			return true
		})

		postfixRunner()
		postfixTime = time.Since(start) - permutationsTime
	}

	totalTime := time.Since(start)

	if enabled["print"] {
		pp.Tracker = append(pp.Tracker, Tracker{
			name:    "Combinations",
			time:    combinationTime,
			count:   combinationTrie.TotalPaths(),
			enabled: enabled["combination"],
		})
		pp.Tracker = append(pp.Tracker, Tracker{
			name:    "Permutations",
			time:    permutationsTime,
			count:   int(permutationCount),
			enabled: enabled["permutation"],
		})
		pp.Tracker = append(pp.Tracker, Tracker{
			name:    "Postfix",
			time:    postfixTime,
			count:   int(equationsCount),
			enabled: enabled["postfix"],
		})
		pp.Tracker = append(pp.Tracker, Tracker{
			name:    "Total",
			time:    totalTime,
			enabled: true,
		})

		prettyPrint(pp)
	}
}
