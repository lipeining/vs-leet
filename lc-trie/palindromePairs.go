package lc

// type Trie struct {
// 	next   map[rune]*Trie
// 	index  int
// 	isWord bool
// }

// /** Initialize your data structure here. */
// func Constructor() Trie {
// 	root := new(Trie)
// 	root.next = make(map[rune]*Trie)
// 	root.isWord = false
// 	root.index = -1
// 	return *root
// }

// /** Inserts a word into the trie. */
// func (this *Trie) Insert(word string, index int) {
// 	for _, v := range word {
// 		if this.next[v] == nil {
// 			node := new(Trie)
// 			//子节点数量为26
// 			node.next = make(map[rune]*Trie)
// 			//初始化节点单词标志为假
// 			node.isWord = false
// 			node.index = -1
// 			this.next[v] = node
// 		}
// 		this = this.next[v]
// 		// this.count++ 统计每个单词前缀出现的次数（也包括统计每个单词出现的次数）
// 	}
// 	this.index = index
// 	this.isWord = true
// }

// /** Returns if the word is in the trie. */
// func (this *Trie) Search(word string) bool {
// 	for _, v := range word {
// 		if this.next[v] == nil {
// 			return -1
// 		}
// 		this = this.next[v]
// 	}
// 	if !this.isWord {
// 		return -1
// 	}
// 	return this.index
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

// /**
//  * 给定一组 互不相同 的单词， 找出所有不同 的索引对(i, j)，使得列表中的两个单词， words[i] + words[j] ，可拼接成回文串。
//  *
//  * 示例 1：
//  *
//  * 输入：["abcd","dcba","lls","s","sssll"]
//  * 输出：[[0,1],[1,0],[3,2],[2,4]]
//  * 解释：可拼接成的回文串为 ["dcbaabcd","abcddcba","slls","llssssll"]
//  * 示例 2：
//  *
//  * 输入：["bat","tab","cat"]
//  * 输出：[[0,1],[1,0]]
//  * 解释：可拼接成的回文串为 ["battab","tabbat"]
//  *
//  * 作者：力扣 (LeetCode)
//  * 链接：https://leetcode-cn.com/leetbook/read/trie/x75mhj/
//  * 来源：力扣（LeetCode）
//  * 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
//  *
//  */
// func palindromePairs(words []string) [][]int {
// 	ans := make([][]int, 0)
// 	n := len(words)
// 	root := Constructor()
// 	for i := 0; i < n; i++ {
// 		root.Insert(words[i])
// 	}
// 	for i := 0; i < n; i++ {
// 		for j := i + 1; j < n; j++ {
// 			rLeft := reverse(words[i])
// 			rRight := reverse(words[j])
// 			fLeft := root.Search(rLeft)
// 			fRight := root.Search(rRight)
// 			if fLeft && fRight {
// 				ans = append(ans, []int{i, j})
// 			}
// 		}
// 	}
// 	return ans
// }

// // func palindromePairs(words []string) [][]int {
// // 	ans := make([][]int, 0)
// // 	n := len(words)
// // 	for i := 0; i < n; i++ {
// // 		for j := 0; j < n; j++ {
// // 			if i != j {
// // 				sentence := words[i] + words[j]
// // 				if check(sentence) {
// // 					ans = append(ans, []int{i, j})
// // 				}
// // 			}
// // 		}
// // 	}
// // 	return ans
// // }
// // func check(word string) bool {
// // 	n := len(word)
// // 	for i := 0; i < n/2; i++ {
// // 		if word[i] != word[n-i-1] {
// // 			return false
// // 		}
// // 	}
// // 	return true
// // }
// func reverse(word string) string {
// 	ans := ""
// 	for i := len(word) - 1; i >= 0; i-- {
// 		ans += string(word[i])
// 	}
// 	return ans
// }
