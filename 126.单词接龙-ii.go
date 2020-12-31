/*
 * @lc app=leetcode.cn id=126 lang=golang
 *
 * [126] 单词接龙 II
 */
package main

import "fmt"

// func main() {
// 	findLadders("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"})
// 	findLadders("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"})
// 	findLadders("a", "c", []string{"a", "b", "c"})
// 	// 超出内存了
// 	// 	"qa"
// 	// ' +
// 	//   '"sq"
// 	// ' +
// 	//   '["si","go","se","cm","so","ph","mt","db","mb","sb","kr","ln","tm","le","av","sm","ar","ci","ca","br","ti","ba","to","ra","fa","yo","ow","sn","ya","cr","po","fe","ho","ma","re","or","rn","au","ur","rh","sr","tc","lt","lo","as","fr","nb","yb","if","pb","ge","th","pm","rb","sh","co","ga","li","ha","hz","no","bi","di","hi","qa","pi","os","uh","wm","an","me","mo","na","la","st","er","sc","ne","mn","mi","am","ex","pt","io","be","fm","ta","tb","ni","mr","pa","he","lr","sq","ye"]
// }

// @lc code=start
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	ans := make([][]string, 0)
	pre := make(map[string][]string)
	n := len(wordList)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if getDiffCount(wordList[i], wordList[j]) == 1 {
				pre[wordList[i]] = append(pre[wordList[i]], wordList[j])
				pre[wordList[j]] = append(pre[wordList[j]], wordList[i])
			}
		}
	}
	if _, ok := pre[beginWord]; !ok {
		for i := 0; i < n; i++ {
			if getDiffCount(beginWord, wordList[i]) == 1 {
				pre[beginWord] = append(pre[beginWord], wordList[i])
			}
		}
	}
	// fmt.Println(pre)
	if _, ok := pre[endWord]; !ok {
		return ans
	}
	// bfs
	type pair struct {
		word string
		path []string
	}
	queue := make([]pair, 0)
	queue = append(queue, pair{word: beginWord, path: []string{beginWord}})
	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			item := queue[i]
			if item.word == endWord {
				// fmt.Println("path", item)
				if len(ans) != 0 {
					if len(item.path) < len(ans[0]) {
						ans = make([][]string, 0)
						ans = append(ans, item.path)
					} else if len(item.path) == len(ans[0]) {
						ans = append(ans, item.path)
					}
				} else {
					ans = append(ans, item.path)
				}
				continue
			}
			for _, next := range pre[item.word] {
				if !includes(item.path, next) {
					nextPath := make([]string, len(item.path))
					copy(nextPath, item.path)
					nextPath = append(nextPath, next)
					queue = append(queue, pair{word: next, path: nextPath})
				}
			}
		}
		queue = queue[size:]
	}
	fmt.Println("ans", ans)
	return ans
}
func includes(list []string, word string) bool {
	for _, w := range list {
		if w == word {
			return true
		}
	}
	return false
}
func getDiffCount(a, b string) int {
	diff := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diff++
		}
	}
	return diff
}

// @lc code=end
