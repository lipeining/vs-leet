/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	cache := make(map[int]*Node)
	var dfs func(n *Node) *Node
	dfs = func(n *Node) *Node {
		if n == nil {
			return nil
		}
		if c, ok := cache[n.Val]; ok {
			return c
		}
		size := len(n.Neighbors)
		t := &Node{
			Val:       n.Val,
			Neighbors: make([]*Node, size),
		}
		cache[n.Val] = t
		for i := 0; i < size; i++ {
			t.Neighbors[i] = dfs(n.Neighbors[i])
		}
		return t
	}
	ans := dfs(node)
	return ans
}