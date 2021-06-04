package main

import (
	"container/heap"
	"fmt"
	"math"
	"strconv"
)

func main() {
	// assignTasks([]int{3, 3, 2}, []int{1, 2, 3, 2, 1, 2})
	// assignTasks([]int{3, 3, 2}, []int{10, 10, 10, 1, 1, 2})
	minSkips([]int{1, 3, 2}, 4, 2)
}

func minSkips(dist []int, speed int, hoursBefore int) int {
	// dp[i][j] 前 i 个，进行了 j 个跳过时，到达终点，花费的时间。
	n := len(dist)
	dp := make([][]float64, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]float64, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = math.MaxFloat32
		}
	}
	min := func(a, b float64) float64 {
		if a < b {
			return a
		}
		return b
	}
	EPS := 1e-7
	ceil := func(a float64) float64 {
		i := int(a)
		if a-float64(i) < EPS {
			return float64(i)
		}
		return float64(i + 1)
	}
	// dp[0][0] = float64(dist[0]) / float64(speed)
	dp[0][0] = 0
	// 到 i 个时，跳过 j 个，完成时间，
	// 然后最后一个不算休息
	for i := 1; i <= n; i++ {
		need := float64(dist[i-1]) / float64(speed)
		for j := 0; j <= i; j++ {
			if j != i {
				dp[i][j] = min(dp[i][j], ceil(dp[i-1][j]+need))
			}
			if j != 0 {
				dp[i][j] = min(dp[i][j], dp[i-1][j-1]+need)
			}
		}
	}
	for i := 0; i <= n; i++ {
		if dp[n][i] < float64(hoursBefore)+EPS {
			return i
		}
	}
	return -1
}

type pair struct{ q, i, free int }
type hp []pair

func (h hp) Len() int { return len(h) }
func (h hp) Less(i, j int) bool {
	a, b := h[i], h[j]
	if a.free != b.free {
		return a.free < b.free
	}
	if a.q != b.q {
		return a.q < b.q
	}
	return a.i < b.i
}
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func (h *hp) push(v pair)        { heap.Push(h, v) }
func (h *hp) pop() pair          { return heap.Pop(h).(pair) }

type minheap struct{ q, i, free int }
type mh []minheap

func (h mh) Len() int { return len(h) }
func (h mh) Less(i, j int) bool {
	a, b := h[i], h[j]
	if a.q != b.q {
		return a.q < b.q
	}
	return a.i < b.i
}
func (h mh) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *mh) Push(v interface{}) { *h = append(*h, v.(minheap)) }
func (h *mh) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func (h *mh) push(v minheap)     { heap.Push(h, v) }
func (h *mh) pop() minheap       { return heap.Pop(h).(minheap) }

func assignTasks(servers []int, tasks []int) []int {
	n := len(servers)
	bussy := make(hp, 0)
	free := make(mh, 0)
	for i := 0; i < n; i++ {
		free.push(minheap{
			i:    i,
			q:    servers[i],
			free: 0,
		})
	}
	m := len(tasks)
	ans := make([]int, m)
	for i := 0; i < m; i++ {
		// 将 Bussy 的到期的全部释放
		for bussy.Len() != 0 {
			if bussy[0].free <= i {
				top := bussy.pop()
				free.push(minheap{q: top.q, i: top.i, free: top.free})
			} else {
				break
			}
		}
		// 如果现在 free 为空的话
		if free.Len() == 0 {
			top := bussy.pop()
			fmt.Println("from bussy", top.i, top.free)
			ans[i] = top.i
			bussy.push(pair{q: top.q, i: top.i, free: top.free + tasks[i]})
		} else {
			now := free.pop()
			fmt.Println("from free", now.i)
			ans[i] = now.i
			bussy.push(pair{q: now.q, i: now.i, free: i + tasks[i]})
		}
	}
	fmt.Println("ansssssssss", ans)
	return ans
}
func maxValue(n string, x int) string {
	neg := false
	if n[0] == '-' {
		neg = true
		n = n[1:]
	}
	if !neg {
		// 找到第一个小于x的值
		i := 0
		for i < len(n) {
			number := int(n[i] - '0')
			if number < x {
				break
			}
			i++
		}
		return n[:i] + strconv.Itoa(x) + n[i:]
	} else {
		i := 0
		for i < len(n) {
			number := int(n[i] - '0')
			if number > x {
				break
			}
			i++
		}
		return "-" + n[:i] + strconv.Itoa(x) + n[i:]
	}
}
func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	toNum := func(str string) int {
		n := len(str)
		sum := 0
		for i := 0; i < n; i++ {
			now := int(str[i] - 'a')
			sum = sum*10 + now
		}
		return sum
	}
	return toNum(firstWord)+toNum(secondWord) == toNum(targetWord)
}
