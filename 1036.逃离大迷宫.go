/*
 * @lc app=leetcode.cn id=1036 lang=golang
 *
 * [1036] 逃离大迷宫
 */

// @lc code=start
func isEscapePossible(blocked [][]int, source []int, target []int) bool {
	blocks := make(map[string]bool)
	toKey := func(a, b int)string{
		return strconv.Itoa(a) + ":" + strconv.Itoa(b)
	}
	for _, block := range blocked {	
		blocks[toKey(block[0], block[1])] = true
	}
	dirs := [][]int{
		{1,0},
		{-1,0},
		{0,1},
		{0,-1},
	}
	limit := int(1e6)
	var bfs func(s, t []int) bool
	bfs = func(s, t []int) bool {
		seen := make(map[string]bool)
		queue := make([][]int, 0)
		queue = append(queue, s)
		seen[toKey(s[0], s[1])] = true
		for len(queue) != 0 {
			p := queue[0]
			queue = queue[1:]
			for _, dir := range dirs {
				nx := p[0] + dir[0]
				ny := p[1] + dir[1]
				k := toKey(nx, ny)
				if nx < 0 || nx >= limit || ny < 0 || ny >= limit {
					continue
				}
				if seen[k] || blocks[k] {
					continue
				}
				if nx == t[0] && ny == t[1] {
					return true
				}
				seen[k] = true
				queue = append(queue, []int{nx, ny})
			}
			if len(seen) >= 20000 {
				return true
			}
		}
		return false
	}
	return bfs(source, target) && bfs(target, source)
}
// @lc code=end

