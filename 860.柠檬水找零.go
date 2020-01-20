/*
 * @lc app=leetcode.cn id=860 lang=golang
 *
 * [860] 柠檬水找零
 */

// @lc code=start
func lemonadeChange(bills []int) bool {
	count5, count10 := 0,0
	for i:=0;i<len(bills);i++{
		if bills[i] == 5 {
			count5++
		} else if bills[i] == 10 {
			if count5 == 0 {
				return false
			}
			count5--
			count10++
		} else {
			if count10 == 0 {
				if count5 < 3 {
					return false
				}
				count5-=3
			} else {
				count10--
				if count5 == 0 {
					return false
				}
				count5--
			}
		}
	}
	return true
}
// @lc code=end

