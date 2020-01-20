/*
 * @lc app=leetcode.cn id=1189 lang=golang
 *
 * [1189] “气球” 的最大数量
 */

// @lc code=start
func maxNumberOfBalloons(text string) int {
    cMap := map[byte]int{
		'b':0,
		'a':0,
		'l':0,
		'o':0,
		'n':0,
	}
	for i:=0;i<len(text);i++ {
		c,ok := cMap[text[i]]
		if ok {
			cMap[text[i]] = c + 1
		}
	}
	fmt.Println(cMap)
	ans := len(text)
	for k,v := range cMap {
		if v == 0 {
			ans = 0
			break
		}
		if k == 'b' || k == 'a' || k=='n' {
			if ans > v {
				ans = v
			}
		} else {
			if ans > v/2 {
				ans = v/2
			}
		}
	}
	return ans
}
// @lc code=end

