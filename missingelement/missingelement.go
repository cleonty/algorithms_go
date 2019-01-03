package missingelement

import "math"

func findMissing(data []int) int {
	size := len(data)
	tmp := make([]int, size)
	for _, datum := range data {
		if datum > 0 && datum <= size {
			tmp[datum-1] = 1
		}
	}
	for i, datum := range tmp {
		if datum == 0 {
			return i + 1
		}
	}
	return size + 1
}

func findMissing2(data []int) int {
	size := len(data)
	sum := 0
	for _, datum := range data {
		sum += datum
	}
	expectedSum := (size + 1) * size / 2
	missing := expectedSum - sum
	if missing == 0 {
		return size + 1
	}
	return int(math.Abs(float64(missing)))
}
