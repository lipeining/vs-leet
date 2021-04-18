package main

import (
	"fmt"
	"sort"
)

// func main() {
// 	// func1()
// 	// func2()
// 	func3()
// 	// func4()
// 	// func5()
// }

func func1() {
	purchasePlans([]int{2, 5, 3, 5}, 6)
	purchasePlans([]int{2, 2, 1, 9}, 10)
	fmt.Println("ansssss")
}
func purchasePlans(nums []int, target int) int {
	n := len(nums)
	sort.Ints(nums)
	ans := 0
	mod := int(1e9 + 7)
	bin := func(start, end int, want int) int {
		l := start
		r := end
		for l < r {
			mid := l + (r-l)/2
			if nums[mid] > want {
				r = mid
			} else {
				l = mid + 1
			}
		}
		fmt.Println("find", want, l, r)
		if l > end {
			return -1
		}
		if nums[l] <= want {
			return l
		}
		return l - 1
	}
	fmt.Println(nums, "nums")
	for i := 0; i < n; i++ {
		left := nums[i]
		right := target - left
		index := bin(i+1, n-1, right)
		if index == -1 || index <= i {
			continue
		}
		fmt.Println("left", left, "right", right, index)
		ans += index - i
	}
	fmt.Println("anssss", ans)
	ans %= mod
	return ans
}
func func2() {
	orchestraLayout(3, 0, 2) // 3
	orchestraLayout(4, 1, 2) // 5
	// orchestraLayout(5, 0, 0) // 1
	// orchestraLayout(5, 1, 0) // 7
	// orchestraLayout(5, 2, 0) // 6
	// orchestraLayout(5, 2, 2) // 7
	// orchestraLayout(5, 1, 1) // 8
	// orchestraLayout(5, 1, 3) // 1
	orchestraLayout(5, 3, 3) // 3
	// orchestraLayout(5, 3, 1) // 5
	orchestraLayout(5, 2, 3) // 2
	// orchestraLayout(5, 2, 1) // 6
	// orchestraLayout(5, 3, 2) // 4
	fmt.Println("ansssss")
}
func orchestraLayout(num int, xPos int, yPos int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	row := min(xPos-0, num-1-xPos)
	col := min(yPos-0, num-1-yPos)
	rc := min(row, col)
	fmt.Println("rc", rc)
	q := num - (rc-1)*2
	out := 4*rc + 4*(rc*(num-2+q-2)/2)
	fmt.Println("out", out)
	q -= 2
	n := num
	i := rc
	j := rc
	cnt := out
	fmt.Println("i", i, "j", j, "and q is ", q)
	if i == xPos {
		cnt += yPos - j + 1
	} else if n-i-1 == xPos {
		cnt += (q-2)*2 + 2
		cnt += i + q - yPos
	} else if n-j-1 == yPos {
		cnt += (q-2)*1 + 1
		cnt += xPos - i + 1
	} else if j == yPos {
		cnt += (q-2)*3 + 3
		cnt += j + q - xPos
	}
	fmt.Println("for now ,it's ", cnt)
	ans := cnt % 9
	fmt.Println("anssssss", ans)
	if ans == 0 {
		return 9
	}
	return ans
}
func orchestraLayoutN(num int, xPos int, yPos int) int {
	// 1e9
	n := num
	i := 0
	j := 0
	// 不断收缩这个包围圈
	// y=-x 的角度出发
	cnt := 0
	q := n
	for {
		if i == xPos {
			cnt += yPos - j + 1
			break
		}
		if n-i-1 == xPos {
			cnt += (q-2)*2 + 2
			cnt += i + q - yPos
			break
		}
		if n-j-1 == yPos {
			cnt += (q-1)*1 + 1
			cnt += yPos - i + 1
			break
		}
		if j == yPos {
			cnt += (q-2)*3 + 3
			cnt += j + q - yPos
			break
		}
		cnt += (q-2)*4 + 4
		i++
		j++
		q -= 2
	}
	fmt.Println("for now ,it's ", cnt)
	ans := cnt % 9
	fmt.Println("anssssss", ans)
	if ans == 0 {
		return 9
	}
	return ans
}
func func3() {

	fmt.Println("ansssss")
}
func magicTower(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum <= 0 {
		return -1
	}
	n := len(nums)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		if nums[i] >= 0 {
			if i > 0 {
				for j := 0; j < i; j++ {
					dp[i][j] = dp[i-1][j] + nums[i]
				}
			} else {
				dp[i][0] = nums[i]
			}
		} else {
			// 如果不跳过，或者跳过
			if i == 0 {
				//
				dp[i][0] = 0
			} else {
				dp[i][0] = dp[i-1][0] + nums[i]
				for j := 1; j < i; j++ {
					dp[i][j+1] = dp[i-1][j]
				}
				// dp[i][1] = dp[i-1][0]
				// dp[i][2] = dp[i-1][1]
			}
		}
	}
	return 0
}
func func4() {
	escapeMaze([][]string{{".#.", "#.."}, {"...", ".#."}, {".##", ".#."}, {"..#", ".#."}})
	escapeMaze([][]string{{".#.", "..."}, {"...", "..."}})
	escapeMaze([][]string{{"...", "...", "..."}, {".##", "###", "##."}, {".##", "###", "##."}, {".##", "###", "##."}, {".##", "###", "##."}, {".##", "###", "##."}, {".##", "###", "##."}})
	fmt.Println("ansssss")
}
func escapeMaze(maze [][]string) bool {
	if len(maze) == 0 {
		return false
	}
	n := len(maze[0])
	m := len(maze[0][1])
	length := len(maze)
	ans := false
	var dfs func(i, j, time int, tem bool, fx, fy int)
	dfs = func(i, j, time int, tem bool, fx, fy int) {
		if ans {
			return
		}
		if length == time {
			return
		}
		if i < 0 || i >= n || j < 0 || j >= m {
			return
		}
		now := maze[time]
		if now[i][j] == '#' && !tem && (fx != i || fy != j) {
			return
		}
		if i == n-1 && j == m-1 {
			fmt.Println("anssss now", now, tem, "fs, fy ", fx, fy)
			ans = true
			return
		}
		need := n - 1 - i + m - 1 - j
		if need > length-time {
			return
		}
		// 四个方向
		// 如果现在遇到了一个#
		if now[i][j] == '#' {
			if fx == i && fy == j {
				dfs(i, j, time+1, tem, fx, fy)
				dfs(i-1, j, time+1, tem, fx, fy)
				dfs(i+1, j, time+1, tem, fx, fy)
				dfs(i, j-1, time+1, tem, fx, fy)
				dfs(i, j+1, time+1, tem, fx, fy)
			} else {
				// 可以变为临时节点，或者，永久节点
				if tem {
					dfs(i, j, time+1, !tem, fx, fy)
					dfs(i-1, j, time+1, !tem, fx, fy)
					dfs(i+1, j, time+1, !tem, fx, fy)
					dfs(i, j-1, time+1, !tem, fx, fy)
					dfs(i, j+1, time+1, !tem, fx, fy)
				} else if fx == -1 && fy == -1 {
					dfs(i, j, time+1, tem, i, j)
					dfs(i-1, j, time+1, tem, i, j)
					dfs(i+1, j, time+1, tem, i, j)
					dfs(i, j-1, time+1, tem, i, j)
					dfs(i, j+1, time+1, tem, i, j)
				} else {
					// 没办法了
					return
				}
			}
		} else {
			dfs(i, j, time+1, tem, fx, fy)
			dfs(i-1, j, time+1, tem, fx, fy)
			dfs(i+1, j, time+1, tem, fx, fy)
			dfs(i, j-1, time+1, tem, fx, fy)
			dfs(i, j+1, time+1, tem, fx, fy)
		}
	}
	dfs(0, 0, 0, true, -1, -1)
	fmt.Println("ansss", ans)
	return ans
}
func func5() {

	fmt.Println("ansssss")
}
