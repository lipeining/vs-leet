/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */
// import "math"
import "fmt"
// @lc code=start
func reverse(x int) int {
	neg := x < 0;
	res := 0;
	if(neg){
		x = -x;
	}
	for(x>0){
		m := x % 10;
		x = x / 10;
		res= res*10 + m;
	}
	// fmt.Print(res)
	max := 2147483648 - 1;
	// fmt.Println(2 >> 31, 2 >> 31 - 1, max)
	// fmt.Println(2 << 31, 2 << 31 - 1, max)
	// fmt.Println(2 << 30, 2 << 30 - 1, max)
	// fmt.Println(1 << 30, 1 << 30 - 1, max)
	// fmt.Println(1 << 31, 1 << 31 - 1, max) // correct
	// max := 2**31 - 1;
	if(res > max) {
		return 0;
	}
	if(neg){
		return -res;
	}
	return res;
	// return (res > 2^31-1) ? 0 : res;
}
// @lc code=end

