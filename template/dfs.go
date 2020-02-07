package template

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// dfs 部分的模板代码
// 105
func buildTree(preorder []int, inorder []int) *TreeNode {
	// 知道的思路是 前序中找到中序中的左子树分支和右子树的划分点
	length := len(preorder)
	var root *TreeNode
	if length == 0 {
		return root
	}
	head := &TreeNode{Val: preorder[0]}
	mid := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			mid = i
			break
		}
	}
	head.Left = buildTree(preorder[1:mid+1], inorder[:mid])
	head.Right = buildTree(preorder[mid+1:], inorder[mid+1:])
	return head
}

// // 106
// func buildTree(inorder []int, postorder []int) *TreeNode {
// 	length := len(postorder)
// 	var root *TreeNode
// 	if length == 0 {
// 		return root
// 	}
// 	head := &TreeNode{Val: postorder[length-1]}
// 	mid := 0
// 	for i := 0; i < len(inorder); i++ {
// 		if inorder[i] == postorder[length-1] {
// 			mid = i
// 			break
// 		}
// 	}
// 	head.Left = buildTree(inorder[:mid], postorder[:mid])
// 	head.Right = buildTree(inorder[mid+1:], postorder[mid:length-1])
// 	return head
// }

// 113
func pathSum(root *TreeNode, sum int) [][]int {
	ans := make([][]int, 0)
	var dfs func(root *TreeNode, path []int, target int)
	dfs = func(root *TreeNode, path []int, target int) {
		if root == nil {
			return
		}
		if isLeaf(root) {
			if root.Val == target {
				path = append(path, root.Val)
				ans = append(ans, path)
			}
		} else {
			t1 := make([]int, len(path))
			t2 := make([]int, len(path))
			copy(t1, path)
			copy(t2, path)
			t1 = append(t1, root.Val)
			t2 = append(t2, root.Val)
			dfs(root.Left, t1, target-root.Val)
			dfs(root.Right, t2, target-root.Val)
		}
	}
	path := make([]int, 0)
	dfs(root, path, sum)
	return ans
}
func isLeaf(t *TreeNode) bool {
	return t.Left == nil && t.Right == nil
}

// 129
func sumNumbers(root *TreeNode) int {
	ans := 0
	var dfs func(root *TreeNode, sum int)
	dfs = func(root *TreeNode, sum int) {
		if root == nil {
			return
		}
		sum = sum*10 + root.Val
		if root.Left == nil && root.Right == nil {
			ans += sum
			return
		} else {
			dfs(root.Left, sum)
			dfs(root.Right, sum)
		}
	}
	dfs(root, 0)
	return ans
}

// 130
func solve(board [][]byte) {
	// 这种染色算法还未熟悉，必须找到模板方法
	m := len(board)
	if m == 0 {
		return
	}
	n := len(board[0])
	if n == 0 {
		return
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			egde := i == 0 || j == 0 || i == m-1 || j == n-1
			if egde && board[i][j] == 'O' {
				dfs(board, i, j)
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == '#' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}
func dfs(board [][]byte, x int, y int) {
	if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) {
		return
	}
	if board[x][y] == 'X' || board[x][y] == '#' {
		return
	}
	board[x][y] = '#'
	dfs(board, x-1, y)
	dfs(board, x+1, y)
	dfs(board, x, y-1)
	dfs(board, x, y+1)
}

// 200
func numIslands(grid [][]byte) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	var dfs func(grid [][]byte, x int, y int)
	dfs = func(grid [][]byte, x int, y int) {
		if x < 0 || x >= m || y < 0 || y >= n {
			return
		}
		if grid[x][y] == '0' {
			return
		}
		grid[x][y] = '0'
		dfs(grid, x-1, y)
		dfs(grid, x+1, y)
		dfs(grid, x, y-1)
		dfs(grid, x, y+1)
	}
	counter := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				counter++
				dfs(grid, i, j)
			}
		}
	}
	return counter
}

// 207 210 547

// 529 695
func maxAreaOfIsland(grid [][]int) int {
	ans := 0
	m := len(grid)
	if m == 0 {
		return ans
	}
	n := len(grid[0])
	if n == 0 {
		return ans
	}
	var dfs func(grid [][]int, x int, y int, seen [][]bool) int
	dfs = func(grid [][]int, x int, y int, seen [][]bool) int {
		if x < 0 || x >= m || y < 0 || y >= n {
			return 0
		}
		if grid[x][y] == 0 {
			return 0
		}
		if seen[x][y] {
			return 0
		}
		seen[x][y] = true
		return 1 + dfs(grid, x-1, y, seen) + dfs(grid, x+1, y, seen) + dfs(grid, x, y-1, seen) + dfs(grid, x, y+1, seen)
	}
	seen := make([][]bool, m)
	for i := 0; i < m; i++ {
		seen[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				tmp := dfs(grid, i, j, seen)
				if tmp > ans {
					ans = tmp
				}
			}
		}
	}
	return ans
}

// 721
