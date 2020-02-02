/*
 * @lc app=leetcode.cn id=721 lang=golang
 *
 * [721] 账户合并
 */

// @lc code=start

// 其实整数和字符串是一样的，不过整数的并查集使用了 array 来简化 map
// map[int]int <=> []int
func find(root map[string]string, node string) string {
	if root[node]==node {
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
	if pRoot==qRoot {
		return
	}
	root[qRoot] = pRoot
}

func accountsMerge(accounts [][]string) [][]string {
	root := make(map[string]string)
	emailToName := make(map[string]string)
	// 初始化单个的 root 
	// 这里需要考虑 emailToName 覆盖问题吧
	for _,a := range accounts {
		for i:=1;i<len(a);i++ {
			emailToName[a[i]] = a[0]
			root[a[i]] = a[i]
		}
	}
	// fmt.Println(emailToName)
	// fmt.Println(root)
	for _,a := range accounts {
		for i:=1;i<len(a)-1;i++ {
			union(root, a[i], a[i+1])
		}
	}
	// fmt.Println(root)
	m := make(map[string]map[string]bool)
	for _,a := range accounts {
		for _,s := range a[1:] {
			r := find(root, s)
			if _,ok := m[r]; !ok {
				m[r] = make(map[string]bool)
			}
			m[r][s] = true
		}
	}
	ans := make([][]string, 0)
	for r,sub := range m {
		list := make([]string, 0)
		for s,_ := range sub {
			list = append(list, s)
		}
		sort.Strings(list)
		// fmt.Println(list)
		name,_ := emailToName[r]
		list = append([]string{name}, list...)
		ans = append(ans, list)
	}
	return ans
}

// @lc code=end
// type Node struct {
// 	parent *Node
// 	account string
// }

// func NewNode(account) *NOde {
// 	n := &Node{}
// 	n.parent = n
// 	n.account = account
// 	return n
// }
// func (n *Node)find() *Node{
// 	if n == nil {
// 		return n
// 	}
// 	if n!=n.parent {
// 		return n.parent.find()
// 	}
// 	return n
// }
// type UnionFindSet struct {
// 	m map[string]Node
// }

// func NewUnionFindSet() *UnionFindSet {
// 	return &UnionFindSet{}
// }
// func (ufs *UnionFindSet)find(str string)*Node{
// 	return ufs.m[str].find()
// }
// func (ufs *UnionFindSet)union(p string, q string) {
// 	pRoot := ufs.m[p].find()
// 	qRoot := ufs.m[q].find()
// 	if pRoot == qRoot {
// 		return
// 	}
// 	pRoot.parent = qRoot
// }
// type UnionSet struct {
// 	Parent []int // 每个并查集元素都有一个parent数组指向它自己
// 	Rank   []int // 维护每个元素在并查集树中的排名
// }
// // 初始化
// func (this *UnionSet) Init(n int) {
// 	this.Parent = make([]int, n)
// 	this.Rank = make([]int, n)
// 	for i := 0; i < n; i++ {
// 		this.Parent[i] = i
// 		this.Rank[i] = 1
// 	}
// }

// // 查询find的根节点
// /*
// 	不断查找ele的parent,直到parent[ele] = ele
// */
// func (this *UnionSet) Find(ele int) int {
// 	return this.find(ele)
// }
// func (this *UnionSet) find(ele int) int {
// 	for ele != this.Parent[ele]  {
// 		this.Parent[ele] = this.Parent[this.Parent[ele]]
// 		ele = this.Parent[ele]
// 	}
// 	return ele
// }

// // 合并
// // 首先分别找到p, q的领头节点、如果相同则直接返回
// // 否则将基本一个的parent置为另一个
// // 最后将并查集中独立集合数量减1
// func (this *UnionSet) Union(p, q int) {
// 	pRoot := this.find(p)
// 	qRoot := this.find(q)
// 	if pRoot == qRoot {
// 		return
// 	}
// 	if this.Rank[pRoot] > this.Rank[qRoot] {
// 		this.Parent[qRoot] = pRoot
// 	} else if this.Rank[pRoot] < this.Rank[qRoot]{
// 		this.Parent[pRoot] = qRoot
// 	} else {
// 		this.Parent[pRoot] = qRoot
// 		this.Rank[qRoot]++
// 	}
// }
