/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {
    if(x<0){
		return false;
	}
	res:=0;
	y:=x;
	for(x>0){
		m := x % 10;
		x = x / 10;
		res= res*10 + m;
	}
	return res == y;
}
// @lc code=end

