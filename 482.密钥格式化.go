/*
 * @lc app=leetcode.cn id=482 lang=golang
 *
 * [482] 密钥格式化
 */

// @lc code=start
func licenseKeyFormatting(S string, K int) string {
	length := len(S)
	res := ""
	counter := 0
	// 可以最简单的先得到全部的字符串，再分割
	for i:= length-1;i>=0;i-- {
		s:= string(S[i])
		if s != "-" {
			res = strings.ToUpper(s) + res
			counter++
		}
		if counter == K {
			counter = 0
			// if i != 0 {
			// 	res = "-" + res
			// }
			res = "-" + res
		}
	}
	fmt.Println(res)
	// res = strings.TrimPrefix(res, string('-'))
	i:= 0;
	for  i=0;i < len(res); i++{
		if string(res[i]) != "-" {
			break;	
		}			
	}
	if i != 0 {
		return res[i:]
	}
	// 
	return res
}
// @lc code=end

