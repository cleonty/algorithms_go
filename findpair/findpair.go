package findpair

import "sort"

func findPair(data []int, value int) bool {
	sort.Ints(data)
	first := 0
	last := len(data) - 1
	if len(data) == 0 {
		return false
	}
	if value > 2*data[last] {
		return false
	}
	if value <= data[first] {
		return false
	}
	for first < last {
		sum := first + last
		if sum == value {
			return true
		} else if sum < value {
			first++
		} else {
			last--
		}
	}
	return false
}
