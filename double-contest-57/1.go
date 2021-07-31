package main

func main() {}

func areOccurrencesEqual(s string) bool {
	n := len(s)
	cnt := make(map[byte]int)
	for i := 0; i < n; i++ {
		cnt[s[i]]++
	}
	num := 0
	for _, v := range cnt {
		if num == 0 {
			num = v
		} else {
			if num != v {
				return false
			}
		}
	}
	return true
}
