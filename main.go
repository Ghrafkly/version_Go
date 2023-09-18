package main

import (
	"flag"
	_ "net/http/pprof"
	_ "version_Go/algorithms"
	app "version_Go/application"
	"version_Go/profiler"
)

var (
	// 177100 combinations; 5322360 permutations; 228904058880 equations
	numbers0 = []int8{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10, 25, 50, 75, 100}

	// 1 combination; 360 permutations; 15482880 equations
	numbers1 = []int8{10, 10, 25, 50, 75, 100}

	// 28 combinations; 10440 permutations; 449003520 equations
	numbers2 = []int8{1, 2, 10, 10, 25, 50, 75, 100}
)

func main() {
	flag.Parse()
	profiler.CPUProfiler()

	var config = map[string]bool{
		"combination": true,
		"permutation": true,
		"postfix":     true,
		"print":       true,
	}

	app.Application(config, numbers2)

	profiler.MemProfiler()
}
