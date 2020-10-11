/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

 func copyRandomList(head *Node) *Node {
	m := make(map[*Node]struct{})
	return dfs(head, m)
}
func dfs(head *Node, m map[*Node]struct{}) *Node {
	if head == nil {
		return nil
	}
	if n, ok := m[head]; ok {
		return n
	}
	node := &Node{Val: head.Val}
	m[head] = node
	node.Next = dfs(head.Next, m)
	node.Random = dfs(head.Random, m)
	return node
}