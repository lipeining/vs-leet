/*
 * @lc app=leetcode.cn id=419 lang=golang
 *
 * [419] 甲板上的战舰
 */

// @lc code=start
func countBattleships(board [][]byte) int {
	n := len(board)
	m := len(board[0])
	cnt := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 'X' && (i == 0 || board[i-1][j] == '.') && (j == 0 || board[i][j-1] == '.') {
				cnt++
			}
		}
	}

	return cnt
}

// @lc code=end

