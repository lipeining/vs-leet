/*
 * @lc app=leetcode.cn id=37 lang=golang
 *
 * [37] 解数独
 */

// @lc code=start
func solveSudoku(board [][]byte) {
	rows, cols := make([][]bool, 9), make([][]bool, 9)
	for i := 0; i < 9; i++ {
		rows[i] = make([]bool, 10)
		cols[i] = make([]bool, 10)
	}
	box := make([][][]bool, 3)
	for i := 0; i < 3; i++ {
		box[i] = make([][]bool, 3)
		for j := 0; j < 3; j++ {
			box[i][j] = make([]bool, 10)
		}
	}
	byteMap := map[int]byte{
		1: '1',
		2: '2',
		3: '3',
		4: '4',
		5: '5',
		6: '6',
		7: '7',
		8: '8',
		9: '9',
	}
	spaces := make([][]int, 0)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == '.' {
				spaces = append(spaces, []int{i, j})
			} else {
				x := board[i][j] - '0'
				rows[i][x] = true
				cols[j][x] = true
				box[i/3][j/3][x] = true
			}
		}
	}
	var dfs func(index int) bool
	dfs = func(index int) bool {
		if index == len(spaces) {
			return true
		}
		space := spaces[index]
		x, y := space[0], space[1]
		for num := 1; num <= 9; num++ {
			if !rows[x][num] && !cols[y][num] && !box[x/3][y/3][num] {
				rows[x][num] = true
				cols[y][num] = true
				box[x/3][y/3][num] = true
				board[x][y] = byteMap[num]
				if dfs(index + 1) {
					return true
				}
				board[x][y] = '.'
				rows[x][num] = false
				cols[y][num] = false
				box[x/3][y/3][num] = false
			}
		}
		return false
	}
	dfs(0)
}

// @lc code=end

