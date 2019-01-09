package median

import "sort"

func median(data []int) float64 {
	n := len(data)
	if n == 0 {
		panic("empty list")
	}
	sort.Ints(data)
	m := float64(data[n/2])
	if n%2 == 0 {
		m = (float64(data[n/2]) + float64(data[n/2-1])) / 2
	}
	return m
}
