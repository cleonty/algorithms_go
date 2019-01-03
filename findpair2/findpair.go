package findpair

func findPair(data []int, value int) bool {
	table := make(map[int]bool)
	for _, datum := range data {
		if _, ok := table[value-datum]; ok {
			return true
		}
		table[datum] = true
	}
	return false
}
