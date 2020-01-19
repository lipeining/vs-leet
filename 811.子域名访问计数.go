/*
 * @lc app=leetcode.cn id=811 lang=golang
 *
 * [811] 子域名访问计数
 */

import "strings"
// @lc code=start
func subdomainVisits(cpdomains []string) []string {
	counter := make(map[string]int)
	for _,cp := range cpdomains {
		arr := strings.Split(string(cp), " ")
		num,_ := strconv.Atoi(arr[0])
		// domains := strings.Split(arr[1], '.')
		tmpStr := arr[1]
		counter[tmpStr] = num
		index := strings.LastIndex(tmpStr, ".")
		for index != -1 {
			tmp := tmpStr[index+1:]
			n, ok := counter[tmp]
			if ok {
				counter[tmp] = n + 1
			} else {
				counter[tmp] = num
			}
			fmt.Println(tmpStr, index, tmp)
			tmpStr = tmpStr[:index]
			index = strings.LastIndex(tmpStr, ".")	
		}
	}
	ans := make([]string, 0)
	for k,v := range counter {
		ans = append(ans, strconv.Itoa(v) + " " + k)
	}
	return ans
}
// @lc code=end

