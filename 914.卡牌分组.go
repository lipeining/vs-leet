/*
 * @lc app=leetcode.cn id=914 lang=golang
 *
 * [914] 卡牌分组
 */

// @lc code=start
func hasGroupsSizeX(deck []int) bool {
	cnt := make(map[int]int)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	size := len(deck)
	for _, num := range deck {
		cnt[num]++
	}
	for _, v := range cnt {
		size = min(size, v)
	}
	// fmt.Println(cnt, size)
	for x := 2; x <= size; x++ {
		flag := true
		for _, v := range cnt {
			if v%x != 0 {
				flag = false
				break
			}
		}
		if flag {
			return true
		}
	}
	return false
}

// @lc code=end

