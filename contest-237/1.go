package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	getXORSum([]int{1, 2, 3}, []int{6, 5})
	getXORSum([]int{12}, []int{4})
	// getOrder([][]int{{1, 2}, {2, 4}, {3, 2}, {4, 1}})
	// getOrder([][]int{{7, 10}, {7, 12}, {7, 5}, {7, 4}, {7, 2}})
	// getOrder([][]int{{19, 13}, {16, 9}, {21, 10}, {32, 25}, {37, 4}, {49, 24}, {2, 15}, {38, 41}, {37, 34}, {33, 6}, {45, 4}, {18, 18}, {46, 39}, {12, 24}})
}
func getXORSum(arr1 []int, arr2 []int) int {
	n := len(arr1)
	m := len(arr2)
	ans := 0
	cur := 0
	j := 0
	for i := 0; i < n; i++ {
		cur = cur ^ (arr1[i] & arr2[j])
	}
	cur ^= 0
	ans ^= cur
	j++
	for j < m {
		cur &= arr2[j]
		j++
		ans ^= cur
	}
	ans = ans ^ 0
	// fmt.Println("ansssss", ans)
	return ans
}
func getOrder(tasks [][]int) []int {
	n := len(tasks)
	for i := 0; i < n; i++ {
		tasks[i] = append(tasks[i], i)
	}
	// task[2] == 原来的 index
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i][0] < tasks[j][0]
	})
	// fmt.Println(tasks)
	ans := make([]int, n)
	finished := 0
	used := make([]bool, n)
	queue := make(map[int]bool)
	start := tasks[0][0]
	inq := 0
	for inq < n {
		if tasks[inq][0] <= start {
			queue[inq] = true
			inq++
		} else {
			break
		}
	}
	getTask := func() int {
		process := math.MaxInt32
		nowIndex := math.MaxInt32
		target := -1
		// fmt.Println(queue)
		for kIndex := range queue {
			task := tasks[kIndex]
			if used[kIndex] {
				continue
			}
			if task[0] <= start {
				if task[1] < process {
					target = kIndex
					process = task[1]
					nowIndex = task[2]
				} else if task[1] == process {
					if task[2] < nowIndex {
						target = kIndex
						process = task[1]
						nowIndex = task[2]
					}
				}
			}
		}
		if target == -1 {
			fmt.Println("wrong target", target)
		}
		// 有可能 start 太小了
		delete(queue, target)
		return target
	}
	for finished != n {
		// 在 queue 里面找到一个合适的进行执行
		taskIndex := getTask()
		ans[finished] = tasks[taskIndex][2]
		finished++
		used[taskIndex] = true
		start += tasks[taskIndex][1]
		// fmt.Println("now", start, "finish task", taskIndex)
		// fmt.Println("queue", queue)
		for inq < n {
			if tasks[inq][0] <= start {
				queue[inq] = true
				inq++
			} else {
				break
			}
		}
	}
	fmt.Println("ansssss", ans)
	return ans
}
func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)
	ans := 0
	n := len(costs)
	for i := 0; i < n; i++ {
		if coins < costs[i] {
			break
		}
		ans++
		coins -= costs[i]
	}
	return ans
}
func checkIfPangram(sentence string) bool {
	m := make(map[byte]bool)
	n := len(sentence)
	for i := 0; i < n; i++ {
		m[sentence[i]] = true
	}
	fmt.Println(m)
	return len(m) == 26
}
