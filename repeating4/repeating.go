package main

import "fmt"

func printRepeating4(data []int, intrange int) {
	size := len(data)
	count := make([]int, intrange)
	fmt.Println(count)
	fmt.Println("Repeating elements are : ")
	for i := 0; i < size; i++ {
		if count[data[i]] == 1 {
			fmt.Println(" ", data[i])
		} else {
			count[data[i]]++
		}
	}
	fmt.Println()
}

func main() {
	printRepeating4([]int{1, 2, 3, 4, 5, 6, 2, 4, 6}, 10)
}
