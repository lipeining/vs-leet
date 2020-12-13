package main

import (
	"fmt"
	"math"
	"sort"
)

func test1610() {
	tests := []struct {
		points   [][]int
		angle    int
		location []int
	}{
		{
			points:   [][]int{{2, 1}, {2, 2}, {3, 3}},
			angle:    90,
			location: []int{1, 1},
		},
		{
			points:   [][]int{{2, 1}, {2, 2}, {3, 4}, {1, 1}},
			angle:    90,
			location: []int{1, 1},
		},
		{
			points:   [][]int{{1, 0}, {2, 1}},
			angle:    13,
			location: []int{1, 1},
		},

		{
			points:   [][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {1, 2}, {2, 1}},
			angle:    0,
			location: []int{1, 1},
		},
	}
	//     输入：points = [[2,1],[2,2],[3,3]], angle = 90, location = [1,1]
	// 输出：3
	// 解释：阴影区域代表你的视野。在你的视野中，所有的点都清晰可见，尽管 [2,2] 和 [3,3]在同一条直线上，你仍然可以看到 [3,3] 。
	// 示例 2：

	// 输入：points = [[2,1],[2,2],[3,4],[1,1]], angle = 90, location = [1,1]
	// 输出：4
	// 解释：在你的视野中，所有的点都清晰可见，包括你所在位置的那个点。
	//     输入：points = [[1,0],[2,1]], angle = 13, location = [1,1]
	// 输出：1
	// 解释：如图所示，你只能看到两点之一。

	//     输入：
	// [[1,1],[2,2],[3,3],[4,4],[1,2],[2,1]]
	// 0
	// [1,1]
	// 输出：
	// 3
	// 预期结果：
	// 4
	for _, test := range tests {
		fmt.Println(visiblePoints(test.points, test.angle, test.location))
		fmt.Println("---------------")
	}
}
func visiblePoints(points [][]int, angle int, location []int) int {
	same := 0
	angles := make([]float64, 0)
	for _, point := range points {
		if checkSame(point, location) {
			same++
			continue
		}
		y := point[1] - location[1]
		x := point[0] - location[0]
		a := math.Atan2(float64(y), float64(x)) * 180 / math.Pi
		if a < 0 {
			a = a + 360
		}
		angles = append(angles)
	}
	sort.Slice(angles, func(i, j int) bool {
		return angles[i] < angles[j]
	})
	fmt.Println(angles)
	n := len(angles)
	for i := 0; i < n; i++ {
		angles = append(angles, angles[i]+360)
	}
	anw := same
	j := 0
	for i := 0; i < len(angles); i++ {
		for j < len(angles) {
			if angles[j]-angles[i] < float64(angle)+float64(1e-5) {
				j++
			} else {
				break
			}
		}
		anw = max(anw, j-i+same)
	}
	return anw
}
func visiblePoints2(points [][]int, angle int, location []int) int {
	// 计算 points  到 location 的 angle
	// n := len(points)
	same := 0
	type node struct {
		point []int
		angle int
	}
	queue := make([]node, 0)
	for _, point := range points {
		if checkSame(point, location) {
			same++
			continue
		}
		angle = countAngle(point, location)
		queue = append(queue, node{point, angle})
	}
	// 视野就是角度范围 [d - angle/2, d + angle/2] 所指示的那片区域

	sort.Slice(queue, func(i, j int) bool {
		return queue[i].angle < queue[j].angle
	})
	morePoints := 0
	fmt.Println(queue)
	for i := 0; i < len(queue); i++ {
		cur := 0
		for j := i; j < len(queue); j++ {
			if queue[j].angle <= queue[i].angle+angle {
				cur++
			} else {
				break
			}
		}
		morePoints = max(morePoints, cur)
	}
	return morePoints + same
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func checkSame(point, location []int) bool {
	return point[0] == location[0] && point[1] == location[1]
}

//   x = abs(x1 - x2)
//   y = abs(y1 - y2)
//   z = math.sqrt(x * x + y * y)
//   angle = round(math.asin(y / z) / math.pi * 180)
func countAngle(point, location []int) int {
	dy := abs(point[1] - location[1])
	dx := abs(point[0] - location[0])
	z := math.Sqrt(float64(dx*dx) + float64(dy*dy))
	angle := math.Asin(float64(dy)/z) / math.Pi * float64(180)
	fmt.Println(point, " angle is", angle)
	angle = math.Floor(angle)
	return int(angle)
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
