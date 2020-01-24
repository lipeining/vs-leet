/*
 * @lc app=leetcode.cn id=1079 lang=golang
 *
 * [1079] 活字印刷
 */

// @lc code=start
func numTilePossibilities(tiles string) int {
	arr := strings.Split(tiles, "")
	sort.Strings(arr)
	fmt.Println(arr)
	ans := 0
	var dfs func(arr []string, index int, k int, now int, used []bool)
	dfs = func(arr []string, index int, k int, now int, used []bool) {
		if k == now {
			ans++
			return
		}
		for i:=index;i<len(arr);i++ {
			if i>index && arr[i-1] == arr[i] && !used[i-1] {
				continue
			}
			used[i] = true
			dfs(arr, index+1, k, now+1, used)
			used[i] = false
		}
	}
	used := make([]bool, len(tiles))
	dfs(arr, 0, 1, 0, used)
	// used := make([]bool, len(tiles))
	// for i:=1;i<=len(tiles);i++ {
	// 	dfs(arr, 0, i, 0, used)
	// }
	return ans
}
// @lc code=end

