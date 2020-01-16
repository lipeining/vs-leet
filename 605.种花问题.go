/*
 * @lc app=leetcode.cn id=605 lang=golang
 *
 * [605] 种花问题
 */

// @lc code=start
func canPlaceFlowers(flowerbed []int, n int) bool {
	res := 0
	for i:=0; i <len(flowerbed);i++ {
		if flowerbed[i] != 1 {
			l := i - 1
			r := i + 1
			if i==0 {
				l = i
			} 
			if i == len(flowerbed) - 1 {
				r = i
			}
			if flowerbed[l] == 0 && flowerbed[r] == 0 {
				res++
				flowerbed[i] = 1
			}
		}
	}
	fmt.Println(flowerbed, res)
	return res >= n
}
// @lc code=end

