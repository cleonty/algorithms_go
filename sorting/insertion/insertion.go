package insertion

func Sort(data []int) {
	for i := range data {
		insert(data, i)
	}
}

func insert(data []int, index int) {
	current := data[index]
	var i int
	shifted := false
	insertionIndex := -1
	for i = index - 1; i >= 0; i-- {
		if data[i] > data[i+1] {
			data[i], data[i+1] = data[i+1], data[i]
			shifted = true
			insertionIndex = i
		} else {
			break
		}
	}
	if shifted {
		data[insertionIndex] = current
	}
}
