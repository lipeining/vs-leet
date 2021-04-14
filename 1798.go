func getMaximumConsecutive(coins []int) int {
	sort.Ints(coins)
	x := 0
	for _, y := range coins {
		if y > x+1 {
			break
		}
		x += y
	}
	return x + 1
}