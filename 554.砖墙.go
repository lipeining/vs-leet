/*
 * @lc app=leetcode.cn id=554 lang=golang
 *
 * [554] 砖墙
 */

// @lc code=start
func leastBricks(wall [][]int) int {
	// 计算边缘的最大值,将长度减去这个边缘的数量，就可以得到结果
	n := len(wall)
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		start := 0
		for j := 0; j < len(wall[i]); j++ {
			if j == len(wall[i])-1 {
				continue
			}
			start += wall[i][j]
			m[start]++
		}
	}
	most := 0
	for _, v := range m {
		if v > most {
			most = v
		}
	}
	return n - most
}

// @lc code=end

