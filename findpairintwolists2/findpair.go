package findpair

import "sort"

func findPair(a, b []int, value int) bool {
	sort.Ints(a)
	sort.Ints(b)
	minIndex := 0
	maxIndex := len(b) - 1
	for minIndex < len(a) && maxIndex >= 0 {
		sum := a[minIndex] + b[maxIndex]
		if sum == value {
			return true
		} else if sum < value {
			minIndex++
		} else {
			maxIndex--
		}
	}
	return false
}
