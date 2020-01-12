/*
 * @lc app=leetcode.cn id=414 lang=golang
 *
 * [414] 第三大的数
 */

// @lc code=start
func thirdMax(nums []int) int {
	minInt := -(1<<31 - 1)-1-1
	first,second,third := minInt,minInt,minInt
	for _,n := range nums {
		if n > first {
			first, second, third = n, first, second
		} else if n== first {
			continue
		} else if n >second {
			second,third = n, second
		} else if n == second {
			continue
		} else if n > third {
			third = n
		}
	}
	if third != minInt {
		return third
	}
	return first
}
// @lc code=end

