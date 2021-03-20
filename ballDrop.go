package main

import "fmt"

// func main() {
// 	ballDrop(100, 10)
// }
func ballDrop(height float64, times int) float64 {
	var s float64 = height
	var h float64 = height
	for i := 1; i < times; i++ {
		s += h
		h /= 2
	}
	fmt.Println("anssss", s)
	return s
}
