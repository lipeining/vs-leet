package template

func find(root map[string]string, node string) string {
	if root[node] == node {
		return node
	}
	return find(root, root[node])
	// for node != root[node] {
	// 	root[node] = root[root[node]]
	// 	node = root[node]
	// }
	// return node
}
func union(root map[string]string, p string, q string) {
	pRoot := find(root, p)
	qRoot := find(root, q)
	if pRoot == qRoot {
		return
	}
	root[qRoot] = pRoot
}

// 399 547 721
