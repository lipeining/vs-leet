package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1 := format(version1)
	v2 := format(version2)
	i, j := 0, 0
	fmt.Println(v1, v2)
	for i < len(v1) && j < len(v2) {
		if v1[i] > v2[j] {
			return 1
		} else if v1[i] < v2[j] {
			return -1
		}
		i++
		j++
	}
	if i != len(v1) {
		return 1
	}
	if j != len(v2) {
		return -1
	}
	return 0
}
func format(version string) []int {
	arr := strings.Split(version, ".")
	ans := make([]int, len(arr))
	skipZero := true
	keepLen := len(arr)
	for i := len(arr) - 1; i >= 0; i-- {
		num, _ := strconv.Atoi(arr[i])
		if num != 0 {
			skipZero = false
		}
		if skipZero {
			keepLen--
		}
		ans[i] = num
	}
	// fix  "0"
	if keepLen == 0 {
		return []int{0}
	}
	return ans[0:keepLen]
}

func main() {
	tests := []struct {
		v1     string
		v2     string
		expect int
	}{
		{"0.1", "1.1", -1},
		{"1.0.1", "1", 1},
		{"7.5.2.4", "7.5.3", -1},
		{"1.01", "1.0001", 0},
		{"1.0", "1.0.0", 0},
		{"0", "1", -1},
		{"1.0", "1.1", -1},
	}
	for _, t := range tests {
		ans := compareVersion(t.v1, t.v2)
		fmt.Println(ans, t.expect)
	}
}
