package main

import (
	"flag"
	"fmt"
	"log"
	"runtime/pprof"

	"os"
)

func square(n int) []int {
	s := make([]int, n)

	for i := 0; i < n; i++ {
		s[i] = i * i
	}
	return s
}

func square2(n int) []int {
	s := make([]int, n)

	for i := 0; i < n; i++ {
		s[i] = i*i + 1
	}
	return s
}

func main() {

	cpu := flag.String("cpu", "", "write cpu profile to cpu.prof file")
	mem := flag.String("mem", "", "write cpu profile to mem.prof file")

	flag.Parse()

	fmt.Println(*cpu, "-", *mem)

	if *cpu != "" {

		f, e := os.Create(*cpu)

		if e != nil {
			log.Fatal(e)
		}

		defer f.Close()
		err := pprof.StartCPUProfile(f)
		if err != nil {
			log.Fatal(err)
		}
		defer pprof.StopCPUProfile()
	}

	if *mem != "" {
		f, e := os.Create(*mem)

		if e != nil {
			log.Fatal(e)
		}

		defer f.Close()

		defer func() {
			if err := pprof.WriteHeapProfile(f); err != nil {
				log.Fatal(err)
			}
		}()
	}

	r1 := square(10)
	r2 := square2(10)

	fmt.Println(r1)
	fmt.Println(r2)
}
