package merge

func Sort(data []int) {
	size := len(data)
	if size > 1 {
		middle := size / 2
		firstHalf := data[0:middle]
		secondHalf := data[middle:]
		Sort(firstHalf)
		Sort(secondHalf)
		result := merge(firstHalf, secondHalf)
		copy(data, result)
	}
}

func merge(first, second []int) []int {
	var result []int
	firstIndex := 0
	secondIndex := 0
	firstSize := len(first)
	secondSize := len(second)
	for {
		if firstIndex < firstSize && secondIndex < secondSize {
			if first[firstIndex] < second[secondIndex] {
				result = append(result, first[firstIndex])
				firstIndex++
			} else {
				result = append(result, second[secondIndex])
				secondIndex++
			}
		} else if firstIndex < firstSize {
			result = append(result, first[firstIndex])
			firstIndex++
		} else if secondIndex < secondSize {
			result = append(result, second[secondIndex])
			secondIndex++
		} else {
			break
		}
	}
	return result
}
