package parallel

func Filter(array []int, predicate func(index int, value int) bool) []int {
	if len(array) == 0 {
		return []int{}
	}
	flags := MapIndexed(array, func(index, value int) int {
		if predicate(index, value) {
			return 1
		} else {
			return 0
		}
	})
	positions := BlockedScan(flags)
	result := make([]int, positions[len(positions)-1])
	PFor(len(array), func(index int) {
		if flags[index] == 1 {
			result[positions[index]-1] = array[index]
		}
	})
	return result
}
