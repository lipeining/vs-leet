package lc

// type Trie struct {
// 	next   map[rune]*Trie
// 	isWord bool
// }

// /** Initialize your data structure here. */
// func NewTrie() *Trie {
// 	root := new(Trie)
// 	root.next = make(map[rune]*Trie)
// 	root.isWord = false
// 	return root
// }

// /** Inserts a word into the trie. */
// func (this *Trie) Insert(word string) {
// 	for _, v := range word {
// 		if this.next[v] == nil {
// 			node := new(Trie)
// 			//子节点数量为26
// 			node.next = make(map[rune]*Trie)
// 			//初始化节点单词标志为假
// 			node.isWord = false
// 			this.next[v] = node
// 		}
// 		this = this.next[v]
// 		// this.count++ 统计每个单词前缀出现的次数（也包括统计每个单词出现的次数）
// 	}
// 	this.isWord = true
// }

// /** Returns if the word is in the trie. */
// func (this *Trie) Search(word string) bool {
// 	for _, v := range word {
// 		if this.next[v] == nil {
// 			return false
// 		}
// 		this = this.next[v]
// 	}
// 	return this.isWord
// }

// /** Returns if there is any word in the trie that starts with the given prefix. */
// func (this *Trie) StartsWith(prefix string) bool {
// 	for _, v := range prefix {
// 		if this.next[v] == nil {
// 			return false
// 		}
// 		this = this.next[v]
// 	}
// 	return true
// }

// type WordDictionary struct {
// 	root *Trie
// }

// /** Initialize your data structure here. */
// func Constructor() WordDictionary {
// 	root := NewTrie()
// 	return WordDictionary{root}
// }

// /** Adds a word into the data structure. */
// func (this *WordDictionary) AddWord(word string) {
// 	this.root.Insert(word)
// }

// /** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
// func (this *WordDictionary) Search(word string) bool {
// 	root := this.root
// 	return root.patternSearch(word)
// }
// func (this *Trie) patternSearch(word string) bool {
// 	for index, v := range word {
// 		if v == '.' {
// 			flag := false
// 			for _, nv := range this.next {
// 				flag = flag || nv.patternSearch(word[index+1:])
// 			}
// 			return flag
// 		} else {
// 			if this.next[v] == nil {
// 				return false
// 			}
// 			this = this.next[v]
// 		}
// 	}
// 	return this.isWord
// }

// /**
//  * Your WordDictionary object will be instantiated and called as such:
//  * obj := Constructor();
//  * obj.AddWord(word);
//  * param_2 := obj.Search(word);
//  */
