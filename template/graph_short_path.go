package template

import (
	"fmt"
	"math"
	"strconv"
)

// https://zhuanlan.zhihu.com/p/33162490

// 使用邻接矩阵存储图之间的距离 a := [][]int{}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// // f[k][i][j] = min(f[k-1][i][j],f[k-1][i][k]+f[k-1][k][j])
// func floyd(a [][]int, n int) {
// 	for k := 0; k < n; k++ {
// 		for i := 0; i < n; i++ {
// 			for j := 0; j < n; j++ {
// 				a[i][j] = min(a[i][j], a[i][k]+a[k][j])
// 			}
// 		}
// 	}
// 	for i := 0; i < n; i++ {
// 		for j := i + 1; j < n; j++ {
// 			line := strconv.Itoa(i) + " " + strconv.Itoa(j) + ":" + strconv.Itoa(a[i][j])
// 			fmt.Println(line)
// 		}
// 	}
// }

func dijkstra(p int, n int, a [][]int) {
	dist := make([]int, n)
	set := make(map[int]bool)
	set[p] = true
	for i := 0; i < n; i++ {
		dist[i] = a[p][i]
	}
	for len(set) != n {
		le := math.MaxInt32
		num := 0
		for i := 0; i < n; i++ {
			if !set[i] && le > dist[i] {
				le = dist[i]
				num = i
			}
		}
		if le == math.MaxInt32 {
			break
		}
		for i := 0; i < n; i++ {
			if !set[i] {
				dist[i] = min(dist[i], dist[num]+a[num][i])
			}
		}
		set[num] = true
	}
	for i := 0; i < n; i++ {
		line := "点" + strconv.Itoa(p) + "到点" + strconv.Itoa(i) + "的距离为：" + strconv.Itoa(dist[i])
		fmt.Println(line)
	}
}

func bellmanFord(s, n int, a [][]int) bool {
	d := make([]int, n)
	for i := 0; i < n; i++ {
		d[i] = a[s][i]
	}
	for i := 0; i < n-1; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if a[j][k] != math.MaxInt32 {
					d[k] = min(d[k], d[j]+a[j][k])
				}
			}
		}
	}
	for j := 0; j < n; j++ {
		for k := 0; k < n; k++ {
			if a[j][k] != math.MaxInt32 {
				if d[k] > d[j]+a[j][k] {
					return false
				}
			}
		}
	}
	return true
}

func spfa(s, n int, a [][]int) {
	queue := make([]int, 0)
	d := make([]int, n)
	flag := make([]bool, n)
	for i := 0; i < n; i++ {
		d[i] = math.MaxInt32
	}
	d[s] = 0
	queue = append(queue, s)
	flag[s] = true
	var u int
	for len(queue) != 0 {
		u = queue[0]
		queue = queue[1:]
		flag[u] = false
		for i := 0; i < n; i++ {
			if a[u][i] != math.MaxInt32 {
				temp := d[u] + a[u][i]
				if temp < d[i] {
					d[i] = temp
					if !flag[i] {
						queue = append(queue, i)
						flag[i] = true
					}
				}
			}
		}
	}
	for i := s + 1; i < n; i++ {
		line := "点" + strconv.Itoa(s) + "到点" + strconv.Itoa(i) + "的距离为：" + strconv.Itoa(d[i])
		fmt.Println(line)
	}
}
