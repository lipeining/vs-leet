/*
 * @lc app=leetcode.cn id=29 lang=golang
 *
 * [29] 两数相除
 */

// @lc code=start
func divide(dividend int, divisor int) int {
	x,y:=dividend, divisor
	isNeg := false
	if (x<0&&y>0) || (x>0&&y<0) {
		isNeg = true
	}
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	mul := func(a, k int)int{
		ans := 0
		for k > 0 {
			if k&1 == 1 {
				ans+=a
			}
			k>>=1
			a+=a
		}
		return ans
	}
	l := 0
	r := x
	for l<r {
		mid := l+r+1>>1
		if mul(mid, y) <= x {
			l=mid
		} else {
			r = mid-1
		}
	}
	ans := l
	if isNeg {
		ans = -ans
	}
	if ans > math.MaxInt32 || x <math.MinInt32 {
		return math.MaxInt32
	}
	return ans
}
// @lc code=end

