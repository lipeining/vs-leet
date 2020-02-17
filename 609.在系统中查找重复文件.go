import "strings"

/*
 * @lc app=leetcode.cn id=609 lang=golang
 *
 * [609] 在系统中查找重复文件
 */

// @lc code=start
func findDuplicate(paths []string) [][]string {
	length := len(paths)
	ans := make([][]string, 0)
	if length == 0 {
		return ans
	}
	m := make(map[string][]string)
	for i := 0; i < length; i++ {
		arr1 := strings.Split(paths[i], " ")
		for j := 1; j < len(arr1); j++ {
			arr2 := strings.SplitN(arr1[j], "(", 2)
			filePath := arr1[0] + "/" + arr2[0]
			fileContent := arr2[1]
			if _, ok := m[fileContent]; !ok {
				m[fileContent] = make([]string, 0)
			}
			m[fileContent] = append(m[fileContent], filePath)
		}
	}
	for _, v := range m {
		if len(v) > 1 {
			ans = append(ans, v)
		}
	}
	return ans
}

// @lc code=end
