/*
 * @lc app=leetcode.cn id=52 lang=golang
 *
 * [52] N皇后 II
 */

// @lc code=start
func totalNQueens(n int) int {
	ans := make([][]string, 0)
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	check := func(path []string, row, col int) string {
		for i, line := range path {
			points := strings.Split(line, "")
			for j, char := range points {
				if char == "Q" {
					if abs(i-row) == abs(j-col) {
						return ""
					}
					if j == col {
						return ""
					}
				}
			}
		}
		ans := ""
		j := 0
		for j < col {
			ans += "."
			j++
		}
		ans += "Q"
		j++
		for j < n {
			ans += "."
			j++
		}
		return ans
	}
	var dfs func(index int, path []string)
	dfs = func(index int, path []string) {
		if index == n {
			p := make([]string, len(path))
			copy(p, path)
			ans = append(ans, p)
			return
		}
		for col := 0; col < n; col++ {
			if line := check(path, index, col); line != "" {
				path = append(path, line)
				dfs(index+1, path)
				path = path[:len(path)-1]
			}
		}
	}
	path := make([]string, 0)
	dfs(0, path)
	return len(ans)
	// return ans
}
// @lc code=end

