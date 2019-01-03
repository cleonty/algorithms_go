package bitoniclist

// A bitonic list comprises of an increasing sequence of integers
// immediately followed by a decreasing sequence of integers.

// findMaxima finds a maxima in a bitonic list. It returns (maximaIndex, true) if found.
// and (-1, false) otherwise
func findMaxima(data []int) (int, bool) {
	if len(data) < 3 {
		return -1, false
	}
	firstIndex := 0
	lastIndex := len(data) - 1
	for firstIndex < lastIndex {
		middleIndex := (firstIndex + lastIndex) / 2
		middle := data[middleIndex]
		prev := data[middleIndex-1]
		next := data[middleIndex+1]
		if middle > prev && next < middle {
			return middleIndex, true
		} else if prev < middle && middle < next {
			firstIndex++
		} else if prev > middle && middle > next {
			lastIndex--
		} else {
			break
		}
	}
	return -1, false
}

// findElement searches for val in bitonic array data.
// returns (index, true) if found, (-1, false) otherwise
func findElement(data []int, val int) (int, bool) {
	maximaIndex, found := findMaxima(data)
	if !found {
		return -1, false
	}
	if val == data[maximaIndex] {
		return maximaIndex, true
	} else if index, found := binarySearch(data[:maximaIndex], val, true); found {
		return index, true
	} else {
		index, found := binarySearch(data[maximaIndex+1:], val, false)
		if found {
			return index + maximaIndex + 1, true
		}
		return -1, false
	}
}

func binarySearch(data []int, val int, inc bool) (int, bool) {
	if len(data) == 0 {
		return -1, false
	}
	lowIndex := 0
	hiIndex := len(data) - 1
	for lowIndex <= hiIndex {
		midIndex := (lowIndex + hiIndex) / 2
		mid := data[midIndex]
		if mid == val {
			return midIndex, true
		} else if val < mid {
			if inc {
				hiIndex = midIndex - 1
			} else {
				lowIndex = midIndex + 1
			}
		} else {
			if inc {
				lowIndex = midIndex + 1
			} else {
				hiIndex = midIndex - 1
			}
		}
	}
	return -1, false
}
