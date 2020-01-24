/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
func myAtoi(str string) int {
	neg := false
	n:=""
	i:=0;
	for i<len(str) {
		if str[i] != ' ' {
			break
		}
		i++
	}
	if i==len(str) {
		return 0
	}
	if str[i] == '-' {
		neg = true
		i++
	} else if str[i] == '+' {
		i++	
	} else if !unicode.IsDigit(rune(str[i])) {
		return 0
	}

	for i<len(str) {
		if !unicode.IsDigit(rune(str[i])) {
			break
		}
		n+=string(str[i])
		i++
	}
	fmt.Println(n, i, neg, str)
	num := atoi(n, neg)
	fmt.Print(num)
	return num
}
func atoi(n string, neg bool) int {
	// max := 1<<31 -1
	// min := -(1<<31)
	// var max int64 = 1<<31 -1
	// var min int64 = -(1<<31)
	// var max64 int64 = 1<<63 - 1
	// var min64 int64 = -(1<<63)
	length := len(n)
	var multi int = 1
	var ans int  = 0
	for i:=length-1;i>=0;i-- {
		tmp := int(n[i]-'0')
		ans += multi * tmp
		multi *= 10
	}
	if neg  && ans * -1 < math.MinInt32{
		return math.MinInt32
	}
	if !neg && ans > math.MaxInt32 {
		return math.MaxInt32
	}
	fmt.Print(ans)
	if neg {
		ans = -ans
	}
	return ans
	// if ans > max64 {
	// 	ans = max64
	// }
	// if ans <
	// if ans > max {
	// 	return int(max)
	// }
	// if ans < min {
	// 	return int(min)
	// }
	// return int(ans)
}
// @lc code=end

