package lc

// const MAXCAP = 26 // a-z

// type Trie struct {
// 	next   map[rune]*Trie
// 	isWord bool
// }

// /** Initialize your data structure here. */
// func Constructor() Trie {
// 	root := new(Trie)
// 	root.next = make(map[rune]*Trie)
// 	root.isWord = false
// 	return *root
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

// func replaceWords(dictionary []string, sentence string) string {
// 	root := Constructor()
// 	for _, dict := range dictionary {
// 		root.Insert(dict)
// 	}
// 	ans := ""
// 	last := 0
// 	for i := 0; i <= len(sentence); i++ {
// 		if i == len(sentence) || sentence[i] == ' ' {
// 			word := sentence[last:i]
// 			j := 0
// 			for j < len(word) {
// 				if root.Search(word[0:j]) {
// 					break
// 				}
// 				j++
// 			}
// 			ans += word[0:j]
// 			if i != len(sentence) {
// 				ans += " "
// 			}
// 			last = i + 1
// 		}
// 	}
// 	return ans
// }
