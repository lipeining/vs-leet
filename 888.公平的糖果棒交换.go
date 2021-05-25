/*
 * @lc app=leetcode.cn id=888 lang=golang
 *
 * [888] 公平的糖果棒交换
 *
 * 
 */

// @lc code=start
func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	am, bm := make(map[int]bool), make(map[int]bool)
	sum1 := 0
	sum2 := 0
	for _, num := range aliceSizes {
		sum1 += num
		am[num] = true
	}
	for _, num := range bobSizes {
		sum2 += num
		bm[num] = true
	}
	// sum1-a+b = sum2-b+a
	// sum1-sum2 = 2 * (a-b)
	diff := sum1-sum2
	diff /= 2
	for a := range am {
		b := a - diff
		if _, ok := bm[b]; ok {
			return []int{a, b}
		}
	}
	return []int{-1, -1}
}
// @lc code=end

