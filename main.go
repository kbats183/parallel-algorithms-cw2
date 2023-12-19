package main

import (
	"github.com/kbats183/parallel-algorithms-cw2/parallel"
	"log"
	"runtime"
	"slices"
	"time"
)

func testImplementation(graph *CubeGridGraph, implementation BFSImplementation) time.Duration {
	start := time.Now()
	result := implementation.BFS(graph)
	t := time.Since(start)
	expectedDist := make([]int, graph.N())

	for f := 0; f < graph.N(); f++ {
		k := f % graph.Side
		j := f / graph.Side % graph.Side
		i := f / graph.Side / graph.Side
		expectedDist[f] = i + j + k
	}

	if !slices.Equal(result, expectedDist) {
		log.Fatalf("incorrect distances\nexpected: %v\nactual: %v", expectedDist, result)
	}

	return t
}

func testImplementationMultiple(graph *CubeGridGraph, implementation BFSImplementation, iterations int) time.Duration {
	var avg int64
	for i := 0; i < iterations; i++ {
		t := testImplementation(graph, implementation)
		avg += t.Milliseconds()
	}
	return time.Duration(avg/int64(iterations)) * time.Millisecond
}

func main() {
	runtime.GOMAXPROCS(4)
	g := CubeGridGraph{Side: 500}

	var sequentialBFS SequentialBFS
	var parallelBFS ParallelBFS

	ts := testImplementationMultiple(&g, &sequentialBFS, 2)
	log.Printf("Avg time for sequentional implementation %s", ts.String())

	parallel.PForDivider = 4
	for i := 0; i < 12; i++ {
		tp := testImplementationMultiple(&g, &parallelBFS, 2)
		log.Printf("Avg time for parallel implementation %s (%d)", tp.String(), parallel.PForDivider)
		log.Printf("Speed up for %v", float64(ts)/float64(tp))
		parallel.PForDivider *= 2
	}
}
