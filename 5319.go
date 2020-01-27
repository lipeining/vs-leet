func removePalindromeSub(s string) int {
    length := len(s)
    if length == 0 {
        return 0
    }
    start := 0
    end := length
    ans := 0
    for {
        if helper(s[start:end]) {
            ans++
            if end == length {
                break
            } else {
                start = end
                end = length
            }
        } else {
           end--
        }  
    }
    return ans
}
func helper(s string) bool {
    for i:=0;i<len(s)/2;i++ {
        if s[i] != s[len(s)-i-1] {
            return false
        }
    }
    return true
}