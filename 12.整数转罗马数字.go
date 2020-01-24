/*
 * @lc app=leetcode.cn id=12 lang=golang
 *
 * [12] 整数转罗马数字
 */

// @lc code=start
func intToRoman(num int) string {

	list := make([]int, 0)
	for num != 0 {
		m := num%10
		list = append(list, m)
		num/=10
	}
	fmt.Println(list)
	ans := ""
	multi := 1
	for i:=0;i<len(list);i++ {
		if list[i]!=0 {
			tmp:= multi * list[i]
			ans = trans(tmp, list[i], i)+ans
		}	
		multi *= 10
	}
	return ans
}
func trans(num int, m int, base int)string {
	sMap := map[int]string{
		1000: "M",
		900: "CM",
		500: "D",
		400: "CD",
		100: "C",
		90: "XC",
		50: "L",
		40: "XL",
		10: "X",
		9: "IX",
		5: "V",
		4: "IV",
		1: "I",
	}
	str,ok := sMap[num]
	fmt.Println(num, m, base)
	if ok {
		return str
	}
	baseMap := map[int]string{
		0: "I",
		1: "X",
		2: "C",
		3: "M",
	}
	fiveMap := map[int]string {
		0: "V",
		1:"L",
		2:"D",
	}
	ans := ""
	b := baseMap[base]
	// 30 3 80 8
	if 0 < m && m < 5 {
		// m==123
	} else if 5 < m && m < 10 {
		// m=678
		m-=5
		five,_ := fiveMap[base]
		ans+=five
	}
	ans += strings.Repeat(b, m)
	fmt.Println(ans)
	return ans
}
// @lc code=end

