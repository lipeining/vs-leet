/*
 * @lc app=leetcode.cn id=1207 lang=golang
 *
 * [1207] 独一无二的出现次数
 */

// @lc code=start
func uniqueOccurrences(arr []int) bool {
	counter := make(map[int]int)
	for i:=0;i<len(arr);i++ {
		n,ok := counter[arr[i]]
		if ok {
			counter[arr[i]] = n + 1
		} else {
			counter[arr[i]] = 1
		}
	}
	bit := make([]bool, len(arr) + 1)
	fmt.Println(bit, counter)
	for _,v := range counter {
		if bit[v] {
			return false
		}
		bit[v] = true
	}
	return true
}
// @lc code=end

