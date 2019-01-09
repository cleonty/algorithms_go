package separate

func separate(data []int) {
	evenIndex := 0
	oddIndex := len(data) - 1
	for evenIndex < oddIndex {
		if isEven(data[evenIndex]) {
			evenIndex++
		} else if isOdd(data[oddIndex]) {
			oddIndex--
		} else {
			swap(data, evenIndex, oddIndex)
		}
	}
}

func isOdd(n int) bool {
	return n%2 == 1
}

func isEven(n int) bool {
	return n%2 == 0
}

func swap(data []int, first, second int) {
	data[first], data[second] = data[second], data[first]
}
