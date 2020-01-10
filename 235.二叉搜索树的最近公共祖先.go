/*
 * @lc app=leetcode.cn id=235 lang=golang
 *
 * [235] 二叉搜索树的最近公共祖先
 */

// @lc code=start
/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
//           // Value of current node or parent node.
// 		  int parentVal = root.val;

// 		  // Value of p
// 		  int pVal = p.val;
  
// 		  // Value of q;
// 		  int qVal = q.val;
  
// 		  if (pVal > parentVal && qVal > parentVal) {
// 			  // If both p and q are greater than parent
// 			  return lowestCommonAncestor(root.right, p, q);
// 		  } else if (pVal < parentVal && qVal < parentVal) {
// 			  // If both p and q are lesser than parent
// 			  return lowestCommonAncestor(root.left, p, q);
// 		  } else {
// 			  // We have found the split point, i.e. the LCA node.
// 			  return root;
// 		  }
  
//   作者：LeetCode
//   链接：https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/solution/er-cha-sou-suo-shu-de-zui-jin-gong-gong-zu-xian--2/
//   来源：力扣（LeetCode）
//   著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
	parentVal := root.Val
	pVal := p.Val
	qVal := q.Val
	if pVal > parentVal && qVal > parentVal {
		return lowestCommonAncestor(root.Right, p, q)
	} else if pVal < parentVal && qVal < parentVal {
		return lowestCommonAncestor(root.Left, p, q)
	} else {
		return root
	}
}
// @lc code=end

