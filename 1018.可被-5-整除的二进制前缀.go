/*
 * @lc app=leetcode.cn id=1018 lang=golang
 *
 * [1018] 可被 5 整除的二进制前缀
 */

// @lc code=start
func prefixesDivBy5(nums []int) []bool {
	ans := make([]bool, len(nums))
	x := 0
	for i, v := range nums {
		x = (x<<1 | v) % 5
		ans[i] = x == 0
	}
	return ans

	// 作者：LeetCode-Solution
	// 链接：https://leetcode-cn.com/problems/binary-prefix-divisible-by-5/solution/ke-bei-5-zheng-chu-de-er-jin-zhi-qian-zh-asih/
	// 来源：力扣（LeetCode）
	// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
}

// @lc code=end

