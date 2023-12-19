package main

import (
	"github.com/kbats183/parallel-algorithms-cw2/parallel"
	"sync/atomic"
)

type ParallelBFS struct{}

func (ParallelBFS) BFS(graph Graph) []int {
	d := make([]int, graph.N())
	dd := make([]atomic.Int32, graph.N())
	dd[0].Store(1)

	arr1 := make([]int, graph.N())
	//positionsC := make([]int, graph.N())
	//newFrontier := make([]int, graph.N())
	currentFrontier := []int{0}

	for i := 0; i < graph.N(); i++ {
		positionsC := parallel.Map(currentFrontier, func(v int) int {
			c := 0
			for _, u := range graph.Edge(v) {
				if d[u] == 0 {
					c++
				}
			}
			return c
		})
		positions := parallel.BlockedScan(positionsC)

		newFrontier := make([]int, positions[len(positions)-1])
		//currentFrontier = currentFrontier

		parallel.PFor(len(currentFrontier), func(index int) {
			v := currentFrontier[index]
			dv := d[v] + 1
			idx := 0
			delta := 0
			if index > 0 {
				delta = positions[index-1]
			}
			for _, u := range graph.Edge(v) {
				if dd[u].CompareAndSwap(0, int32(dv)) {
					d[u] = dv
					newFrontier[delta+idx] = u
					idx++
				}
			}
		})

		//newSize := pfilter.Filter(newFrontier[:newFrontierLen], func(index int, value int) bool {
		//	return value != 0
		//}, arr1)
		newSize := sequenceFilter(newFrontier, func(index int, value int) bool {
			return value != 0
		}, arr1)
		currentFrontier = arr1[:newSize]
		if len(currentFrontier) == 0 {
			break
		}
	}

	return d
}

func atomicInt32ArrayToArray(array []atomic.Int32) []int {
	result := make([]int, len(array))
	parallel.PFor(len(array), func(index int) {
		result[index] = int(array[index].Load()) - 1
	})
	return result
}

func sequenceFilter(array []int, predicate func(index int, value int) bool, result []int) int {
	cnt := 0
	j := 0
	for i := 0; i < len(array); i++ {
		if predicate(i, array[i]) {
			result[j] = array[i]
			j++
			cnt++
		}
	}
	return j
}
