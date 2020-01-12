/*
 * @lc app=leetcode.cn id=443 lang=golang
 *
 * [443] 压缩字符串
 */

// @lc code=start
func compress(chars []byte) int {
	res := 0
	slow := 0
	c := 0
	// 参考标准答案的话，是可以使用 fast + 1 = len 或者 fast != slow 合并代码
	// 对于那个 helper 可以看标准的来做
	for fast:= 0;fast<len(chars);fast++ {
		if chars[fast] != chars[slow] {
			if c == 1 {
				chars[res] = chars[slow]
				res = res + 1
				slow = fast
				c = 1
			} else {
				num := helper(c)
				chars[res] = chars[slow]
				for j:= 0; j<len(num); j++ {
					chars[res+1+j] = num[j]
				}
				fmt.Println(res, slow, fast, num)
				res = res + 1 + len(num)
				slow = fast
				c = 1
			}
		} else {
			c++
		}
	}
	if c == 0 {
		return res
	} else if c == 1 {
		chars[res] = chars[slow]
		return res + 1
	} else {
		num := helper(c)
		chars[res] = chars[slow]
		for j:= 0; j<len(num); j++ {
			chars[res+1+j] = num[j]
		}
		fmt.Println(res)
		// fmt.Println(res, slow, fast, num)
		res = res + 1 + len(num)
		return res
	}
	// fmt.Println(res)
	// return res
}
func helper(c int) []byte {
	s := []byte("")
	for c != 0 {
		i := c % 10
		c = c / 10
		s = append([]byte(strconv.Itoa(i)), s...)
	}
	return s
}
// @lc code=end

