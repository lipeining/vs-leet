/*
 * @lc app=leetcode.cn id=371 lang=golang
 *
 * [371] 两整数之和
 */

// @lc code=start
func getSum(a int, b int) int {
// 	while(b != 0){
// 		int temp = a ^ b;
// 		b = (a & b) << 1;
// 		a = temp;
// 	}
// 	return a;

// 作者：phoenixfei
// 链接：https://leetcode-cn.com/problems/sum-of-two-integers/solution/li-yong-wei-cao-zuo-shi-xian-liang-shu-qiu-he-by-p/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
for b != 0 {
	tmp := a ^ b
	b = (a&b)<<1
	a = tmp
}
return a
}
// @lc code=end

