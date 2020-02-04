/*
 * @lc app=leetcode.cn id=676 lang=golang
 *
 * [676] 实现一个魔法字典
 */

// @lc code=start

type TrieNode struct {
	isEnd bool
	links map[string]*TrieNode
}
type MagicDictionary struct {
	root *TrieNode
}

/** Initialize your data structure here. */
func Constructor() MagicDictionary {
	return MagicDictionary{root: &TrieNode{links: make(map[string]*TrieNode)}}
}
func (this *MagicDictionary) insert(word string) {
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

/** Build a dictionary through a list of words */
func (this *MagicDictionary) BuildDict(dict []string) {
	for i := 0; i < len(dict); i++ {
		this.insert(dict[i])
	}
}

/** Returns if there is any word in the trie that equals to the given word after modifying exactly one character */
func (this *MagicDictionary) Search(word string) bool {
	// [[], [["hello","hallo","leetcode"]], ["hello"], ["hhllo"], ["hell"], ["leetcoded"]]
	// hello <=> hallo
	ans := false
	var dfs func(root *TrieNode, word string, flag bool)
	dfs = func(root *TrieNode, word string, flag bool) {
		if root == nil {
			return
		}
		if len(word) == 0 {
			if root.isEnd && flag {
				ans = true
			}
			return
		}
		char := string(word[0])
		next, first := root.links[char]
		if flag {
			if !first {
				return
			}
			dfs(next, word[1:], true)
		} else {
			if !first {
				for _, v := range root.links {
					dfs(v, word[1:], true)
				}
			} else {
				dfs(next, word[1:], false)
			}
		}
	}
	dfs(this.root, word, false)
	return ans
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dict);
 * param_2 := obj.Search(word);
 */
// @lc code=end
