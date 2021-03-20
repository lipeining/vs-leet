/*
 * @lc app=leetcode.cn id=388 lang=golang
 *
 * [388] 文件的最长绝对路径
 */
package main

import (
	"strings"
)

// func main() {
// 	lengthLongestPath("dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext")
// 	lengthLongestPath("a")
// 	lengthLongestPath("file1.txt\nfile2.txt\nlongfile.txt")
// }

// @lc code=start
func lengthLongestPath(input string) int {
	list := strings.Split(input, "\n")
	ans := 0
	levels := make(map[int]int)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 0; i < len(list); i++ {
		now := list[i]
		start := strings.LastIndex(now, "\t")
		isFile := strings.Index(now, ".") != -1
		level := start + 1
		// fmt.Println("now", now, len(now), now[start+1:], start, isFile, level, ans)
		if start == -1 {
			// 当前是 level 0，重置 prev
			levels = make(map[int]int)
			if isFile {
				ans = max(ans, len(now))
			} else {
				levels[0] = len(now)
			}
		} else {
			if isFile {
				prev := levels[level-1]
				ans = max(ans, prev+len(now[start+1:])+1)
			} else {
				cur := len(now[start+1:]) + 1
				levels[level] = cur + levels[level-1]
			}
		}
		// fmt.Println(levels, ans)
	}
	// fmt.Println("ansssss", ans)
	return ans
}

// @lc code=end
