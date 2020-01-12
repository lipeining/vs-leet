/*
 * @lc app=leetcode.cn id=415 lang=golang
 *
 * [415] 字符串相加
 */
import "strconv"
// @lc code=start
func addStrings(num1 string, num2 string) string {
	flag := 0
	res := ""
	i:= len(num1)-1
	j:= len(num2)-1
	for i >= 0 || j >= 0 {
		l := 0
		if i>=0 {
			ll,ok := strconv.Atoi(string(num1[i]))
			if ok == nil {
				l = ll
			}
		}
		r := 0
		if j >= 0 {
			rr,ok := strconv.Atoi(string(num2[j]))
			if ok == nil {
				r = rr
			}
		}
		t := flag + l + r
		// fmt.Println(flag, t, res, i, j)
		if t >= 10 {
			flag = 1
			t = t - 10
		} else {
			flag = 0
		}
		res = strconv.Itoa(t) + res

		i=i-1
		j=j-1
	}
	if flag == 1 {
		return "1" + res
	}
	return res
}
// @lc code=end

