/*
 * @lc app=leetcode.cn id=507 lang=golang
 *
 * [507] 完美数
 */

// @lc code=start
func checkPerfectNumber(num int) bool {
    if num <= 0 {
		return false
	}
	sum := 0
	for i:=1;i*i<num;i++ {
		if num %i == 0 {
			sum += i
			if i*i!=num {
				sum+=num/i
			}
		}
	}
	return sum == num * 2
}
// @lc code=end

