import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode.cn id=1300 lang=golang
 *
 * [1300] 转变数组后最接近目标值的数组和
 */

// @lc code=start
func findBestValue(arr []int, target int) int {
	left := 0
	right := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > right {
			right = arr[i]
		}
	}
	for left < right {
		mid := left + (right-left)>>1
		fmt.Println(left, right, mid)
		if helper(arr, target, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	lower := left - 1
	upper := left
	lowerSum := count(arr, lower)
	upperSum := count(arr, upper)
	if math.Abs(float64(lowerSum-target)) <= math.Abs(float64(upperSum-target)) {
		return lower
	}
	return upper
}
func count(arr []int, K int) int {
	sum := 0
	for _, n := range arr {
		if n > K {
			sum += K
		} else {
			sum += n
		}
	}
	return sum
}
func helper(arr []int, target int, K int) bool {
	sum := count(arr, K)
	return sum > target
}

// @lc code=end
