/*
 * @lc app=leetcode.cn id=970 lang=golang
 *
 * [970] 强整数
 */

// @lc code=start
func powerfulIntegers(x int, y int, bound int) []int {
	// ans := make([]int,0)
	// xmap := getMap(x, bound)
	// ymap := getMap(y, bound)
	// tmap := make(map[int]bool)
	// for x0,_ := range xmap {
	// 	for y0,_ := range ymap {
	// 		tmp := x0+y0
	// 		if tmp <= bound {
	// 			tmap[tmp] = true
	// 		}
	// 	}
	// }
	// for k,_ := range tmap {
	// 	ans = append(ans, k)
	// }
	// return ans
	uniq := make(map[int]bool)
	for i:=0;i<20;i++ {
		for j:=0;j<20;j++ {
			tmp := int(math.Pow(float64(x), float64(i))) + int(math.Pow(float64(y), float64(j)))
			if tmp <= bound {
				uniq[tmp] = true
			}
		}
	}
	ans := make([]int, len(uniq))
	i:=0
	for k,_ := range uniq {
		ans[i] = k
		i++
	}
	return ans
}
// func getMap(a int, bound int) map[int]int {
// 	i:=0
// 	t:=make(map[int]int)
// 	for {
// 		tmp:=int(math.Pow(float64(a), float64(i)))
// 		if tmp >= bound {
// 			break
// 		} else {
// 			t[tmp]=i
// 			i++
// 		}
// 	}
// 	fmt.Println(t)
// 	return t
// }
// @lc code=end

