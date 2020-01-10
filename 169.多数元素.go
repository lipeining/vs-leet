/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */
import "fmt"
// @lc code=start
func majorityElement(nums []int) int {
	counter := map[int]int{}
	for _,num := range nums {
		c, ok := counter[num]
		if ok {
			counter[num] = c + 1
		} else {
			counter[num] = 1
		}
	}
	max := 0
	res := 0
	for k,v := range counter {
		// fmt.Println(k, v)
		if v > max {
			max = v
			res = k
		}
	}
	return res
}
// @lc code=end

