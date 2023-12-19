package parallel

func Map(arr []int, f func(value int) int) []int {
	r := make([]int, len(arr))
	PFor(len(arr), func(index int) {
		r[index] = f(arr[index])
	})
	return r
}

func MapIndexed(arr []int, f func(index int, value int) int) []int {
	r := make([]int, len(arr))
	PFor(len(arr), func(index int) {
		r[index] = f(index, arr[index])
	})
	return r
}
