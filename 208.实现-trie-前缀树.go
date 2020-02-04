/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 */

// @lc code=start
type TrieNode struct {
	isEnd bool
	links map[string]*TrieNode
}
type Trie struct {
	root *TrieNode
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{root: &TrieNode{links: make(map[string]*TrieNode)}}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this.root
	for i := 0; i < len(word); i++ {
		char := string(word[i])
		if _, ok := node.links[char]; !ok {
			tmpNode := &TrieNode{links: make(map[string]*TrieNode)}
			node.links[char] = tmpNode
		}
		node = node.links[char]
	}
	node.isEnd = true
}
func (this *Trie) searchPrefix(word string) *TrieNode {
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

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this.searchPrefix(word)
	return node != nil && node.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this.searchPrefix(prefix)
	return node != nil
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end
