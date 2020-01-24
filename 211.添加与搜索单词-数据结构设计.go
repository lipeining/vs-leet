/*
 * @lc app=leetcode.cn id=211 lang=golang
 *
 * [211] 添加与搜索单词 - 数据结构设计
 */

// @lc code=start
type WordDictionary struct {
    words []string
}


/** Initialize your data structure here. */
func Constructor() WordDictionary {
    return WordDictionary{}
}


/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string)  {
    this.words = append(this.words, word)
}


/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
    for i:=0;i<len(this.words);i++ {
		if helper(this.words[i], word) {
			return true
		}
	}
	return false
}
func helper(word string, pattern string) bool {
	// var dfs func(word string, pattern string, path string)
	// dfs = func(word string, pattern string, path string) {
	// 	if path == word {
	// 		ans = true
	// 		return
	// 	}
	// 	if len(path) >= len(word) {
	// 		return
	// 	}
	// 	for i:=0;i< len(pattern);i++ {

	// 	}
	// }
	// dfs(word, pattern, "")
	if len(word) != len(pattern) {
		return false
	}
	for i:=0;i<len(word);i++ {
		if pattern[i] == '.' {
			continue
		} else if pattern[i] != word[i] {
			return false
		}
	}
	return true
}


/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
// @lc code=end

