package main

import "bytes"

func checkMove(board [][]byte, rMove int, cMove int, color byte) bool {
	n := len(board)
	// n = 8
	dirs := [][]int{
		{0, -1},
		{0, 1},
		{-1, 0},
		{1, 0},
		{-1, -1},
		{1, 1},
		{-1, 1},
		{1, -1},
	}
	check := func(i, j int) bool {
		if i < 0 || i >= n || j < 0 || j >= n {
			return false
		}
		return true
	}
	for _, dir := range dirs {
		dx, dy := dir[0], dir[1]
		same := 0
		flag := true
		x0, y0 := rMove, cMove
		for {
			nx, ny := x0+dx, y0+dy
			if !check(nx, ny) {
				break
			}
			if board[nx][ny] == '.' {
				flag = false
				break
			} else if board[nx][ny] == color {
				same++
				if same > 1 {
					flag = false
					break
				}
			}
			x0, y0 = nx, ny
		}
		if flag && same == 1 {
			return true
		}
	}
	return false
}

func makeFancyString(s string) string {
	var buf bytes.Buffer
	i := 0
	n := len(s)
	c := 0
	for i < n {
		if c == 0 {
			buf.WriteByte(s[i])
			c++
		} else if c == 1 {
			if s[i] == s[i-1] {
				c++
			} else {
				c = 1
			}
			buf.WriteByte(s[i])
		} else {
			if s[i] == s[i-1] {
				//
			} else {
				c = 1
				buf.WriteByte(s[i])
			}
		}
		i++
	}
	return buf.String()
}
