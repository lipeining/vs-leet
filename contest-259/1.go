package main

func sumOfBeauties(nums []int) int {
	n := len(nums)
	prefix := make([]int, n)
	suffix := make([]int, n)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	prefix[0] = nums[0]
	for i := 1; i < n; i++ {
		prefix[i] = max(prefix[i-1], nums[i-1])
	}
	suffix[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		suffix[i] = min(suffix[i+1], nums[i+1])
	}
	ans := 0
	for i := 1; i <= n-2; i++ {
		if prefix[i] < nums[i] && nums[i] < suffix[i] {
			ans += 2
		} else if nums[i-1] < nums[i] && nums[i] < nums[i+1] {
			ans++
		}
	}
	return ans
}
func finalValueAfterOperations(operations []string) int {
	ans := 0
	for _, op := range operations {
		if op[0] == '+' || op[2] == '+' {
			ans++
		} else {
			ans--
		}
	}
	return ans
}
