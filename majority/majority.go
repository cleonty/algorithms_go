package majority

import (
	"sort"
)

func getMajority(data []int) (int, bool) {
	if len(data) == 0 {
		return 0, false
	}
	sort.Ints(data)
	candidateIndex := len(data) / 2
	candidate := data[candidateIndex]
	count := 0
	for _, datum := range data {
		if datum == candidate {
			count++
		}
	}
	if count > len(data)/2 {
		return candidate, true
	}
	return 0, false
}
