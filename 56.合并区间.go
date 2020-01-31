/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func (i int, j int)bool {
		left := intervals[i][0] - intervals[j][0]
		if left != 0 {
			return left < 0
		}
		return intervals[i][1] < intervals[j][1]
	})
	fmt.Println(intervals)
	stack := make([][]int, 0)
	for i:=0;i<len(intervals);i++ {
		length := len(stack)
		if length == 0 {
			stack = append(stack, intervals[i])
		} else {
			top := stack[length-1]
			m := getMerge(top, intervals[i])
			fmt.Println(m)
			if len(m) == 0 {
				stack = append(stack, intervals[i])
			} else {
				stack = stack[:length-1]
				stack = append(stack, m)
			}
		}
	}
	return stack
}
func getMerge(top []int, interval []int) []int {
	ans := make([]int, 0)
	if top[1] < interval[0] {
		return ans
	}
	max := interval[1]
	if max < top[1] {
		max = top[1]
	}
	ans = append(ans, top[0], max)
	return ans
}
// @lc code=end

