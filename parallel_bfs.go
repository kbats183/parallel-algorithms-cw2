package main

import (
	"github.com/kbats183/parallel-algorithms-cw2/parallel"
	"github.com/kbats183/parallel-algorithms-cw2/parallel/pfilter"
	"sync/atomic"
)

type ParallelBFS struct{}

func (ParallelBFS) BFS(graph Graph) []int {
	d := make([]atomic.Int32, graph.N())
	d[0].Store(1)

	currentFrontier := []int{0}

	for i := 0; i < graph.N(); i++ {
		positionsC := parallel.Map(currentFrontier, func(v int) int {
			c := 0
			for _, u := range graph.Edge(v) {
				if d[u].Load() == 0 {
					c++
				}
			}
			return c
		})
		positions := parallel.BlockedScan(positionsC)
		//fmt.Printf("d: %v\n", atomicInt32ArrayToArray(d))
		//fmt.Printf("f: %v\n", currentFrontier)
		//fmt.Printf("p: %v\n", positionsC)
		//fmt.Printf("p: %v\n", positions)
		//fmt.Printf("\n")

		newFrontier := make([]int, positions[len(positions)-1])
		//currentFrontier = currentFrontier

		parallel.PFor(len(currentFrontier), func(index int) {
			v := currentFrontier[index]
			dv := d[v].Load() + 1
			idx := 0
			delta := 0
			if index > 0 {
				delta = positions[index-1]
			}
			for _, u := range graph.Edge(v) {
				if d[u].CompareAndSwap(0, dv) {
					newFrontier[delta+idx] = u
					idx++
				}
			}
		})

		currentFrontier = pfilter.Filter(newFrontier, func(index int, value int) bool {
			return value != 0
		})
		if len(currentFrontier) == 0 {
			break
		}
	}

	return atomicInt32ArrayToArray(d)
}

func atomicInt32ArrayToArray(array []atomic.Int32) []int {
	result := make([]int, len(array))
	parallel.PFor(len(array), func(index int) {
		result[index] = int(array[index].Load()) - 1
	})
	return result
}
