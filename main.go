package main

import (
	"flag"
	"log"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/pprof"
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
	operators                 = []int8{-1, -2, -3, -4}
	numbers0                  = []int8{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10, 25, 50, 75, 100} // 177100 combinations; 5322360 permutations; 228904058880 equations
	numbers1                  = []int8{10, 10, 25, 50, 75, 100}                                                       // 1 combination; 360 permutations; 15482880 equations
	numbers2                  = []int8{1, 2, 10, 10, 25, 50, 75, 100}                                                 // 28 combinations; 10440 permutations; 449003520 equations
	combinationTrie           = NewTrie()                                                                             // Avoids duplicate combinations
	permutationTrie           = NewTrie()                                                                             // Avoids duplicate permutations
	equationsCount            int64                                                                                   // Tallies the total number of postfix equations
	permutationCount          int64                                                                                   // Tallies the total number of permutations
	combinationPermutationMap sync.Map                                                                                // Stores permutations by combination
	permutationPostfixMap     sync.Map                                                                                // Stores postfix equations by permutation
	solutionsMap              sync.Map                                                                                // Running sum of each three-digit solution found
	wg                        sync.WaitGroup
	testMap                   = make(map[*[]int8][][]int8)
	cpuprofile                = flag.String("cpuprofile", "", "write cpu profile to `file`")
	memprofile                = flag.String("memprofile", "", "write memory profile to `file`")
)

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	var config = map[string]bool{
		"combination": true,
		"permutation": true,
		"postfix":     true,
		"print":       true,
	}

	application(
		config,
		numbers2,
	)

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		runtime.GC()    // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}

func application(enabled map[string]bool, numbers []int8) {
	var combinationTime time.Duration
	var permutationsTime time.Duration
	var postfixTime time.Duration

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

	//if enabled["permutation"] {
	//	permutationRunner(combinationTrie.getPaths())
	//	permutationsTime = time.Since(start) - combinationTime
	//}
	//
	//if enabled["postfix"] {
	//	combinationPermutationMap.Range(func(key, value interface{}) bool {
	//		testMap[key.(*[]int8)] = value.([][]int8)
	//		return true
	//	})
	//
	//	postfixRunner()
	//	postfixTime = time.Since(start) - permutationsTime
	//}

	if enabled["permutation"] {
		permutationRunner2(combinationTrie.getPaths())
		permutationsTime = time.Since(start) - combinationTime
	}

	if enabled["postfix"] {
		postfixRunner2(permutationTrie.getPaths())
		postfixTime = time.Since(start) - permutationsTime
	}

	totalTime := time.Since(start)

	if enabled["print"] {
		pp.Tracker = append(pp.Tracker, Tracker{
			name:    "Combinations",
			time:    combinationTime,
			count:   combinationTrie.totalPaths(),
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

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		runtime.GC()    // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}
