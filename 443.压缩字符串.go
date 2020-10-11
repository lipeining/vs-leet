/*
 * @lc app=leetcode.cn id=443 lang=golang
 *
 * [443] 压缩字符串
 */

// @lc code=start
func compress(chars []byte) int {
	slow,fast := 0,0
	n := len(chars)
	if n <= 1 {
			return n
	}
	for fast < n {
			count := 1
			i:=fast
			for i<n-1 && chars[i] == chars[i+1] {
					i++
					count++
			}
			chars[slow] = chars[fast]
			slow++
			if count >= 2 {
					num := trans(count)
					// 计算得到的值
					for j:=0;j<len(num);j++ {
							chars[slow] = num[j]
							slow++
					}
			}
			fast = i+1
	}
	return slow
	// 处理剩下的字符串
}
func trans(count int)[]byte {
	m := []byte{'0','1','2','3','4','5','6','7','8','9'}
	ans := make([]byte, 0)
	for count != 0 {
			t := count%10
			ans = append([]byte{m[t]}, ans...)
			count /=10
	}
	fmt.Println(ans)
	return ans
}
// @lc code=end

