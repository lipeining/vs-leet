/*
 * @lc app=leetcode.cn id=682 lang=golang
 *
 * [682] 棒球比赛
 */

// @lc code=start
func calPoints(ops []string) int {
	stack := make([]int, 0)
	for _,op := range ops {
		length := len(stack)
		if op == "+" {
			num := stack[length-1] + stack[length-2]
			stack = append(stack, num)
		} else if op == "D" {
			stack = append(stack, stack[length-1] * 2)
		} else if op == "C" {
			stack = stack[:length-1]
		} else {
			num, _ := strconv.Atoi(op)
			stack = append(stack, num)
		}
		fmt.Println(stack)
	}
	res := 0
	for _,num := range stack {
		res += num
	}
	return res
}
// @lc code=end

