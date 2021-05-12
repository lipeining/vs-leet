package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	closestRoom([][]int{{2, 2}, {1, 2}, {3, 2}}, [][]int{{3, 1}, {3, 3}, {5, 2}})
	closestRoom([][]int{{1, 4}, {2, 3}, {3, 5}, {4, 1}, {5, 2}}, [][]int{{2, 3}, {2, 4}, {2, 5}})
}
func closestRoom(rooms [][]int, queries [][]int) []int {
	sort.Slice(rooms, func(i, j int) bool {
		if rooms[i][1] == rooms[j][1] {
			return rooms[i][0] < rooms[j][0]
		}
		return rooms[i][1] < rooms[j][1]
	})
	fmt.Println(rooms)
	// 房间按照 size 排序，再按照 id 排序
	n := len(queries)
	ans := make([]int, n)
	bs := func(pid, msize int) int {
		l := 0
		r := len(rooms) - 1
		for l < r {
			mid := l + (r-l)/2
			if rooms[mid][1] < msize {
				l = mid + 1
			} else if rooms[mid][1] > msize {
				r = mid
			} else {
				if rooms[mid][0] > pid {
					r--
				} else if rooms[mid][0] < pid {
					l++
				} else {
					return rooms[mid][0]
				}
				// // 左右遍历吧
				// j := mid
				// for j >= 0 {
				// 	if rooms[j][1] == msize {
				// 		j--
				// 	} else {
				// 		break
				// 	}
				// }
				// return j + 1
			}
		}
		// 此时 l = r
		if rooms[l][1] < msize {
			return -1
		}
		return rooms[l][0]
	}
	for i := 0; i < n; i++ {
		pid, msize := queries[i][0], queries[i][1]
		ans[i] = bs(pid, msize)
	}
	fmt.Println(ans)
	return ans
}
func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	sort.Ints(arr)
	arr[0] = 1
	n := len(arr)
	ans := 0
	for i := 1; i < n; i++ {
		if arr[i]-arr[i-1] <= 1 {
			// do nothing
		} else {
			arr[i] = arr[i-1] + 1
		}
	}
	ans = arr[n-1]
	return ans
}

type pair struct{ t, i int }
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { a, b := h[i], h[j]; return a.t < b.t }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func (h *hp) push(v pair)        { heap.Push(h, v) }
func (h *hp) pop() pair          { return heap.Pop(h).(pair) }

type SeatManager struct {
	seats hp
}

func Constructor(n int) SeatManager {
	seats := make(hp, n)
	for i := 0; i < n; i++ {
		seats[i] = pair{t: i + 1, i: i + 1}
	}
	heap.Init(&seats)
	return SeatManager{seats: seats}
}

func (this *SeatManager) Reserve() int {
	p := this.seats.pop()
	return p.t
}

func (this *SeatManager) Unreserve(seatNumber int) {
	this.seats.push(pair{t: seatNumber, i: seatNumber})
}

func replaceDigits(s string) string {
	// table := "abcdefghijklmnopqrstuvwxyz"
	n := len(s)
	ans := ""
	shift := func(char byte, offset int) byte {
		return byte(int(char) + offset)
	}
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			ans += string(s[i])
		} else {
			ans += string(shift(s[i-1], int(s[i]-'0')))
		}
	}
	return ans
}
