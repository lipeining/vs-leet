package main

import "fmt"

func main() {

}
func countPalindromicSubsequence(s string) int {
	// x k x
	// x 有 26 个字母
	n := len(s)
	record := make([][]int, 26)
	for i := 0; i < 26; i++ {
		record[i] = make([]int, 2)
		record[i][0] = -1
	}
	for i := 0; i < n; i++ {
		char := s[i]
		index := int(char - 'a')
		if record[index][0] == -1 {
			record[index][0] = i
		}
		record[index][1] = i
	}
	ans := 0
	for index := 0; index < 26; index++ {
		first, last := record[index][0], record[index][1]
		if last-first < 2 {
			continue
		}
		cnt := make(map[byte]bool)
		for i := first + 1; i < last; i++ {
			cnt[s[i]] = true
		}
		ans += len(cnt)
		fmt.Println("for index", index, len(cnt))
	}
	return ans
}
func getConcatenation(nums []int) []int {
	n := len(nums)
	ans := make([]int, n*2)
	for i := 0; i < n; i++ {
		ans[i] = nums[i]
		ans[i+n] = nums[i]
	}
	return ans
}
