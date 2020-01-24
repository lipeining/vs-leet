/*
 * @lc app=leetcode.cn id=93 lang=golang
 *
 * [93] 复原IP地址
 */

// @lc code=start
func restoreIpAddresses(s string) []string {
	ans := make([]string, 0)
	var dfs func(s string, index int, path string)
	dfs = func(s string, index int, path string) {
		if index == 4 {
			if s != "" {
				return
			}
			ans = append(ans, path)
			return
		}
		for k:=1;k<=3;k++ {
			if len(s) < k {
				break
			}
			val,_ := strconv.Atoi(s[0:k])
			if val > 255 || k != len(strconv.Itoa(val)) {
				continue
			}
			t := path + s[0:k]
			if index != 3 {
				t += "."
			}
			dfs(s[k:], index+1, t)
		}
		// 分步行走，1位或者2位或者3位
	}
	path := ""
	dfs(s, 0, path)
	return ans
	// ans := make([]string, 0)
	// var dfs func(s string, index int, path []string)
	// dfs = func(s string, index int, path []string) {
	// 	if len(path) == 4 {
	// 		if index < len(s) {
	// 			return
	// 		}
	// 		for i:=0;i<len(path);i++ {
	// 			if !check(path[i]) {
	// 				return
	// 			}
	// 		}
	// 		ans = append(ans, strings.Join(path, "."))
	// 		return
	// 	}
	// 	if index>= len(s) {
	// 		return
	// 	}
	// 	// 分步行走，1位或者2位或者3位
	// }
	// path := make([]string, 0)
	// dfs(s, 0, path)
	// return ans
}
func check(s string)bool {
	// 将字符串转为整数，判断是否符合要求
	t,_ := strconv.Atoi(s)
	return t>=0 && t<=255
}
// @lc code=end

