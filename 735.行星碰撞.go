/*
 * @lc app=leetcode.cn id=735 lang=golang
 *
 * [735] 行星碰撞
 */

// @lc code=start
func asteroidCollision(asteroids []int) []int {
	// 当 flag 不同时，需要处理 stack
	stack := make([]int, 0)
	if len(asteroids) == 0 {
		return stack
	}
	// 找到第一个往右星球
	r:=0
	for r<len(asteroids) {
		if asteroids[r] > 0 {
			break
		}
		stack = append(stack, asteroids[r])
		r++
	}
	for i:=r;i<len(asteroids);i++ {
		cur := asteroids[i] > 0
		if cur {
			// 方向相同
			stack = append(stack, asteroids[i])
		} else {
			if len(stack) == 0 {
				stack = append(stack, asteroids[i])
				continue
			}
			neg := asteroids[i]
			for {
				if len(stack) == 0 {
					break
				}
				top := stack[len(stack)-1]
				if top < 0 {
					break
				}
				sum := top + neg
				if sum > 0 {
					neg = top
					break
				} else if sum < 0 {
					stack = stack[:len(stack)-1]
				} else {
					// 刚好抵消
					stack = stack[:len(stack)-1]
					neg = 0
					break
				}
			}
			if neg < 0 {
				stack = append(stack, neg)
			}
		}
	}
	return stack
}
// @lc code=end

