import "math"

/*
 * @lc app=leetcode.cn id=875 lang=golang
 *
 * [875] 爱吃香蕉的珂珂
 */

// @lc code=start
func minEatingSpeed(piles []int, H int) int {
	// big := 0
	// small := math.MaxInt32
	// length := len(piles)
	// for i := 0; i < length; i++ {
	// 	big = max(big, piles[i])
	// 	small = min(small, piles[i])
	// }
	// // 二分得到合适的值，此时，left = right
	// // 保留 right
	// left := big / H
	// right := small / H
	left := 1
	right := 1000000000
	for left < right {
		K := left + (right-left)>>1
		// fmt.Println(left, right, K)
		if count(piles, H, K) {
			right = K
		} else {
			left = K + 1
		}
	}
	return left
}
func count(piles []int, H int, K int) bool {
	now := 0
	for _, pile := range piles {
		n := 0
		if pile > K {
			n = pile / K
			if pile%K != 0 {
				n++
			}
		} else {
			n++
		}
		now += n
		// fmt.Println("now:", now)
		if now > H {
			return false
		}
	}
	return true
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
