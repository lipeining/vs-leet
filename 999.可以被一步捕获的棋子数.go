/*
 * @lc app=leetcode.cn id=999 lang=golang
 *
 * [999] 可以被一步捕获的棋子数
 */

// @lc code=start
func numRookCaptures(board [][]byte) int {
	m := len(board)
	n := len(board[0])
	ans := 0
	var dfs func(i, j int, dir string)
	dfs = func(i, j int, dir string) {
		if i < 0 || i == m || j < 0 || j == n {
			return
		}
		if board[i][j] == 'p' {
			ans++
			return
		}
		if board[i][j] == 'B' {
			return
		}
		if dir == "left" {
			dfs(i-1, j, dir)
		} else if dir == "right" {
			dfs(i+1, j, dir)
		} else if dir == "up" {
			dfs(i, j-1, dir)
		} else {
			dfs(i, j+1, dir)
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'R' {
				dfs(i, j, "left")
				dfs(i, j, "right")
				dfs(i, j, "up")
				dfs(i, j, "down")
			}
		}
	}
	return ans
}

// @lc code=end

