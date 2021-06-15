package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func alertNames(keyName []string, keyTime []string) []string {
	m := make(map[string][]int)
	n := len(keyName)
	for i := 0; i < n; i++ {
		mini := parseKeyTime(keyTime[i])
		user := keyName[i]
		if _, ok := m[user]; !ok {
			m[user] = make([]int, 0)
		}
		m[user] = append(m[user], mini)
	}
	fmt.Println(m)
	ans := make([]string, 0)
	for k, v := range m {
		if pass(v) {
			ans = append(ans, k)
		}
	}
	sort.Strings(ans)
	return ans
}
func pass(counter []int) bool {
	sort.Ints(counter)
	for i := 0; i < len(counter); i++ {
		// num := 0
		// for j := i + 1; j < len(counter) && j < i+3; j++ {
		// 	if counter[j]-counter[i] <= 60 {
		// 		num++
		// 	}
		// }
		// if num >= 2 {
		// 	return true
		// }
		if i+2 < len(counter) && counter[i+2]-counter[i] <= 60 {
			return true
		}
	}
	return false
}
func parseKeyTime(keyTime string) int {
	if keyTime == "" {
		return 0
	}
	strArr := strings.Split(keyTime, ":")
	hour, _ := strconv.Atoi(strArr[0])
	mini, _ := strconv.Atoi(strArr[1])
	return hour*60 + mini
}

func main() {
	// fmt.Println(
	// 	alertNames(
	// 		[]string{"daniel", "daniel", "daniel", "luis", "luis", "luis", "luis"},
	// 		[]string{"10:00", "10:40", "11:00", "09:00", "11:00", "13:00", "15:00"},
	// 	),
	// )
	tests := []struct {
		keyName []string
		keyTime []string
	}{
		{
			[]string{"daniel", "daniel", "daniel", "luis", "luis", "luis", "luis"},
			[]string{"10:00", "10:40", "11:00", "09:00", "11:00", "13:00", "15:00"},
		},
		{
			[]string{"alice", "alice", "alice", "bob", "bob", "bob", "bob"},
			[]string{"12:01", "12:00", "18:00", "21:00", "21:20", "21:30", "23:00"},
		},
		{
			[]string{"john", "john", "john"},
			[]string{"23:58", "23:59", "00:01"},
		},
		{
			[]string{"leslie", "leslie", "leslie", "clare", "clare", "clare", "clare"},
			[]string{"13:00", "13:20", "14:00", "18:00", "18:51", "19:30", "19:49"},
		},
	}
	for _, test := range tests {
		fmt.Println(alertNames(test.keyName, test.keyTime))
		fmt.Println("---------------")
	}
}
