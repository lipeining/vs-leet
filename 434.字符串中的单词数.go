/*
 * @lc app=leetcode.cn id=434 lang=golang
 *
 * [434] 字符串中的单词数
 */

// @lc code=start
func countSegments(s string) int {
	slow, fast := 0, 0
	res := 0
	for _,c := range s {
		if string(c) == " " {
			if slow != fast {
				res++
				fast++
				slow = fast
			} else {
				slow++
				fast++
			}
		} else {
			fast++
		}
		// fmt.Println(slow, fast, res)
	}
	if fast != slow {
		res++
	}
	return res
}
// @lc code=end

