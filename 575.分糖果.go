/*
 * @lc app=leetcode.cn id=575 lang=golang
 *
 * [575] 分糖果
 */

// @lc code=start
func distributeCandies(candies []int) int {
	counter := make(map[int]int)
	for _,c := range candies {
		n,ok:=counter[c]
		if ok {
			counter[c] = n +1
		} else {
			counter[c] = 1
		}
	}
	one := 0
	other := 0
	for _,num := range counter {
		if num == 1 {
			one++
		} else {
			other++
		}
	}
	half := len(candies)/2
	res := one + other
	if res >= half {
		return half
	}
	return res
}
// @lc code=end

