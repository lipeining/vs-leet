/*
 * @lc app=leetcode.cn id=297 lang=golang
 *
 * [297] 二叉树的序列化与反序列化
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

 type Codec struct {
    
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    if root == nil {
        return "x"
    }
    return strconv.Itoa(root.Val) + "," +  this.serialize(root.Left)+ "," + this.serialize(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {    
    list := strings.Split(data, ",")
    var dfs func() *TreeNode
    dfs = func() *TreeNode {
        val := list[0]
        list = list[1:]
        if val == "x" {
            return nil
        }
        Val,_ := strconv.Atoi(val)
        root := &TreeNode{Val: Val}
        root.Left = dfs()
        root.Right = dfs()
        return root
    }
    return dfs()
}


/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
// @lc code=end

