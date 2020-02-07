import (
	"strconv"
	"unicode"
)

/*
 * @lc app=leetcode.cn id=241 lang=golang
 *
 * [241] 为运算表达式设计优先级
 */

// @lc code=start
func diffWaysToCompute(input string) []int {
	// merge
	length := len(input)
	// 网友的思路：
	// 分解：按运算符分成左右两部分，分别求解
	// 解决：实现一个递归函数，输入算式，返回算式解
	// 合并：根据运算符合并左右两部分的解，得出最终解

	ans := make([]int, 0)
	// 分隔符比较难找，可能有很大的数字
	i := 0
	alone := true
	for i < length {
		if unicode.IsDigit(rune(input[i])) {
			i++
			continue
		}
		alone = false
		left := diffWaysToCompute(input[:i])
		right := diffWaysToCompute(input[i+1:])
		for _, l := range left {
			for _, r := range right {
				if input[i] == '+' {
					ans = append(ans, l+r)
				} else if input[i] == '-' {
					ans = append(ans, l-r)
				} else {
					ans = append(ans, l*r)
				}
			}
		}
		i++
	}
	if alone {
		num, _ := strconv.Atoi(input)
		ans = append(ans, num)
		return ans
	}
	return ans
}

// @lc code=end
