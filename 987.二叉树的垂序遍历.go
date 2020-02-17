import "sort"

/*
 * @lc app=leetcode.cn id=987 lang=golang
 *
 * [987] 二叉树的垂序遍历
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func verticalTraversal(root *TreeNode) [][]int {
	m := make(map[int]map[int][]int)
	var dfs func(root *TreeNode, x, y int)
	dfs = func(root *TreeNode, x, y int) {
		if root == nil {
			return
		}
		if _, ok := m[x]; !ok {
			m[x] = make(map[int][]int)
		}
		if _, ok := m[x][y]; !ok {
			m[x][y] = make([]int, 0)
		}
		m[x][y] = append(m[x][y], root.Val)
		dfs(root.Left, x-1, y-1)
		dfs(root.Right, x+1, y-1)
	}
	dfs(root, 0, 0)
	ans := make([][]int, 0)
	xRange := make([]int, 0)
	for k, _ := range m {
		xRange = append(xRange, k)
	}
	sort.Ints(xRange)
	// fmt.Println("xRange", xRange)
	for _, x := range xRange {
		yRange := make([]int, 0)
		for k, _ := range m[x] {
			yRange = append(yRange, k)
		}
		// sort.Ints(yRange)
		sort.Sort(sort.Reverse(sort.IntSlice(yRange)))
		// fmt.Println("yRange", yRange)
		Ylist := make([]int, 0)
		for _, y := range yRange {
			list,_ := m[x][y]
			sort.Ints(list)
			Ylist = append(Ylist, list...)
		}
		ans = append(ans, Ylist)
	}
	return ans
}

// @lc code=end
