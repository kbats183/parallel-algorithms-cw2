package pfilter

import (
	"slices"
	"testing"
)

func TestFilter(t *testing.T) {
	result := make([]int, 1024)
	cnt := FilterInplace([]int{1}, func(index int, value int) bool {
		return true
	}, result)
	if !slices.Equal(result[:cnt], []int{1}) {
		t.Error("Test 1.1")
	}
	cnt = FilterInplace([]int{1}, func(index int, value int) bool {
		return false
	}, result)
	if !slices.Equal(result[:cnt], []int{}) {
		t.Error("Test 1.2")
	}

	cnt = FilterInplace([]int{1, 2}, func(index int, value int) bool {
		return true
	}, result)
	if !slices.Equal(result[:cnt], []int{1, 2}) {
		t.Error("Test 2.anybody")
	}
	cnt = FilterInplace([]int{1, 2}, func(index int, value int) bool {
		return false
	}, result)
	if !slices.Equal(result[:cnt], []int{}) {
		t.Error("Test 2.nobody")
	}
	cnt = FilterInplace([]int{1, 2}, func(index int, value int) bool {
		return index%2 == 0
	}, result)
	if !slices.Equal(result[:cnt], []int{1}) {
		t.Error("Test 2.even")
	}
	cnt = FilterInplace([]int{1, 2}, func(index int, value int) bool {
		return index%2 == 1
	}, result)
	if !slices.Equal(result[:cnt], []int{2}) {
		t.Error("Test 2.odd")
	}

	cnt = FilterInplace([]int{1, 2, 3, 4, 5}, func(index int, value int) bool {
		return value%2 == 0
	}, result)
	if !slices.Equal(result[:cnt], []int{2, 4}) {
		t.Error("Test 3.even")
	}
	cnt = FilterInplace([]int{1, 2, 3, 4, 5}, func(index int, value int) bool {
		return value%2 == 1
	}, result)
	if !slices.Equal(result[:cnt], []int{1, 3, 5}) {
		t.Error("Test 3.odd")
	}

	cnt = FilterInplace([]int{1, 2, 3, 4, 5, 6, 7, 8}, func(index int, value int) bool {
		return value <= 1 || value > 4 && value%2 == 0
	}, result)
	if !slices.Equal(result[:cnt], []int{1, 6, 8}) {
		t.Error("Test 4")
	}
}

func TestFilterHuge(t *testing.T) {
	arr := make([]int, 1024)
	res := make([]int, 1024)
	ans := make([]int, 0)
	for i := 0; i < 1024; i++ {
		arr[i] = i
		if i%3 == 2 || i%7 == 5 {
			ans = append(ans, i)
		}
	}
	cnt := FilterInplace(arr, func(index int, value int) bool {
		return value%3 == 2 || value%7 == 5
	}, res)
	if !slices.Equal(res[:cnt], ans) {
		t.Errorf("Test failed %v != %v", ans, res)
	}
}
