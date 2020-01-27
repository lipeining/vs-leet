/*
 * @lc app=leetcode.cn id=173 lang=golang
 *
 * [173] 二叉搜索树迭代器
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
type BSTIterator struct {
    stack []int
}


func Constructor(root *TreeNode) BSTIterator {
	stack := make([]int, 0)
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode){
		if root == nil {
			return
		}
		dfs(root.Left)
		stack = append([]int{root.Val}, stack...)
		// fmt.Println(root.Val)
		dfs(root.Right)
	}
	dfs(root)
	// fmt.Println(stack)
	return BSTIterator{stack}
}


/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	length := len(this.stack)
	top := this.stack[length-1]
	this.stack = this.stack[:length-1]
	return top
}


/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
    return len(this.stack) > 0
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
// @lc code=end

