package main

import "fmt"

func main() {

}
func nearestExit(maze [][]byte, entrance []int) int {
	m := len(maze)
	n := len(maze[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	si, sj := entrance[0], entrance[1]
	visited[si][sj] = true
	toId := func(i, j int) int {
		return i*n + j
	}
	toIJ := func(num int) (int, int) {
		return num / n, num % n
	}
	step := 0
	queue := make([]int, 0)
	queue = append(queue, toId(si, sj))
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(queue) != 0 {
		size := len(queue)
		for k := 0; k < size; k++ {
			id := queue[k]
			i, j := toIJ(id)
			if i == 0 || i == m-1 || j == 0 || j == n-1 {
				if i != si || j != sj {
					fmt.Println("ans,,,,", i, j, step)
					return step
				}
			}
			for _, dir := range dirs {
				x, y := i+dir[0], j+dir[1]
				if x < 0 || y < 0 || x >= m || y >= n {
					continue
				}
				if maze[x][y] != '.' {
					continue
				}
				if visited[x][y] {
					continue
				}
				visited[x][y] = true
				queue = append(queue, toId(x, y))
			}
		}
		step++
		queue = queue[size:]
	}
	return -1
}
func countTriples(n int) int {
	targetm := make(map[int]bool)
	for i := 1; i <= n; i++ {
		targetm[i*i] = true
	}
	ans := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			sum := i*i + j*j
			if targetm[sum] {
				ans++
			}
		}
	}
	return ans
}
