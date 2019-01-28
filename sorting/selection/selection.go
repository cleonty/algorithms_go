package selection

func Sort(data []int) {
	for endIndex := len(data) - 1; endIndex > 0; endIndex-- {
		maxIndex := findMax(data, endIndex)
		data[endIndex], data[maxIndex] = data[maxIndex], data[endIndex]
	}
}

func findMax(data []int, endIndex int) int {
	maxIndex := 0
	max := data[maxIndex]
	for i := 1; i <= endIndex; i++ {
		if data[i] > max {
			max = data[i]
			maxIndex = i
		}
	}
	return maxIndex
}
