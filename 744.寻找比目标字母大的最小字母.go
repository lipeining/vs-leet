/*
 * @lc app=leetcode.cn id=744 lang=golang
 *
 * [744] 寻找比目标字母大的最小字母
 */

// @lc code=start
func nextGreatestLetter(letters []byte, target byte) byte {
	length:=len(letters)
	// double := make([]byte, length * 2)
	// for i:=0;i<length;i++ {
	// 	double[i] = double[i+length] = letters
	// }
	for i:=0; i <length;i++ {
		if letters[i] > target {
			return letters[i]
		}
	}
	return letters[0]
}
// @lc code=end

