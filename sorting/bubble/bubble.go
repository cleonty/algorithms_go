package bubble

import "fmt"

func Sort(data []int) []int {
	size := len(data)
	totalSwapCount := 0
	for i := size - 1; i > 0; i-- {
		swapCount := bubbleUp(data, i)
		totalSwapCount += swapCount
		if swapCount == 0 {
			break
		}
	}
	fmt.Printf("total swap count is %d\n", totalSwapCount)
	return data
}

func bubbleUp(data []int, stop int) int {
	swapCount := 0
	for i := 0; i < stop; i++ {
		if data[i] > data[i+1] {
			data[i], data[i+1] = data[i+1], data[i]
			swapCount++
		}
	}
	return swapCount
}
