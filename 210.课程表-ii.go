/*
 * @lc app=leetcode.cn id=210 lang=golang
 *
 * [210] 课程表 II
 */

// @lc code=start
func findOrder(numCourses int, prerequisites [][]int) []int {
    indegrees := make([]int, numCourses)
	for i:=0;i<len(prerequisites);i++ {
		indegrees[prerequisites[i][0]]++
	}
	queue := make([]int, 0)
	// 将度为0的入队
	for i:=0;i<numCourses;i++ {
		if indegrees[i] == 0 {
			queue = append(queue, i)
		}
	}
	// 循环将最先学习的课程删除，只有开始学习的课程的度为1
	ans := make([]int, 0)
	for len(queue) != 0 {
		pre := queue[0]
		queue = queue[1:]
		numCourses--
		ans = append(ans, pre)
    //  如果这里不用 for 循环的话，需要维护一个邻接表
		for i:=0;i<len(prerequisites);i++ {
			if prerequisites[i][1] != pre {
				continue
			}
			indegrees[prerequisites[i][0]]--
			if indegrees[prerequisites[i][0]] == 0 {
				queue = append(queue, prerequisites[i][0])
			}
		}
	}
	if numCourses != 0 {
		t := make([]int, 0)
		return t
	}
	return ans
}
// @lc code=end

