/*
 * @lc app=leetcode.cn id=427 lang=golang
 *
 * [427] 建立四叉树
 */

// @lc code=start
/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */

func construct(grid [][]int) *Node {
	var dfs func(x1, y1, x2, y2 int) *Node
	dfs = func(x1, y1, x2, y2 int) *Node {
		if x1 == x2-1 {
			val := false
			if grid[x1][y1] == 1 {
				val = true
			}
			return &Node{
				Val:    val,
				IsLeaf: true,
			}
		}
		mx := x1 + (x2-x1)/2
		my := y1 + (y2-y1)/2
		tl := dfs(x1, y1, mx, my)
		tr := dfs(x1, my, mx, y2)
		bl := dfs(mx, y1, x2, my)
		br := dfs(mx, my, x2, y2)
		if tl.IsLeaf && tr.IsLeaf && bl.IsLeaf && br.IsLeaf {
			if tl.Val && tr.Val && bl.Val && br.Val {
				return &Node{
					Val:    true,
					IsLeaf: true,
				}
			}
			if !tl.Val && !tr.Val && !bl.Val && !br.Val {
				return &Node{
					Val:    false,
					IsLeaf: true,
				}
			}
		}
		return &Node{
			Val:         false,
			IsLeaf:      false,
			TopLeft:     tl,
			TopRight:    tr,
			BottomLeft:  bl,
			BottomRight: br,
		}
	}
	n := len(grid)
	return dfs(0, 0, n, n)
}

// @lc code=end

