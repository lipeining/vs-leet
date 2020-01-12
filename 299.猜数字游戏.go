/*
 * @lc app=leetcode.cn id=299 lang=golang
 *
 * [299] 猜数字游戏
 */
import "fmt"
// @lc code=start
func getHint(secret string, guess string) string {
    smap := make(map[string]int)
	gmap := make(map[string]int)
	bullCounter := 0
	cowCounter := float64(0)
	for i:=0; i<len(secret); i++ {
		if secret[i] == guess[i] {
			bullCounter++
		} else {
			if sc,ok := smap[string(secret[i])]; ok {
				smap[string(secret[i])] = sc + 1
			} else {
				smap[string(secret[i])] = 1
			}
			if gc,ok := gmap[string(guess[i])]; ok {
				gmap[string(guess[i])] = gc + 1
			} else {
				gmap[string(guess[i])] = 1
			}
		}
	}
	for k,v := range smap {
		if gc, ok := gmap[k]; ok {
			cowCounter += math.Min(float64(v), float64(gc))
		}
	}
	fmt.Println(smap)
	fmt.Println(gmap)
	fmt.Println(bullCounter)
	fmt.Println(cowCounter)
	return strconv.Itoa(bullCounter) + "A" + strconv.Itoa(int(cowCounter)) + "B"
}
// @lc code=end

