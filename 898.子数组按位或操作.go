/*
 * @lc app=leetcode.cn id=898 lang=golang
 *
 * [898] 子数组按位或操作
 */
package main

// import "fmt"

// func main() {
// 	subarrayBitwiseORs([]int{0})
// 	subarrayBitwiseORs([]int{1, 1, 2})
// 	subarrayBitwiseORs([]int{1, 2, 4})
// }

// @lc code=start
func subarrayBitwiseORs(arr []int) int {
	ans := make(map[int]bool)
	cur := make(map[int]bool)
	cur[0] = true
	for _, x := range arr {
		cur2 := make(map[int]bool)
		for y := range cur {
			cur2[x|y] = true
		}
		cur2[x] = true
		cur = cur2
		for k := range cur {
			ans[k] = true
		}
	}
	// fmt.Println(ans)
	return len(ans)
	// n := len(arr)
	// m := make(map[int]bool)
	// for i := 0; i < n; i++ {
	// 	cur := arr[i]
	// 	m[cur] = true
	// 	for j := i + 1; j < n; j++ {
	// 		cur |= arr[j]
	// 		m[cur] = true
	// 	}
	// }
	// // fmt.Println(m)
	// return len(m)
}

// @lc code=end
