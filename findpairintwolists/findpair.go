package findpair

func findPair(a, b []int, value int) bool {
	tableA := make(map[int]bool)
	for _, val := range a {
		tableA[val] = true
	}
	for _, val := range b {
		if _, ok := tableA[value-val]; ok {
			return true
		}
	}
	return false
}
