/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		want := target - num
		if j, ok := m[want]; ok {
			return []int{j, i}
		}
		m[num] = i
	}
	return nil
}

// @lc code=end

