package segmentTree

type Node struct {
	Val   int
	Ls    int
	Rs    int
	Left  *Node
	Right *Node
}

type SegmentTree struct {
	N    int
	Vals []int
	Root *Node
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func build(l, r int, vals []int) *Node {
	if r-l < 1 {
		return nil
	}
	node := &Node{Ls: l, Rs: r}
	if r-l == 1 {
		node.Val = vals[l]
	} else {
		m := (l + r) >> 1
		node.Left = build(l, m, vals)
		node.Right = build(m, r, vals)
		node.Val = min(node.Left.Val, node.Right.Val)
	}
	return node
}

func NewSegmentTree(n int, vals []int) *SegmentTree {
	root := build(0, n, vals)
	return &SegmentTree{N: n, Vals: vals, Root: root}
}

func (self *SegmentTree) Update(k, v int) {
	update(self.Root, k, v)
}

func update(u *Node, k, v int) {
	if u == nil {
		return
	}
	if u.Ls <= k && k < u.Rs {
		u.Val = min(u.Val, v)
		m := (u.Ls + u.Rs) >> 1
		if k < m {
			update(u.Left, k, v)
		} else {
			update(u.Right, k, v)
		}
	}
}

func (self *SegmentTree) Query(l, r int) int {
	return query(self.Root, l, r)
}

func query(u *Node, l, r int) int {
	if l <= u.Ls && r >= u.Rs {
		return u.Val
	} else if r <= u.Left.Rs {
		return query(u.Left, l, r)
	} else if l >= u.Right.Ls {
		return query(u.Right, l, r)
	}
	return min(query(u.Left, l, r), query(u.Right, l, r))
}
