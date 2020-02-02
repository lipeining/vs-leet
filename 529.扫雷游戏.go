/*
 * @lc app=leetcode.cn id=529 lang=golang
 *
 * [529] 扫雷游戏
 */

// @lc code=start
func updateBoard(board [][]byte, click []int) [][]byte {
	// 完全抄别人思路代码
	m := len(board)
	if m == 0 {
		return board
	}
	n := len(board[0])
	if n == 0 {
		return board
	}
	x0 := click[0]
	y0 := click[1]
	if board[x0][y0] == 'M' {
		board[x0][y0] = 'X'
		return board
	}
	dx := [8]int{-1,-1,-1,1,1,1,0,0}
	dy := [8]int{-1,1,0,-1,1,0,-1,1}
	var getNumber func(board [][]byte, x int, y int) int 
	getNumber = func(board [][]byte, x int, y int) int {
		num:=0
		for i:=0;i<len(dx);i++ {
			xt := x+dx[i]
			yt := y+dy[i]
			if xt<0 || xt>=m || yt<0 || yt>=n {
				continue
			}
			if board[xt][yt] == 'M' || board[xt][yt] == 'X' {
				num++
			}
		}
		return num
	}
	byteMap := map[int]byte{
		1:'1',
		2:'2',
		3:'3',
		4:'4',
		5:'5',
		6:'6',
		7:'7',
		8:'8',
	}
	var dfs func(board [][]byte, x int, y int)
	dfs = func(board [][]byte, x int, y int) {
		if x<0 || x>=m || y<0 || y>=n {
			return
		}
		if board[x][y] != 'E' {
			return
		}
		num := getNumber(board, x, y)
		if num == 0 {
			board[x][y] = 'B'
			for i:=0;i<len(dx);i++ {
				xt := x+dx[i]
				yt := y+dy[i]
				dfs(board, xt, yt)
			}
		} else {
			// board[x][y] = strconv.Itoa(num)
			// board[x][y] = uint8(num)
			board[x][y] = byteMap[num]
		}
	}
	dfs(board, x0, y0)
	return board
}
// @lc code=end

