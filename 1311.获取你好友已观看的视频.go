import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=1311 lang=golang
 *
 * [1311] 获取你好友已观看的视频
 */

// @lc code=start
func watchedVideosByFriends(watchedVideos [][]string, friends [][]int, id int, level int) []string {
	m := make(map[string]int)
	visited := make([]bool, len(watchedVideos))
	visited[id] = true
	dfs(watchedVideos, friends, id, level, 0, m, visited)
	fmt.Println(m)
	ans := make([]string, 0)
	for k, _ := range m {
		ans = append(ans, k)
	}
	sort.Slice(ans, func(i, j int) bool {
		diff := m[ans[i]] - m[ans[j]]
		if diff != 0 {
			return diff < 0
		}
		return ans[i] < ans[j]
	})
	return ans
}
func dfs(watchedVideos [][]string, friends [][]int, id int, level int, nowLevel int, m map[string]int, visited []bool) {
	if nowLevel == level {
		watched := watchedVideos[id]
		for _, movie := range watched {
			m[movie]++
		}
		return
	}
	myFriends := friends[id]
	for _, user := range myFriends {
		if visited[user] {
			continue
		}
		visited[user] = true
		dfs(watchedVideos, friends, user, level, nowLevel+1, m, visited)
	}
}

// @lc code=end
