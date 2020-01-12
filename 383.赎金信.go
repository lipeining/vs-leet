/*
 * @lc app=leetcode.cn id=383 lang=golang
 *
 * [383] 赎金信
 */

// @lc code=start
func canConstruct(ransomNote string, magazine string) bool {
    if len(magazine) < len(ransomNote) {
		return false
	}
	smap := make(map[rune]int)
	for _,c := range ransomNote {
		n,ok := smap[c]
		if ok {
			smap[c] = n+1
		} else {
			smap[c]=1
		}
	}
	for _,c := range magazine {
		n,ok := smap[c]
		if ok {
			smap[c] = n - 1
		}
	}
	for _,v := range smap {
		if v > 0 {
			return false
		}
	}
	return true
}
// @lc code=end

