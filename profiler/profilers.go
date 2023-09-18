package profiler

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
)

var (
	cpuprofile = flag.String("cpuprofile", "profiler/cpu.prof", "write cpu profile to `file`")
	memprofile = flag.String("memprofile", "profiler/mem.prof", "write memory profile to `file`")
)

func CPUProfiler() func() {
	if *cpuprofile != "" {
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

	return func() {}
}

func MemProfiler() func() {
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal(err)
		}
		return func() {
			pprof.WriteHeapProfile(f)
			f.Close()
		}
	}

	return func() {}
}
