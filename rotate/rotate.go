package main

func rotate(arr []int, k int) []int {
	reverse(arr, 0, k)
	reverse(arr, k, len(arr))
	reverse(arr, 0, len(arr))
	return arr
}

func reverse(arr []int, start int, end int) {
	i := start
	j := end - 1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}
