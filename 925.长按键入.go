/*
 * @lc app=leetcode.cn id=925 lang=golang
 *
 * [925] 长按键入
 */

// @lc code=start
func isLongPressedName(name string, typed string) bool {
	slow := 0
	for i:=0;i<len(name);i++ {
		if slow >= len(typed) {
			return false
		}
		if name[i] != typed[slow] {
			if slow == 0 || typed[slow] != typed[slow-1] {
				return false
			}
			for slow < len(typed) && typed[slow] == typed[slow-1] {
				slow++
			}
			if slow == len(typed) || typed[slow]!=name[i] {
				return false
			}	
		}
		slow++
	}
	return true
}
// @lc code=end

