package main

import "fmt"

func main() {
	func2()
	// func3()
}
func func2() {
	// 输入：n = 2, languages = [[1],[2],[1,2]], friendships = [[1,2],[1,3],[2,3]]
	// 输出：1
	// 解释：你可以选择教用户 1 第二门语言，也可以选择教用户 2 第一门语言。
	// 示例 2：

	// 输入：n = 3, languages = [[2],[1,3],[1,2],[3]], friendships = [[1,4],[1,2],[3,4],[2,3]]
	// 输出：2
	// 解释：教用户 1 和用户 2 第三门语言，需要教 2 名用户。
	minimumTeachings(2, [][]int{{1}, {2}, {1, 2}}, [][]int{{1, 2}, {1, 3}, {2, 3}})
	minimumTeachings(3, [][]int{{2}, {1, 3}, {1, 2}, {3}}, [][]int{{1, 4}, {1, 2}, {3, 4}, {2, 3}})
	// 	11
	// [[3,11,5,10,1,4,9,7,2,8,6],[5,10,6,4,8,7],[6,11,7,9],[11,10,4],[6,2,8,4,3],[9,2,8,4,6,1,5,7,3,10],[7,5,11,1,3,4],[3,4,11,10,6,2,1,7,5,8,9],[8,6,10,2,3,1,11,5],[5,11,6,4,2]]
	// [[7,9],[3,7],[3,4],[2,9],[1,8],[5,9],[8,9],[6,9],[3,5],[4,5],[4,9],[3,6],[1,7],[1,3],[2,8],[2,6],[5,7],[4,6],[5,8],[5,6],[2,7],[4,8],[3,8],[6,8],[2,5],[1,4],[1,9],[1,6],[6,7]]
}
func minimumTeachings(n int, languages [][]int, friendships [][]int) int {
	luMap := make(map[int]map[int]bool)
	nuMap := make(map[int]map[int]bool)
	for i := 1; i <= n; i++ {
		luMap[i] = make(map[int]bool)
		nuMap[i] = make(map[int]bool)
	}
	m := len(languages)
	for i := 0; i < m; i++ {
		for j := 0; j < len(languages[i]); j++ {
			learned := languages[i][j]
			luMap[learned][i+1] = true
		}
	}
	ans := m
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i := 0; i < len(friendships); i++ {
		left := friendships[i][0]
		right := friendships[i][1]
		// 如果 left right 没有共同语言，那么对每一种语言都加一
		pass := false
		for j := 1; j <= n; j++ {
			if luMap[j][left] && luMap[j][right] {
				pass = true
				break
			}
		}
		if pass {
			continue
		}
		for j := 1; j <= n; j++ {
			if !luMap[j][left] {
				nuMap[j][left] = true
			}
			if !luMap[j][right] {
				nuMap[j][right] = true
			}
		}
	}
	for l, v := range nuMap {
		need := len(v)
		fmt.Println("language ", l, "need", need)
		ans = min(ans, need)
	}
	return ans
}
func minimumTeachingsW(n int, languages [][]int, friendships [][]int) int {
	luMap := make(map[int]map[int]bool)
	for i := 1; i <= n; i++ {
		luMap[i] = make(map[int]bool)
	}
	m := len(languages)
	for i := 0; i < m; i++ {
		for j := 0; j < len(languages[i]); j++ {
			learned := languages[i][j]
			luMap[learned][i+1] = true
		}
	}
	fTable := make(map[int][]int)
	for _, friend := range friendships {
		fTable[friend[0]] = append(fTable[friend[0]], friend[1])
		fTable[friend[1]] = append(fTable[friend[1]], friend[0])
	}
	ans := m
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	// 能沟通的意思，不是指，使用同一种语言沟通。
	// 所以可以避免
	for l := 1; l <= n; l++ {
		left := len(luMap[l])
		if left == m {
			ans = 0
			break
		}
		for u := range luMap[l] {
			for _, f := range fTable[u] {
				luMap[l][f] = true
			}
		}
		right := len(luMap[l])
		if right != m {
			continue
		}
		need := right - left
		fmt.Println("language i", l, "need", need)
		ans = min(ans, need)
	}
	fmt.Println("ansssss", ans)
	return ans
}
func minimumTeachingsOld(n int, languages [][]int, friendships [][]int) int {
	fTable := make(map[int][]int)
	lessonMap := make(map[int][]int)
	m := len(languages)
	for i := 0; i < m; i++ {
		for _, language := range languages[i] {
			lessonMap[language] = append(lessonMap[language], i+1)
		}
	}
	for _, friend := range friendships {
		fTable[friend[0]] = append(fTable[friend[0]], friend[1])
		fTable[friend[1]] = append(fTable[friend[1]], friend[0])
	}
	// fmt.Println(lessonMap)
	// fmt.Println(fTable)
	ans := n
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for l := 1; l <= n; l++ {
		// 计算第 i 个语言需要的数量
		users := make([]bool, m+1)
		fmt.Println("lesson", lessonMap[l])
		for _, begin := range lessonMap[l] {
			users[begin] = true
		}
		need := 0
		pass := len(lessonMap[l])
		if pass != m {
			for _, begin := range lessonMap[l] {
				for _, f := range fTable[begin] {
					if !users[f] {
						need++
						users[f] = true
					}
				}
			}
		}
		if pass+need != m {
			// 当前的不符合全部人能学会
			continue
		}
		fmt.Println("language i ", l, " need ", need)
		ans = min(ans, need)
	}
	fmt.Println("ansss", ans)
	return ans
}
func func3() {
	// vector<int> ans(encoded.size()+1);
	//       ans[0]=first;
	//       for(int i=0;i<encoded.size();i++)
	//           ans[i+1]=ans[i]^encoded[i];//若a^b=c,则a=b^c
	//       return ans;
	// 输入：encoded = [3,1]
	// 输出：[1,2,3]
	// 解释：如果 perm = [1,2,3] ，那么 encoded = [1 XOR 2,2 XOR 3] = [3,1]
	// 示例 2：

	// 输入：encoded = [6,5,4,6]
	// 输出：[2,4,1,5,3]
	decode([]int{3, 1})
	decode([]int{6, 5, 4, 6})
	decode([]int{3, 12, 1, 15, 5, 2, 3, 7})

	// 	输入：
	// [3,12,1,15,5,2,3,7]
	// 输出：
	// [6,5,9,8,7,2,0,3,4]
	// 预期：
	// [7,4,8,9,6,3,1,2,5]
}
func decode(encoded []int) []int {
	n := len(encoded)
	var ans []int
	m := n + 1
	for i := 1; i <= m; i++ {
		cur := make([]int, m)
		cur[0] = i
		uniq := make([]bool, m+1)
		uniq[i] = true
		flag := true
		for j := 0; j < n; j++ {
			cur[j+1] = cur[j] ^ encoded[j]
			if cur[j+1] <= 0 || cur[j+1] > m {
				flag = false
				break
			}
			if uniq[cur[j+1]] {
				flag = false
				break
			}
			uniq[cur[j+1]] = true
		}
		if flag {
			ans = cur
			break
		}
	}
	fmt.Println("ansssss", ans)
	return ans
}
