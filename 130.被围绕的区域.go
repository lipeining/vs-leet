/*
 * @lc app=leetcode.cn id=130 lang=golang
 *
 * [130] 被围绕的区域
 */

// @lc code=start
func solve(board [][]byte)  {
	// 这种染色算法还未熟悉，必须找到模板方法
	m := len(board)
	if m == 0 {
		return
	}
	n := len(board[0])
	if n == 0 {
		return
	}
	for i:=0;i<m;i++ {
		for j:=0;j<n;j++ {
			egde := i==0 || j==0 || i==m-1 || j==n-1
			if egde && board[i][j] == 'O' {
				dfs(board,i,j)
			}
		}
	}
	for i:=0;i<m;i++ {
		for j:=0;j<n;j++ {
			if board[i][j] == '#' {
				board[i][j] = 'O'
			} else 	if board[i][j] == 'O' {
				board[i][j] = 'X'
			}	
		}
	}
}
func dfs(board [][]byte, x int, y int){
	if x<0 || x>=len(board) || y<0 || y>=len(board[0]) {
		return
	}
	if board[x][y] == 'X' || board[x][y] == '#' {
		return
	}
	board[x][y] = '#'
	dfs(board,x-1,y)
	dfs(board,x+1,y)
	dfs(board,x,y-1)
	dfs(board,x,y+1)
}
// @lc code=end

