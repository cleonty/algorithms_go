package countingsort

func SortAges(ages []int) []int {
	const max = 130
	ageArr := make([]int, max+1)
	for _, age := range ages {
		if age < max {
			ageArr[age]++
		}
	}
	index := 0
	for age, count := range ageArr {
		for i := 0; i < count; i++ {
			ages[index] = age
			index++
		}
	}
	return ages[:]
}
