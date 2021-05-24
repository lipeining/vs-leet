/*
 * @lc app=leetcode.cn id=36 lang=golang
 *
 * [36] 有效的数独
 */

// @lc code=start
func isValidSudoku(board [][]byte) bool {
	n := 9
	row := make([]map[int]bool, n)
	col := make([]map[int]bool, n)
	box := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		row[i] = make(map[int]bool)
		col[i] = make(map[int]bool)
		box[i] = make(map[int]bool)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == '.' {
				continue
			}
			num := int(board[i][j] - '0')
			if _, ok := row[i][num]; ok {
				return false
			}
			if _, ok := col[j][num]; ok {
				return false
			}
			b := i/3*3 + j/3
			// fmt.Println("i-j", i, j, b)
			if _, ok := box[b][num]; ok {
				return false
			}
			row[i][num] = true
			col[j][num] = true
			box[b][num] = true
		}
	}
	return true
}

// @lc code=end

