package parallel

func Pow2Size(n int) int {
	if n == 0 {
		return 0
	}
	a := 1
	for a < n {
		a <<= 1
	}
	return a
}

func Scan(array []int) []int {
	if len(array) == 0 {
		return array
	}
	s := Pow2Size(len(array))
	b := make([]int, (s<<1)-1)
	c := make([]int, len(array))
	blockSize := (len(array) + PForDivider - 1) / PForDivider
	if blockSize < 1 {
		blockSize = 1
	}
	ReduceImpl(array, 0, len(array), 0, b, blockSize)
	scanPropagate(b, 0, len(array), 0, 0, c, blockSize)
	return c
}

func scanPropagate(
	array []int,
	l int,
	r int,
	x int,
	fromLeft int,
	result []int,
	blockSize int) {
	if l+1 == r {
		result[l] = array[x] + fromLeft
		return
	}
	m := (l + r) / 2
	if r-l <= blockSize {
		scanPropagate(array, l, m, x*2+1, fromLeft, result, blockSize)
		scanPropagate(array, m, r, x*2+2, fromLeft+array[x*2+1], result, blockSize)
	} else {
		Fork2Join(
			func() {
				scanPropagate(array, l, m, x*2+1, fromLeft, result, blockSize)
			}, func() {
				scanPropagate(array, m, r, x*2+2, fromLeft+array[x*2+1], result, blockSize)
			})
	}
}

func BlockedScan(array []int) []int {
	if len(array) == 0 {
		return array
	}
	s := Pow2Size(PForDivider)
	b := make([]int, (s<<1)-1)
	c := make([]int, len(array))
	blockSize := (len(array) + PForDivider - 1) / PForDivider
	if blockSize < 1 {
		blockSize = 1
	}
	blockedReduceImpl(array, 0, len(array), 0, b, blockSize)
	blockedScanPropagate(array, b, 0, len(array), 0, 0, c, blockSize)
	return c
}

func blockedScanPropagate(
	array []int,
	reduced []int,
	l int,
	r int,
	x int,
	fromLeft int,
	result []int,
	blockSize int) {
	if r-l <= blockSize {
		result[l] = fromLeft + array[l]
		for i := l + 1; i < r; i++ {
			result[i] = result[i-1] + array[i]
		}
	} else {
		m := (l + r) / 2
		Fork2Join(
			func() {
				blockedScanPropagate(array, reduced, l, m, x*2+1, fromLeft, result, blockSize)
			}, func() {
				blockedScanPropagate(array, reduced, m, r, x*2+2, fromLeft+reduced[x*2+1], result, blockSize)
			})
	}
}
