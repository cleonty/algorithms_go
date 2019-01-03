package main

import (
	"fmt"
	"sort"
)

func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}

func findSumClosestToZero(data []int) int {
	sort.Ints(data)
	firstIndex := 0
	lastIndex := len(data) - 1
	minSum := abs(data[firstIndex] + data[lastIndex])
	for firstIndex < lastIndex {
		sum := data[firstIndex] + data[lastIndex]
		fmt.Printf("first [%d] = %d last [%d] = %d sum (%d, %d)\n",
			firstIndex, data[firstIndex], lastIndex, data[lastIndex], sum, abs(sum))
		if abs(sum) < minSum {
			minSum = abs(sum)
		}
		if sum > 0 {
			lastIndex--
		} else if sum < 0 {
			firstIndex++
		} else {
			break
		}
	}
	return minSum
}
