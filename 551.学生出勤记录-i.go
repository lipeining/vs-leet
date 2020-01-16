/*
 * @lc app=leetcode.cn id=551 lang=golang
 *
 * [551] 学生出勤记录 I
 */

// @lc code=start
func checkRecord(s string) bool {
	aNum := 0
	lNum := 0
	for i:=0;i<len(s);i++ {
		if "A" == string(s[i]) {
			aNum++
			lNum=0
			if aNum == 2 {
				return false
			}
		} else if "L" == string(s[i]) {
			lNum++
			if lNum == 3 {
				return false
			}
		} else {
			lNum = 0
		}
	}
	return true
}
// @lc code=end

