import "fmt"

/*
 * @lc app=leetcode.cn id=1011 lang=golang
 *
 * [1011] 在 D 天内送达包裹的能力
 */

// @lc code=start
func shipWithinDays(weights []int, D int) int {
	left := 1
	right := 100000000
	for left < right {
		K := left + (right-left)>>1
		fmt.Println(left, right, K)
		if count(weights, D, K) {
			right = K
		} else {
			left = K + 1
		}
	}
	return left
}
func count(weights []int, D int, K int) bool {
	cur := K
	for _, w := range weights {
		if w > K {
			return false
		}
		if cur < w {
			cur = K
			D--
		}
		cur -= w
	}
	return D > 0
}

// @lc code=end
