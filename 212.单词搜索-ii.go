/*
 * @lc app=leetcode.cn id=212 lang=golang
 *
 * [212] 单词搜索 II
 */
package main

import "fmt"

func main() {
	// 	输入：board = [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]], words = ["oath","pea","eat","rain"]
	// 输出：["eat","oath"]
	// 输入：board = [["a","b"],["c","d"]], words = ["abcb"]
	// 输出：[]
	findWords([][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}},
		[]string{"oath", "pea", "eat", "rain"})
	findWords([][]byte{{'a', 'b'}, {'c', 'd'}}, []string{"abcb"})
	findWords([][]byte{{'a', 'a'}}, []string{"a"})
}

// @lc code=start
func findWords(board [][]byte, words []string) []string {
	root := Constructor()
	for _, word := range words {
		root.Insert(word)
	}
	n := len(board)
	m := len(board[0])
	ans := make([]string, 0)
	directions := [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	var dfs func(i, j int, parent *Trie)
	dfs = func(i, j int, parent *Trie) {
		letter := board[i][j]
		cur := parent.next[rune(letter)]
		if cur.isWord {
			ans = append(ans, cur.word)
			cur.word = ""
			cur.isWord = false
		}
		board[i][j] = '#'
		for _, direction := range directions {
			x := i + direction[0]
			y := j + direction[1]
			if x >= 0 && x < n && y >= 0 && y < m {
				next := board[x][y]
				if _, ok := cur.next[rune(next)]; ok {
					dfs(x, y, cur)
				}
			}
		}
		board[i][j] = letter
		fmt.Println(cur)
		if len(cur.next) == 0 {
			delete(parent.next, rune(letter))
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if _, ok := root.next[rune(board[i][j])]; ok {
				dfs(i, j, root)
			}
		}
	}
	fmt.Println("ans", ans)
	return ans
}

const MAXCAP = 26 // a-z

type Trie struct {
	next   map[rune]*Trie
	isWord bool
	word   string
}

/** Initialize your data structure here. */
func Constructor() *Trie {
	root := new(Trie)
	root.next = make(map[rune]*Trie, MAXCAP)
	root.isWord = false
	return root
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	for _, v := range word {
		if this.next[v] == nil {
			node := new(Trie)
			//子节点数量为26
			node.next = make(map[rune]*Trie, MAXCAP)
			//初始化节点单词标志为假
			node.isWord = false
			this.next[v] = node
		}
		this = this.next[v]
		// this.count++ 统计每个单词前缀出现的次数（也包括统计每个单词出现的次数）
	}
	this.isWord = true
	this.word = word
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for _, v := range word {
		if this.next[v] == nil {
			return false
		}
		this = this.next[v]
	}
	return this.isWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for _, v := range prefix {
		if this.next[v] == nil {
			return false
		}
		this = this.next[v]
	}
	return true
}

// @lc code=end
