package majority

func getMajority(data []int) (int, bool) {
	size := len(data)
	if size == 0 {
		return 0, false
	}
	majIndex := 0
	count := 1
	for i := 1; i < size; i++ {
		if data[i] == data[majIndex] {
			count++
		} else {
			count--
		}
		if count == 0 {
			majIndex = i
			count = 1
		}
	}
	candidate := data[majIndex]
	count = 0
	for _, datum := range data {
		if datum == candidate {
			count++
		}
	}
	if count > size/2 {
		return candidate, true
	}
	return 0, false
}
