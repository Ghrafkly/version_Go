package main

import (
	"log"
	"os"
	"runtime/pprof"
)

func cpuProfiler() func() {
	f, err := os.Create(*cpuprofile)
	if err != nil {
		log.Fatal(err)
	}
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal(err)
	}
	return func() {
		pprof.StopCPUProfile()
		f.Close()
	}
}

func memProfiler() func() {
	f, err := os.Create(*memprofile)
	if err != nil {
		log.Fatal(err)
	}
	return func() {
		pprof.WriteHeapProfile(f)
		f.Close()
	}
}
