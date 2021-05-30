/*
 * @lc app=leetcode.cn id=781 lang=golang
 *
 * [781] 森林中的兔子
 */

// @lc code=start
func numRabbits(answers []int) int {
	counter := make(map[int]int)
	for i := 0; i < len(answers); i++ {
		counter[answers[i]]++
	}
	// fmt.Println(counter)
	ans := 0
	for k, v := range counter {
		if k == 0 {
			ans += v
			continue
		}
		should := k + 1
		// now := v
		// fmt.Println(should, k, v)
		if should < v {
			time := v / should
			if v%should == 0 {
				ans += time * should
			} else {
				ans += (time + 1) * should
			}
		} else {
			ans += should
		}
	}
	return ans
}

// @lc code=end

