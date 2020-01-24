/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */

// @lc code=start
func exist(board [][]byte, word string) bool {
	m := len(board)
	if m == 0 {
		return false
	}
	n := len(board[0])
	used := make([][]bool, m)
	// fmt.Println(used)
	for i:=0; i<m;i++ {
		sub := make([]bool, n)
		used[i]=sub
	}
	// fmt.Println(used)
	ans := false
	var dfs func(board [][]byte, word string, x int, y int, used [][]bool, path string)
	dfs = func(board [][]byte, word string, x int, y int, used [][]bool, path string) {
		if ans {
			return
		}
		if len(word) == len(path) {
			if word == path {
				ans = true
			}
			fmt.Println(path)
			return
		}
		if len(path) > len(word) {
			return
		}
		if x < 0 || x >= m {
			return
		}
		if y < 0 || y >= n {
			return
		}
		if used[x][y] {
			return
		}
		// 向四个方向散步吧
		cur := path + string(board[x][y])
		// 这里判断 cur 是否是 word 前缀
		if !strings.HasPrefix(word, cur) {
			return
		}
		used[x][y] = true
		dfs(board, word, x-1, y, used, cur)
		dfs(board, word, x+1, y, used, cur)
		dfs(board, word, x, y-1, used, cur)
		dfs(board, word, x, y+1, used, cur)
		used[x][y] = false
	}
	for i:=0;i<m;i++ {
		for j:=0;j<n;j++ {
			dfs(board, word, i, j, used, "")
		}
	}
	return ans
}
// @lc code=end

