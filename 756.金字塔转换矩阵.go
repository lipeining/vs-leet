/*
 * @lc app=leetcode.cn id=756 lang=golang
 *
 * [756] 金字塔转换矩阵
 */

// @lc code=start
func pyramidTransition(bottom string, allowed []string) bool {
	// next 字符串表示上一层的模样，
	// 直到 next = 1 才终止 dfs
	// 需要记录 bottom, allowed, next, used index, 几个参数
	// 考虑边界条件
	var dfs func(bottom string, allowed []string, index int, next string) bool
	dfs = func(bottom string, allowed []string, index int, next string) bool {
		if len(bottom) == 1 {
			return true
		}
		if len(bottom) - len(next) > 1 {
			for i:=0;i<len(allowed);i++ {
				// 检查是否匹配当前的的开头
				flag := true
				for j:=0;j<2;j++ {
					if allowed[i][j] != bottom[index+j] {
						flag = false
						break
					}
				}
				if flag {
					nextNext := next+string(allowed[i][2])
					nextIndex := index+1
					fmt.Println(bottom, index, nextIndex, nextNext)
					if dfs(bottom, allowed, nextIndex, nextNext) {
						return true
					}
				}			
			}
			return false
		} else {
			return dfs(next, allowed, 0, "")
		}
	}
	ans := dfs(bottom, allowed, 0, "")
	return ans
}
// @lc code=end

