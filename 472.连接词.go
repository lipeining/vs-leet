/*
 * @lc app=leetcode.cn id=472 lang=golang
 *
 * [472] 连接词
 */
package main

// func main() {
// 	findAllConcatenatedWordsInADict([]string{"cat", "cats", "catsdogcats", "dog", "dogcatsdog", "hippopotamuses", "rat", "ratcatdogcat"})
// }

// @lc code=start
func findAllConcatenatedWordsInADict(words []string) []string {
	n := len(words)
	set := make(map[string]int)
	for i := 0; i < n; i++ {
		set[words[i]] = i
	}
	ans := make([]string, 0)
	var dfs func(word string, index, cnt int) bool
	dfs = func(word string, index, cnt int) bool {
		if index == len(word) && cnt >= 2 {
			return true
		}
		if index >= len(word) {
			return false
		}
		for j := index; j <= len(word); j++ {
			if _, ok := set[word[index:j]]; ok {
				if dfs(word, j, cnt+1) {
					return true
				}
			}
		}
		return false
	}
	for i := 0; i < n; i++ {
		if len(words[i]) == 0 {
			continue
		}
		if dfs(words[i], 0, 0) {
			ans = append(ans, words[i])
		}
	}
	// fmt.Println("ans", ans)
	return ans
}

// @lc code=end
