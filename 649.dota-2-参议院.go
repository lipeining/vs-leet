/*
 * @lc app=leetcode.cn id=649 lang=golang
 *
 * [649] Dota2 参议院
 */

// @lc code=start
func predictPartyVictory(senate string) string {
	radiant, dire := make([]int, 0), make([]int, 0)
	n := len(senate)
	for i := 0; i < n; i++ {
		if senate[i] == 'R' {
			radiant = append(radiant, i)
		} else {
			dire = append(dire, i)
		}
	}
	for len(radiant) != 0 && len(dire) != 0 {
		r := radiant[0]
		d := dire[0]
		if r < d {
			radiant = append(radiant, r+n)
		} else {
			dire = append(dire, d+n)
		}
		radiant = radiant[1:]
		dire = dire[1:]
	}
	if len(radiant) != 0 {
		return "Radiant"
	}
	return "Dire"
}

// @lc code=end

