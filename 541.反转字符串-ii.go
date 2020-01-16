/*
 * @lc app=leetcode.cn id=541 lang=golang
 *
 * [541] 反转字符串 II
 */

// @lc code=start
func reverseStr(s string, k int) string {
	length := len(s)
	// fmt.Println(s, s[k:])
	i:=0;
	res:= ""
	for i < length {
		mid := i+k
		end := i+2*k
		if mid > length {
			// less than k
			res += helper(s[i:length])
		} else if mid <= length && end>length {
			res += helper(s[i:mid]) + s[mid:length]
		} else {
			res += helper(s[i:mid]) + s[mid:end]
		}
		i += 2*k
	}
	return res
}
func helper(s string) string {
	length := len(s)
	res := ""
	for i:=0; i<length;i++ {
		// t := s[i]
		// s[i] = s[length-i-1]
		// s[length-i-1] = t
		res+=string(s[length-i-1])
	}
	fmt.Println(res)
	return res
}
// @lc code=end

