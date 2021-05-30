/*
 * @lc app=leetcode.cn id=990 lang=golang
 *
 * [990] 等式方程的可满足性
 */
package main

// func main() {
// 	equationsPossible([]string{"a==b", "b!=a"})
// 	equationsPossible([]string{"a==b", "b==a"})
// 	equationsPossible([]string{"a==b", "b==c", "a==c"})
// 	equationsPossible([]string{"a==b", "b!=c", "c==a"})
// 	equationsPossible([]string{"c==c", "b==d", "x!=z"})
// }

// @lc code=start
func equationsPossible(equations []string) bool {
	eq := make(map[byte]map[byte]bool)
	n := len(equations)
	for i := 0; i < n; i++ {
		exp := equations[i]
		left := exp[0]
		right := exp[3]
		equal := exp[1] == '='
		if equal {
			if eq[left] == nil {
				eq[left] = make(map[byte]bool)
			}
			eq[left][right] = true
			if eq[right] == nil {
				eq[right] = make(map[byte]bool)
			}
			eq[right][left] = true
		} else {

		}
	}
	var bfs func(start, end byte) bool
	bfs = func(start, end byte) bool {
		if start == end {
			return true
		}
		visited := make(map[byte]bool)
		queue := make([]byte, 0)
		queue = append(queue, start)
		visited[start] = true
		for len(queue) != 0 {
			size := len(queue)
			for i := 0; i < size; i++ {
				node := queue[i]
				for next := range eq[node] {
					if !visited[next] {
						if next == end {
							return true
						}
						queue = append(queue, next)
						visited[next] = true
					}
				}
			}
			queue = queue[size:]
		}
		return false
	}
	for i := 0; i < n; i++ {
		exp := equations[i]
		left := exp[0]
		right := exp[3]
		equal := exp[1] == '='
		if !equal {
			// 找出 left 和 right 是否已经相等了
			ret := bfs(left, right)
			if ret {
				// fmt.Println("falseeee", string(left), string(right))
				return false
			}
			ret = bfs(right, left)
			if ret {
				// fmt.Println("falseeee", string(left), string(right))
				return false
			}
		}
	}
	// fmt.Println("trueeeeeeeeeeeeeeeee")
	return true
}

// @lc code=end
