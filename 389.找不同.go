/*
 * @lc app=leetcode.cn id=389 lang=golang
 *
 * [389] 找不同
 */

// @lc code=start
func findTheDifference(s string, t string) byte {
	num := 0
	for _, char := range s {
		num ^= int(char-'a')
	}
	for _, char := range t {
		num ^= int(char-'a')
	}
	num ^= 0
	return byte(num+97)
	// smap := make(map[rune]int)
	// for _,c := range t {
	// 	n,ok := smap[c]
	// 	if ok {
	// 		smap[c] = n+1
	// 	} else {
	// 		smap[c]=1
	// 	}
	// }
	// for _,c := range s {
	// 	n,ok := smap[c]
	// 	if ok {
	// 		smap[c] = n - 1
	// 	} else {
	// 		return byte(c)
	// 	}
	// }
	// for k,v := range smap {
	// 	if v > 0 {
	// 		return byte(k)
	// 	}
	// }
	// return t[0]
}
// @lc code=end

