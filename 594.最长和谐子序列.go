/*
 * @lc app=leetcode.cn id=594 lang=golang
 *
 * [594] 最长和谐子序列
 */

// @lc code=start
func findLHS(nums []int) int {
	// 可以考虑原地交换的方法，那样的话，只需要快慢指针即可
	counter := make(map[int]int)
	for _,n :=range nums {
		c,ok:=counter[n]
		if ok {
			counter[n] = c + 1
		} else {
			counter[n] = 1
		}
	}
	max := 0
	for n,c := range counter {
		up := n +1
		down := n - 1
		if upNum,ok:=counter[up];ok{
			if (c+upNum) > max {
				max = upNum + c
			}
		}
		if downNum,ok:=counter[down];ok{
			if (c+downNum) > max {
				max = downNum + c
			}
		}
	}
	return max
}
// @lc code=end

