package quick

import "fmt"

func Select(data []int, index int) int {
	size := len(data)
	qsortUtil(data, 0, size-1, index int)
	return data[index]
}

func partition(data []int, left, right int) (pivotIndex int) {
	pivotValue := data[right]
	wallIndex := left - 1
	for currentIndex := left; currentIndex < right; currentIndex++ {
		if data[currentIndex] <= pivotValue {
			wallIndex = moveWall(wallIndex)
			swap(data, wallIndex, currentIndex)
		}
	}
	wallIndex = moveWall(wallIndex)
	swap(data, wallIndex, right)
	return wallIndex
}

func swap(data []int, i, j int) {
	data[i], data[j] = data[j], data[i]
}

func moveWall(wallIndex int) int {
	return wallIndex + 1
}

func moveInFrontOfWall(data []int, wallIndex int, elementIndex int) {
	swap(data, wallIndex, elementIndex)
}

func qsortUtil(data []int, left, right int, index int) {
	if left < right {
		q := partition(data, left, right)
		qsortUtil(data, left, q-1)
		qsortUtil(data, q+1, right)
	}
}

func printSlice(data []int, a, b int) {
	result := ""
	for i, value := range data {
		if i == a || i == b {
			result += " [" + fmt.Sprintf("%d", value) + "]"
		} else {
			result += fmt.Sprintf("  %d  ", value)
		}
	}
	fmt.Println(result)
}
