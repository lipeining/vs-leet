/*
 * @lc app=leetcode.cn id=677 lang=golang
 *
 * [677] 键值映射
 */

// @lc code=start

type TrieNode struct {
	isEnd bool
	val   int
	links map[string]*TrieNode
}
type MapSum struct {
	root *TrieNode
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{root: &TrieNode{links: make(map[string]*TrieNode)}}
}

func (this *MapSum) Insert(key string, val int) {
	node := this.root
	for i := 0; i < len(key); i++ {
		char := string(key[i])
		if _, ok := node.links[char]; !ok {
			tmpNode := &TrieNode{links: make(map[string]*TrieNode)}
			node.links[char] = tmpNode
		}
		node = node.links[char]
	}
	node.val = val
	node.isEnd = true
}

/** Inserts a word into the trie. */
func (this *MapSum) searchPrefix(word string) *TrieNode {
	node := this.root
	for i := 0; i < len(word); i++ {
		char := string(word[i])
		if _, ok := node.links[char]; ok {
			node = node.links[char]
		} else {
			return nil
		}
	}
	return node
}
func (node *TrieNode) Val() int {
	val := node.val
	fmt.Println(node)
	// if !node.isEnd {
	// 	val = 0
	// }
	for _, v := range node.links {
		val += v.Val()
	}
	return val
}

func (this *MapSum) Sum(prefix string) int {
	node := this.searchPrefix(prefix)
	if node == nil {
		return 0
	}
	return node.Val()
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
// @lc code=end
