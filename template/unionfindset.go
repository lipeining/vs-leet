package template

// parent := [10001]int{}
// rank   := [10001]int{}
var parent []int = make([]int, 10001)
var rank []int = make([]int, 10001)

// 递归压缩可能会导致栈溢出
func findStack(x int) int {
	if x != parent[x] {
		parent[x] = find(parent[x])
	}
	return parent[x]
}
func find(x int) int {
	var j, k, r int
	r = x
	for r != parent[r] {
		r = parent[r]
	}
	k = x
	for k != r {
		j = parent[k]
		parent[k] = r
		k = j
	}
	return r
}
func union(x, y int) {
	x = find(x)
	y = find(y)
	if x == y {
		return
	}
	if rank[x] >= rank[y] {
		parent[y] = x
		rank[x] += rank[y]
	} else {
		parent[x] = y
		rank[y] += rank[x]
	}
}

// func find(root map[string]string, node string) string {
// 	if root[node] == node {
// 		return node
// 	}
// 	return find(root, root[node])
// 	// for node != root[node] {
// 	// 	root[node] = root[root[node]]
// 	// 	node = root[node]
// 	// }
// 	// return node
// }
// func union(root map[string]string, p string, q string) {
// 	pRoot := find(root, p)
// 	qRoot := find(root, q)
// 	if pRoot == qRoot {
// 		return
// 	}
// 	root[qRoot] = pRoot
// }

// // 399 547 721
