import "fmt"

/*
 * @lc app=leetcode.cn id=211 lang=golang
 *
 * [211] 添加与搜索单词 - 数据结构设计
 */

// type Trie struct {
// 	root *TrieNode
// }

// /** Initialize your data structure here. */
// func Constructor() Trie {
// 	return Trie{root: &TrieNode{links: make(map[string]*TrieNode)}}
// }

// /** Inserts a word into the trie. */
// func (this *Trie) Insert(word string) {
// 	node := this.root
// 	for i := 0; i < len(word); i++ {
// 		char := string(word[i])
// 		if _, ok := node.links[char]; !ok {
// 			tmpNode := &TrieNode{links: make(map[string]*TrieNode)}
// 			node.links[char] = tmpNode
// 		}
// 		node = node.links[char]
// 	}
// 	node.isEnd = true
// }

// /** Returns if the word is in the trie. */
// func (this *Trie) Search(word string) bool {
// 	node := this.searchPrefix(word)
// 	return node != nil && node.isEnd
// }

// /** Returns if there is any word in the trie that starts with the given prefix. */
// func (this *Trie) StartsWith(prefix string) bool {
// 	node := this.searchPrefix(prefix)
// 	return node != nil
// }

// @lc code=start

type TrieNode struct {
	isEnd bool
	links map[string]*TrieNode
}

type WordDictionary struct {
	root *TrieNode
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{root: &TrieNode{links: make(map[string]*TrieNode)}}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
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

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {

	ans := false
	var dfs func(root *TrieNode, word string)
	dfs = func(root *TrieNode, word string) {
		// fmt.Println(root, word)
		if root == nil {
			return
		}
		if len(word) == 0 {
			if root.isEnd {
				ans = true
			}
			return
		}
		char := string(word[0])
		if char == "." {
			for _, v := range root.links {
				dfs(v, word[1:])
			}
		} else {
			// fmt.Println("here,here", char, root.links[char])
			dfs(root.links[char], word[1:])
		}
	}
	dfs(this.root, word)
	return ans
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
// @lc code=end

// type WordDictionary struct {
//     words []string
// }

// /** Initialize your data structure here. */
// func Constructor() WordDictionary {
//     return WordDictionary{}
// }

// /** Adds a word into the data structure. */
// func (this *WordDictionary) AddWord(word string)  {
//     this.words = append(this.words, word)
// }

// /** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
// func (this *WordDictionary) Search(word string) bool {
//     for i:=0;i<len(this.words);i++ {
// 		if helper(this.words[i], word) {
// 			return true
// 		}
// 	}
// 	return false
// }
// func helper(word string, pattern string) bool {
// 	// var dfs func(word string, pattern string, path string)
// 	// dfs = func(word string, pattern string, path string) {
// 	// 	if path == word {
// 	// 		ans = true
// 	// 		return
// 	// 	}
// 	// 	if len(path) >= len(word) {
// 	// 		return
// 	// 	}
// 	// 	for i:=0;i< len(pattern);i++ {

// 	// 	}
// 	// }
// 	// dfs(word, pattern, "")
// 	if len(word) != len(pattern) {
// 		return false
// 	}
// 	for i:=0;i<len(word);i++ {
// 		if pattern[i] == '.' {
// 			continue
// 		} else if pattern[i] != word[i] {
// 			return false
// 		}
// 	}
// 	return true
// }
