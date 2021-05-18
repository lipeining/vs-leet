/*
 * @lc app=leetcode.cn id=638 lang=golang
 *
 * [638] 大礼包
 */

// @lc code=start
// func shoppingOffers(price []int, special [][]int, needs []int) int {

// }
func shoppingOffers(price []int, special [][]int, needs []int) int {
	var dfs func(price []int, special [][]int, needs []int) int
	dfs = func(price []int, special [][]int, needs []int) int {
		res := dot(price, needs)
		for i := 0; i < len(special); i++ {
			nextNeeds := make([]int, len(needs))
			copy(nextNeeds, needs)
			j := 0
			for j < len(nextNeeds) {
				diff := nextNeeds[j] - special[i][j]
				if diff < 0 {
					break
				}
				nextNeeds[j] = diff
				j++
			}
			if j == len(nextNeeds) {
				nextSum := dfs(price, special, nextNeeds)
				if res > nextSum+special[i][j] {
					res = nextSum + special[i][j]
				}
			}
		}
		return res
	}
	return dfs(price, special, needs)
}
func dot(price []int, needs []int) int {
	sum := 0
	for i := 0; i < len(price); i++ {
		sum += price[i] * needs[i]
	}
	return sum
}

// @lc code=end
// func shoppingOffers(price []int, special [][]int, needs []int) int {
// 	num := len(price)
//     for i:=0;i<num;i++ {
// 		t := make([]int, num + 1)
// 		t[i] = 1
// 		t[num] = price[i]
// 		special = append(special, t)
// 	}
// 	fmt.Println(special)
// 	ans := math.MaxInt32
// 	// 把只买一个的也当做大礼包吧
// 	var dfs func(special [][]int, needs []int, path []int, sum int)
// 	dfs = func(special [][]int, needs []int, path []int, sum int) {
// 		flag := true
// 		// needs[i] < path[i] 这里超出之后是否是剪枝了
// 		for i:=0;i<len(needs);i++ {
// 			if needs[i] < path[i] {
// 				return
// 			} else if needs[i] > path[i] {
// 				flag = false
// 				break
// 			}
// 		}
// 		if flag {
// 			if sum < ans {
// 				ans = sum
// 			}
// 			return
// 		}
// 		// 还需要递归查找答案
// 		for i:=0;i<len(special);i++ {
// 			nextSum := sum + special[i][len(special[i])-1]
// 			nextPath := make([]int, len(path))
// 			copy(nextPath, path)
// 			for j:=0;j<len(special[i])-1;j++ {
// 				nextPath[j]+=special[i][j]
// 			}
// 			fmt.Println(nextPath, nextSum)
// 			dfs(special, needs, nextPath, nextSum)
// 		}
// 	}
// 	path := make([]int, len(needs))
// 	dfs(special, needs, path, 0)
// 	return ans
// }

