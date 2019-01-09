package maxprofit

func maxProfit(data []int) int {
	if len(data) == 0 {
		return 0
	}
	min := data[0]
	maxProfit := 0
	for _, datum := range data {
		currentProfit := datum - min
		if currentProfit > maxProfit {
			maxProfit = currentProfit
		}
		if datum < min {
			min = datum
		}
	}
	return maxProfit
}
