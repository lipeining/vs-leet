/*
 * @lc app=leetcode.cn id=1002 lang=golang
 *
 * [1002] 查找常用字符
 */

// @lc code=start
func commonChars(A []string) []string {
	count := [26]int{}
	for i:=0;i<len(count);i++ {
		count[i]=101
	}
	for i:=0;i<len(A);i++ {
		word:=[26]int{}
		for j:=0;j<len(A[i]);j++ {
			word[A[i][j]-'a']++
		}
		for k:=0;k<len(count);k++{
			if count[k] > word[k] {
				count[k]=word[k]
			}
		}
	}
	ans := make([]string, 0)
	for i:=0;i<len(count);i++ {
		if count[i]!=0 {
			for j:=0;j<count[i];j++ {
				ans = append(ans, string('a'+i))
			}			
		}
	}
	return ans
}
// @lc code=end

