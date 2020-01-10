/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */
// import "fmt";
// @lc code=start
func romanToInt(s string) int {
	sMap := map[string]int{"M":1000,"CM":900,"D":500,"CD":400,"C":100,"XC":90,"L":50,"XL":40,"X":10,"IX":9,"V":5,"IV":4,"I":1}
	res := 0
	length := len(s)
	for i := 0; i < length; i++ {
		tmp := 0;
		if i < length - 1 {
			if _, ok := sMap[s[i:i+2]]; ok {
				tmp = sMap[s[i:i+2]];
				i++;
				res = res + tmp;
				continue;
			}
		}
		tmp = sMap[string(s[i])];
		res = res + tmp;
	}
	return res;
	// sMap := map[string]int     {"M":1000,"CM":900,"D":500,"CD":400,"C":100,"XC":90,"L":50,"XL":40,"X":10,"IX":9,"V":5,"IV":4,"I":1}
    // result :=0
    // for i:=0; i<len(s); i++ {
    //     if i<len(s)-1 {
    //         if _, ok := sMap[s[i:i+2]]; ok {
    //             result = result+sMap[s[i:i+2]]
    //             i++
    //             continue
    //         }
    //     } 
    //     if _, ok := sMap[s[i:i+1]]; ok {
    //         result = result+sMap[s[i:i+1]]
    //     }
    // }
    // return result
}
// @lc code=end

