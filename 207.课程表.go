/*
 * @lc app=leetcode.cn id=207 lang=golang
 *
 * [207] 课程表
 */

// @lc code=start

// public boolean canFinish(int numCourses, int[][] prerequisites) {
// 	int[][] adjacency = new int[numCourses][numCourses];
// 	int[] flags = new int[numCourses];
// 	for(int[] cp : prerequisites)
// 		adjacency[cp[1]][cp[0]] = 1;
// 	for(int i = 0; i < numCourses; i++){
// 		if(!dfs(adjacency, flags, i)) return false;
// 	}
// 	return true;
// }
// private boolean dfs(int[][] adjacency, int[] flags, int i) {
// 	if(flags[i] == 1) return false;
// 	if(flags[i] == -1) return true;
// 	flags[i] = 1;
// 	for(int j = 0; j < adjacency.length; j++) {
// 		if(adjacency[i][j] == 1 && !dfs(adjacency, flags, j)) return false;
// 	}
// 	flags[i] = -1;
// 	return true;
// }

// 作者：jyd
// 链接：https://leetcode-cn.com/problems/course-schedule/solution/course-schedule-tuo-bu-pai-xu-bfsdfsliang-chong-fa/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func canFinish(numCourses int, prerequisites [][]int) bool {
	adjacency := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		adjacency[i] = make([]int, numCourses)
	}
	for i := 0; i < len(prerequisites); i++ {
		adjacency[prerequisites[i][0]][prerequisites[i][1]] = 1
	}
	flags := make([]int, numCourses)
	for i := 0; i < numCourses; i++ {
		if !dfs(adjacency, flags, i) {
			return false
		}
	}
	return true
}

func dfs(adjacency [][]int, flags []int, i int) bool {
	if flags[i] == 1 {
		return false
	}
	if flags[i] == -1 {
		return true
	}
	flags[i] = 1
	for j := 0; j < len(adjacency); j++ {
		if adjacency[i][j] == 1 && !dfs(adjacency, flags, j) {
			return false
		}
	}
	flags[i] = -1
	return true
}

// -----------------------------------------

// func canFinish(numCourses int, prerequisites [][]int) bool {
// 	indegrees := make([]int, numCourses)
// 	for i:=0;i<len(prerequisites);i++ {
// 		indegrees[prerequisites[i][0]]++
// 	}
// 	queue := make([]int, 0)
// 	// 将度为0的入队
// 	for i:=0;i<numCourses;i++ {
// 		if indegrees[i] == 0 {
// 			queue = append(queue, i)
// 		}
// 	}
// 	// 循环将最先学习的课程删除，只有开始学习的课程的度为1
// 	for len(queue) != 0 {
// 		pre := queue[0]
// 		queue = queue[1:]
// 		numCourses--
//      如果这里不用 for 循环的话，需要维护一个邻接表
// 		for i:=0;i<len(prerequisites);i++ {
// 			if prerequisites[i][1] != pre {
// 				continue
// 			}
// 			indegrees[prerequisites[i][0]]--
// 			if indegrees[prerequisites[i][0]] == 0 {
// 				queue = append(queue, prerequisites[i][0])
// 			}
// 		}
// 	}
// 	return numCourses == 0
// 	public boolean canFinish(int numCourses, int[][] prerequisites) {
//         int[] indegrees = new int[numCourses];
//         for(int[] cp : prerequisites) indegrees[cp[0]]++;
//         LinkedList<Integer> queue = new LinkedList<>();
//         for(int i = 0; i < numCourses; i++){
//             if(indegrees[i] == 0) queue.addLast(i);
//         }
//         while(!queue.isEmpty()) {
//             Integer pre = queue.removeFirst();
//             numCourses--;
//             for(int[] req : prerequisites) {
//                 if(req[1] != pre) continue;
//                 if(--indegrees[req[0]] == 0) queue.add(req[0]);
//             }
//         }
//         return numCourses == 0;
//     }

// 作者：jyd
// 链接：https://leetcode-cn.com/problems/course-schedule/solution/course-schedule-tuo-bu-pai-xu-bfsdfsliang-chong-fa/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
// }

// -----------------------------------------

// func canFinish(numCourses int, prerequisites [][]int) bool {
// 	// 学习回溯算法的操作
// 	// path array used array

// 	ans := false
// 	var dfs func(prerequisites [][]int, learned int, used []bool, path []int)
// 	dfs = func(prerequisites [][]int, learned int, used []bool, path []int) {
// 		if learned == numCourses || len(prerequisites) == 0 {
// 			ans = true
// 			return
// 		}
// 		for i:=0;i<len(prerequisites);i++ {
// 			want := prerequisites[i][0]
// 			need := prerequisites[i][1]
// 			if used[want] {
// 				// 循环学习 need 这门课程了
// 				continue
// 			}
// 			used[need] = true
// 			tmp := make([]int, len(path))
// 			copy(tmp, path)
// 			tmp = append(tmp, want)
// 			dfs(prerequisites, learned+1, used, tmp)
// 			used[need] = false
// 		}
// 	}
// 	used := make([]bool, numCourses)
// 	path := make([]int, 0)
// 	dfs(prerequisites, 0, used, path)
// 	return ans
// }
// @lc code=end

