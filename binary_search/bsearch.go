package bsearch

func BinarySearch(data []int, value int) bool {
	low := 0
	size := len(data)
	hi := size - 1
	mid := 0
	for low <= hi {
		mid = (low + hi) / 2
		if data[mid] == value {
			return true
		} else if data[mid] < value {
			low = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return false
}

func BinarySearchRecursive(data []int, value int) bool {
	low := 0
	size := len(data)
	hi := size - 1
	mid := 0
	mid = (low + hi) / 2
	if size == 0 {
		return false
	}
	if data[mid] == value {
		return true
	} else if data[mid] < value {
		low = mid + 1
		return BinarySearchRecursive(data[low:hi+1], value)
	} else {
		hi = mid - 1
		return BinarySearchRecursive(data[low:hi+1], value)
	}
}
