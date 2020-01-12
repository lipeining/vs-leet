/*
 * @lc app=leetcode.cn id=532 lang=golang
 *
 * [532] 数组中的K-diff数对
 */
import "sort"
// @lc code=start
func findPairs(nums []int, k int) int {
	if k < 0 {
		return 0
	}
	nMap := make(map[int]bool) 
	dMap := make(map[int]bool) 
	for _,n := range nums {
		if nMap[n-k] {
			dMap[n-k] = true
		}
		if nMap[n+k] {
			dMap[n] = true
		}
		nMap[n] = true
	}
	return len(dMap)
// if k<0:
// return 0
// saw, diff = set(), set()
// for i in nums:
// 	if i-k in saw:
// 		diff.add(i-k)
// 	if i+k in saw:
// 		diff.add(i)
// 	saw.add(i)
// return len(diff)

// 作者：yybeta
// 链接：https://leetcode-cn.com/problems/k-diff-pairs-in-an-array/solution/ha-xi-onzui-jian-dan-jie-fa-by-mai-mai-mai-mai-zi/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
}
// @lc code=end

