/*
 * @lc app=leetcode.cn id=483 lang=golang
 *
 * [483] 最小好进制
 */

// @lc code=start
func smallestGoodBase(n string) string {
	num, _ := strconv.Atoi(n)
	for m := int(math.Log2(float64(num))); m >= 1; m-- {
		l, r := 2, int(math.Pow(float64(num), 1/float64(m)))+1
		for l < r {
			k := (l + r) / 2
			sum := 1
			for j := 0; j < m; j++ {
				sum = sum*k + 1
			}
			if sum == num {
				return strconv.Itoa(k)
			} else if sum < num {
				l = k + 1
			} else {
				r = k
			}
		}
	}
	return strconv.Itoa(num - 1)
}

// @lc code=end

