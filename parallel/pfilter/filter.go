package pfilter

import (
	"github.com/kbats183/parallel-algorithms-cw2/parallel"
)

func Filter(array []int, predicate func(index int, value int) bool) []int {
	if len(array) == 0 {
		return []int{}
	}

	dv := parallel.PForDivider
	if len(array) < dv {
		dv = len(array)
	}
	blockSize := (len(array) + dv - 1) / dv
	if blockSize < 1 {
		blockSize = 1
	}

	s := parallel.Pow2Size(dv)
	b := make([]int, (s<<1)-1)
	filterReduceImpl(array, 0, len(array), 0, b, predicate, blockSize)
	filterScanPropagate(b, 0, len(array), 0, 0, blockSize)
	result := make([]int, b[0])
	filterFinalizeImpl(array, 0, len(array), 0, b, 0, result, blockSize)

	return result
}

func FilterInplace(array []int, predicate func(index int, value int) bool, result []int) int {
	if len(array) == 0 {
		return 0
	}

	dv := parallel.PForDivider
	if len(array) < dv {
		dv = len(array)
	}
	blockSize := (len(array) + dv - 1) / dv
	if blockSize < 1 {
		blockSize = 1
	}

	s := parallel.Pow2Size(dv)
	b := make([]int, (s<<1)-1)
	filterReduceImpl(array, 0, len(array), 0, b, predicate, blockSize)
	filterScanPropagate(b, 0, len(array), 0, 0, blockSize)
	filterFinalizeImpl(array, 0, len(array), 0, b, 0, result, blockSize)

	return b[0]
}

func filterReduceImpl(array []int, l int, r int, x int, reduced []int, predicate func(index int, value int) bool, blockSize int) {
	if r-l <= blockSize {
		cnt := 0
		j := l
		for i := l; i < r; i++ {
			if predicate(i, array[i]) {
				array[j] = array[i]
				j++
				cnt++
			}
		}
		reduced[x] = cnt
	} else {
		m := (l + r) / 2
		parallel.Fork2Join(func() {
			filterReduceImpl(array, l, m, x*2+1, reduced, predicate, blockSize)
		}, func() {
			filterReduceImpl(array, m, r, x*2+2, reduced, predicate, blockSize)
		})
		reduced[x] = reduced[x*2+1] + reduced[x*2+2]
	}
}

func filterScanPropagate(
	reduced []int,
	l int,
	r int,
	x int,
	fromLeft int,
	blockSize int) {
	if r-l <= blockSize {
		reduced[x] += fromLeft
	} else {
		m := (l + r) / 2
		gg := reduced[x*2+1]
		parallel.Fork2Join(
			func() {
				filterScanPropagate(reduced, l, m, x*2+1, fromLeft, blockSize)
			}, func() {
				filterScanPropagate(reduced, m, r, x*2+2, fromLeft+gg, blockSize)
			})
		reduced[x] += fromLeft
	}
}

func filterFinalizeImpl(array []int, l int, r int, x int, reduced []int, prev int, result []int, blockSize int) {
	if r-l <= blockSize {
		for i := 0; i < (reduced[x] - prev); i++ {
			result[prev+i] = array[l+i]
		}
	} else {
		m := (l + r) / 2
		parallel.Fork2Join(func() {
			filterFinalizeImpl(array, l, m, x*2+1, reduced, prev, result, blockSize)
		}, func() {
			filterFinalizeImpl(array, m, r, x*2+2, reduced, reduced[x*2+1], result, blockSize)
		})
	}
}
