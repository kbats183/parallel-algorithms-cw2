package parallel

import "sync"

func Fork2Join(body1 func(), body2 func()) {
	//var wg sync.WaitGroup
	//wg.Add(2)
	//go func() {
	//	body1()
	//	wg.Done()
	//}()
	//go func() {
	//	body2()
	//	wg.Done()
	//}()
	//wg.Wait()
	body1()
	body2()
}

var PForDivider = 4

func PFor(n int, body func(index int)) {
	blockSize := (n + PForDivider - 1) / PForDivider
	if blockSize < 1 {
		blockSize = 1
	}
	var wg sync.WaitGroup
	for i := 0; i < n; i += blockSize {
		wg.Add(1)
		iFinal := i
		go func() {
			for j := iFinal; j < n && j < iFinal+blockSize; j++ {
				body(j)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
