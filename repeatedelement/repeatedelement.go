package repeatedelement

func findRepeatedElement(data []int) int {
	m := make(map[int]int)
	for _, datum := range data {
		if _, ok := m[datum]; ok {
			return datum
		}
		m[datum] = 1
	}
	return 0
}
