package fall

func numsGame(nums []int) []int {
	return nil
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func dis(a, b int) int {
	if a >= b {
		return abs(a - b - 1)
	}
	return abs(b - a - 1)
}
func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
