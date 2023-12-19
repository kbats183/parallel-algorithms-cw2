package parallel

import (
	"testing"
)

func TestPow2Size(t *testing.T) {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := []int{0, 1, 2, 4, 4, 8, 8, 8, 8, 16}
	for i := 0; i < len(a); i++ {
		if pow2size(a[i]) != b[i] {
			t.Errorf("pow2size(%d) != %d", a[i], b[i])
		}
	}
}
