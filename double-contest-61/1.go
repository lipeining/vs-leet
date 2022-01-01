package main

func main() {
	findOriginalArray([]int{1, 2, 3, 4, 6, 8})
	findOriginalArray([]int{6, 3, 0, 1})
	findOriginalArray([]int{1})
	findOriginalArray([]int{0, 0, 0, 0})
	findOriginalArray([]int{0, 3, 2, 4, 6, 0})
}
func findOriginalArray(changed []int) []int {
	m := len(changed)
	if m%2 == 1 {
		return []int{}
	}

	return ans 
}
func countKDifference(nums []int, k int) int {
	n := len(nums)
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	ans := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if abs(nums[i]-nums[j]) == k {
				ans++
			}
		}
	}
	return ans
}
