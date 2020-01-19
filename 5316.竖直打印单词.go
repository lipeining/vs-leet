func printVertically(s string) []string {
	words := strings.Split(s, " ")
	// 倒叙生成结果才对，不足的补空格就好
	maxLength := 0
	for i:=0;i<len(words);i++ {
		if len(words[i]) > maxLength {
			maxLength = len(words[i])
		}
	}
	ans := make([]string, maxLength)
	for i:=0;i<maxLength;i++ {
		tmp:=""
		for j:=0;j<len(words);j++ {
			if i >= len(words[j]) {
				tmp += " "
			} else {
				tmp += string(words[j][i])
			}			
		}
		ans[i] = strings.TrimRight(tmp, " ")
	}
	return ans
}