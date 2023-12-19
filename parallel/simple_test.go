package parallel

import (
	"testing"
)

func TestFork2Join(t *testing.T) {
	r := make([]int, 2)
	Fork2Join(func() {
		r[0] = 1
	}, func() {
		r[1] = 2
	})

	if r[0] != 1 {
		t.Error("failed: r[0] != 1")
	}
	if r[1] != 2 {
		t.Error("failed r[1] != 2")
	}
}

func TestPFor(t *testing.T) {
	r := make([]int, 3)
	PFor(3, func(i int) {
		r[i] = i + 1
	})

	if r[0] != 1 {
		t.Error("failed: r[0] != 1")
	}
	if r[1] != 2 {
		t.Error("failed r[1] != 2")
	}
	if r[2] != 3 {
		t.Error("failed r[2] != 3")
	}
}
