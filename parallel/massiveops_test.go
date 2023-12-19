package parallel

import (
	"slices"
	"testing"
)

func TestMap(t *testing.T) {
	a := make([]int, 10)
	c := make([]int, 10)
	for i := 0; i < 10; i++ {
		a[i] = i * i
		c[i] = a[i] + 1
	}
	b := Map(a, func(v int) int {
		return v + 1
	})

	if !slices.Equal(b, c) {
		t.Fail()
	}
}

func TestReduceImpl(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := make([]int, 7)
	c := []int{10, 3, 7, 1, 2, 3, 4}
	ReduceImpl(a, 0, len(a), 0, b, 1)

	if !slices.Equal(b, c) {
		t.Fail()
	}
}

func TestScanPropagate(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := make([]int, 7)
	c := make([]int, 4)
	d := []int{1, 3, 6, 10}
	ReduceImpl(a, 0, len(a), 0, b, 1)
	scanPropagate(b, 0, len(a), 0, 0, c, 1)

	if !slices.Equal(c, d) {
		t.Fail()
	}
}

func TestScan(t *testing.T) {
	if !slices.Equal(Scan([]int{1}), []int{1}) {
		t.Error("Test 0 failed")
	}
	if !slices.Equal(Scan([]int{1, 2}), []int{1, 3}) {
		t.Error("Test 0.5 failed")
	}
	if !slices.Equal(Scan([]int{1, 2, 3, 4}), []int{1, 3, 6, 10}) {
		t.Error("Test 1 failed")
	}
	if !slices.Equal(Scan([]int{1, 2, 3, 4, 5}), []int{1, 3, 6, 10, 15}) {
		t.Error("Test 2 failed")
	}
	if !slices.Equal(Scan([]int{1, 2, 3, 4, 5, 6}), []int{1, 3, 6, 10, 15, 21}) {
		t.Error("Test 3 failed")
	}
	if !slices.Equal(Scan([]int{1, 2, 3, 4, 5, 6, 7}), []int{1, 3, 6, 10, 15, 21, 28}) {
		t.Error("Test 4 failed")
	}
}

func TestBlockedScan(t *testing.T) {
	if !slices.Equal(BlockedScan([]int{1}), []int{1}) {
		t.Error("Test 0 failed")
	}
	if !slices.Equal(BlockedScan([]int{1, 2}), []int{1, 3}) {
		t.Error("Test 0.5 failed")
	}
	if !slices.Equal(BlockedScan([]int{1, 2, 3, 4}), []int{1, 3, 6, 10}) {
		t.Error("Test 1 failed")
	}
	if !slices.Equal(BlockedScan([]int{1, 2, 3, 4, 5}), []int{1, 3, 6, 10, 15}) {
		t.Error("Test 2 failed")
	}
	if !slices.Equal(BlockedScan([]int{1, 2, 3, 4, 5, 6}), []int{1, 3, 6, 10, 15, 21}) {
		t.Error("Test 3 failed")
	}
	if !slices.Equal(BlockedScan([]int{1, 2, 3, 4, 5, 6, 7}), []int{1, 3, 6, 10, 15, 21, 28}) {
		t.Error("Test 4 failed")
	}
}

func TestFilter(t *testing.T) {
	if !slices.Equal(Filter([]int{1}, func(index int, value int) bool {
		return true
	}), []int{1}) {
		t.Error("Test 1.1 failed")
	}
	if !slices.Equal(Filter([]int{1}, func(index int, value int) bool {
		return false
	}), []int{}) {
		t.Error("Test 1.2 failed")
	}

	if !slices.Equal(Filter([]int{1, 2}, func(index int, value int) bool {
		return true
	}), []int{1, 2}) {
		t.Error("Test 2.anybody failed")
	}
	if !slices.Equal(Filter([]int{1, 2}, func(index int, value int) bool {
		return false
	}), []int{}) {
		t.Error("Test 2.nobody failed")
	}
	if !slices.Equal(Filter([]int{1, 2}, func(index int, value int) bool {
		return index%2 == 0
	}), []int{1}) {
		t.Error("Test 2.even")
	}
	if !slices.Equal(Filter([]int{1, 2}, func(index int, value int) bool {
		return index%2 == 1
	}), []int{2}) {
		t.Error("Test 2.odd")
	}

	if !slices.Equal(Filter([]int{1, 2, 3, 4, 5}, func(index int, value int) bool {
		return value%2 == 0
	}), []int{2, 4}) {
		t.Error("Test 3.even")
	}
	if !slices.Equal(Filter([]int{1, 2, 3, 4, 5}, func(index int, value int) bool {
		return value%2 == 1
	}), []int{1, 3, 5}) {
		t.Error("Test 3.odd")
	}

	if !slices.Equal(Filter([]int{1, 2, 3, 4, 5, 6, 7, 8}, func(index int, value int) bool {
		return value <= 1 || value > 4 && value%2 == 0
	}), []int{1, 6, 8}) {
		t.Error("Test 4")
	}
}
