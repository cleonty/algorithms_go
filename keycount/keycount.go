package keycount

func keyCountSlow(data []int, key int) int {
	count := 0
	for _, datum := range data {
		if datum == key {
			count++
		}
	}
	return count
}
func keyCountFast(data []int, key int) int {
	firstIndex := findFirstIndex(data, key, 0, len(data)-1)
	lastIndex := findLastIndex(data, key, 0, len(data)-1)
	if firstIndex == -1 || lastIndex == -1 {
		return 0
	}
	return (lastIndex - firstIndex + 1)
}

func findFirstIndex(data []int, key int, start, end int) int {
	if end < start {
		return -1
	}
	mid := (start + end) / 2
	if key == data[mid] && (mid == start || data[mid-1] != key) {
		return mid
	}
	if key <= data[mid] {
		return findFirstIndex(data, key, start, mid-1)
	}
	return findFirstIndex(data, key, mid+1, end)
}

func findLastIndex(data []int, key int, start, end int) int {
	if end < start {
		return -1
	}
	mid := (start + end) / 2
	midElement := data[mid]
	endElement := data[end]
	if key == midElement && (endElement == key || key != data[mid+1]) {
		return mid
	}
	if key >= midElement {
		return findLastIndex(data, key, mid+1, end)
	}
	return findLastIndex(data, key, start, mid-1)
}
