package bucket

func Sort(data []int) {
	if len(data) < 2 {
		return
	}
	min, max := findMinAndMax(data)
	bucket := fillBucket(data, min, max)
	restoreFromBucket(bucket, data, min)
}

func findMinAndMax(data []int) (int, int) {
	min := data[0]
	max := data[0]
	for _, val := range data {
		if val < min {
			min = val
		} else if val > max {
			max = val
		}
	}
	return min, max
}

func fillBucket(data []int, min, max int) []int {
	bucket := make([]int, max-min+1)
	for _, val := range data {
		bucket[val-min]++
	}
	return bucket
}

func restoreFromBucket(bucket []int, data []int, min int) {
	index := 0
	for valueMinusMin, count := range bucket {
		for i := 0; i < count; i++ {
			value := valueMinusMin + min
			data[index] = value
			index++
		}
	}
}
