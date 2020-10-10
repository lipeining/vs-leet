package lc

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

// // [["a","b"],["c","d"]]
// // ["ab","cb","ad","bd","ac","ca","da","bc","db","adcb","dabc","abb","acb"]

// // [["a","b"],["a","a"]]
// // ["aba","baa","bab","aaab","aaa","aaaa","aaba"]

// func findWords(board [][]byte, words []string) []string {
// 	root := Constructor()
// 	for _, word := range words {
// 		root.Insert(word)
// 	}
// 	n := len(board)
// 	m := len(board[0])
// 	visited := make([][]bool, n)
// 	for i := 0; i < n; i++ {
// 		visited[i] = make([]bool, m)
// 	}
// 	ansMap := make(map[string]bool)
// 	var dfs func(i, j int, path string)
// 	dfs = func(i, j int, path string) {
// 		if i < 0 || i >= n || j < 0 || j >= m {
// 			return
// 		}
// 		if visited[i][j] {
// 			return
// 		}
// 		if !root.StartsWith(path) {
// 			return
// 		}

// 		next := path + string(board[i][j])
// 		if root.Search(next) {
// 			ansMap[next] = true
// 			// return
// 		}
// 		visited[i][j] = true
// 		dfs(i-1, j, next)
// 		dfs(i+1, j, next)
// 		dfs(i, j-1, next)
// 		dfs(i, j+1, next)
// 		visited[i][j] = false
// 	}
// 	for i := 0; i < n; i++ {
// 		for j := 0; j < m; j++ {
// 			dfs(i, j, "")
// 		}
// 	}
// 	ans := make([]string, 0)
// 	for k, _ := range ansMap {
// 		ans = append(ans, k)
// 	}
// 	return ans
// }
