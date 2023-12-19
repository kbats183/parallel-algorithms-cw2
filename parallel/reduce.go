package parallel

func ReduceImpl(array []int, l int, r int, x int, reduced []int, blockSize int) {
	if l+1 == r {
		reduced[x] = array[l]
		return
	}
	m := (l + r) / 2
	if r-l <= blockSize {
		ReduceImpl(array, l, m, x*2+1, reduced, blockSize)
		ReduceImpl(array, m, r, x*2+2, reduced, blockSize)
	} else {
		Fork2Join(func() {
			ReduceImpl(array, l, m, x*2+1, reduced, blockSize)
		}, func() {
			ReduceImpl(array, m, r, x*2+2, reduced, blockSize)
		})
	}
	reduced[x] = reduced[x*2+1] + reduced[x*2+2]
}

func blockedReduceImpl(array []int, l int, r int, x int, reduced []int, blockSize int) {
	if l+1 == r {
		reduced[x] = array[l]
		return
	}
	m := (l + r) / 2
	if r-l <= blockSize {
		for i := l; i < r; i++ {
			reduced[x] += array[i]
		}
	} else {
		Fork2Join(func() {
			blockedReduceImpl(array, l, m, x*2+1, reduced, blockSize)
		}, func() {
			blockedReduceImpl(array, m, r, x*2+2, reduced, blockSize)
		})
		reduced[x] = reduced[x*2+1] + reduced[x*2+2]
	}
}
