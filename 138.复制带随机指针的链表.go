/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	m := make(map[*Node]*Node)
	var dfs func(head *Node) *Node
	dfs = func(head *Node) *Node {
		if head == nil {
			return nil
		}
		if n, ok := m[head]; ok {
			return n
		}
		node := &Node{Val: head.Val}
		m[head] = node
		node.Next = dfs(head.Next)
		node.Random = dfs(head.Random)
		return node
	}
	ans := dfs(head)
	return ans
}