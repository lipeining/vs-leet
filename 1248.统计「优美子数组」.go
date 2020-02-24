/*
 * @lc app=leetcode.cn id=1248 lang=golang
 *
 * [1248] 统计「优美子数组」
 */

// @lc code=start
func numberOfSubarrays(nums []int, k int) int {
	length := len(nums)
	p := make([]int, length+1)
	for i := 1; i <= length; i++ {
		p[i] = p[i-1]
		if nums[i-1]%2 == 1 {
			p[i]++
		}
	}
	// p[j] - p[i] = k  的问题
	m := make(map[int]int)
	ans := 0
	for i := 0; i <= length; i++ {
		n, _ := m[p[i]-k]
		ans += n
		m[p[i]]++
	}
	return ans
}

// @lc code=end
