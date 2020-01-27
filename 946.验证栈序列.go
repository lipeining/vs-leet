/*
 * @lc app=leetcode.cn id=946 lang=golang
 *
 * [946] 验证栈序列
 */

// @lc code=start
func validateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, 0)
	popMap := getMap(popped)
	for i:=1;i<len(pushed);i++ {
		
		prev,_ := popMap[pushed[i-1]]
		next,_ := popMap[pushed[i]]
		if next < prev {
			return false
		}
	}
	return true
}
func getMap(list []int) map[int]int {
	smap := make(map[int]int)
	for i:=0;i<len(list);i++ {
		smap[list[i]] = i
	}
	return smap
}
// @lc code=end

