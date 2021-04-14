func maxValue(n int, index int, maxSum int) int {
	lo := 1
	hi := maxSum + 1
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	calcSum := func(x, n, index int) int {
		sum := n - 1 + x
		left := min(index, x-1)
		sum += ((x - left) + (x - 1)) * left / 2
		right := min(n-index-1, x-1)
		sum += ((x - 1) + (x - right)) * right / 2
		sum = sum - left - right
		return sum
	}
	for lo < hi {
		mid := lo + (hi-lo)/2
		if calcSum(mid, n, index) <= maxSum {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo - 1
}