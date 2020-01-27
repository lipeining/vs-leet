/*
 * @lc app=leetcode.cn id=71 lang=golang
 *
 * [71] 简化路径
 */

// @lc code=start
func simplifyPath(path string) string {
	stack := make([]string, 0)
	sp := strings.Split(path, "/")
	fmt.Println(sp, len(sp))

	for k:=0; k<len(sp);k++ {
		if sp[k]=="" {
			continue
		} else {
			length := len(stack)
			if sp[k] == "." {
				continue
			} else if sp[k] == ".." {
				if length > 0 {
					stack = stack[:length-1]
				}
			} else {
				stack = append(stack, sp[k])
			}
		}
	}
	ans := "/"
	for i:=0;i<len(stack);i++ {
		ans+= stack[i]
		if i != len(stack)-1 {
			ans+="/"
		}
	}
	return ans
}
// @lc code=end

