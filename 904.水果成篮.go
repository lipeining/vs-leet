import "fmt"

/*
 * @lc app=leetcode.cn id=904 lang=golang
 *
 * [904] 水果成篮
 */

// @lc code=start
func totalFruit(tree []int) int {
	left, right := 0, 0
	m := make(map[int]int)
	ans := 0
	for right < len(tree) {
		m[tree[right]]++
		for len(m) >= 3 {
			ori, _ := m[tree[left]]
			// fmt.Println(m, left, ori)
			if ori == 1 {
				delete(m, tree[left])
			} else {
				m[tree[left]] = ori - 1
			}		
			left++
		}
		ans = max(ans, right-left+1)
		right++
	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func totalFruit(tree []int) int {
// 	ans := 0
// 	for index := 0; index < len(tree); index++ {
// 		tmp := helper(tree, index)
// 		if ans < tmp {
// 			ans = tmp
// 		}
// 	}
// 	return ans
// }
func helper(tree []int, index int) int {
	m := make(map[int]int)
	for i := index; i < len(tree); i++ {
		length := len(m)
		_, ok := m[tree[i]]
		if length == 2 && !ok {
			return sumM(m)
		} else {
			m[tree[i]]++
		}
	}
	return sumM(m)
}
func sumM(m map[int]int) int {
	sum := 0
	for _, v := range m {
		sum += v
	}
	return sum
}

// @lc code=end
