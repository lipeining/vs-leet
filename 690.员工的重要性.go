/**
 * Definition for Employee.
 * type Employee struct {
 *     Id int
 *     Importance int
 *     Subordinates []int
 * }
 */

func getImportance(employees []*Employee, id int) int {
	queue := make([]int, 0)
	queue = append(queue, id)
	ans := 0
	m := make(map[int]*Employee, 0)
	for i := 0; i < len(employees); i++ {
		m[employees[i].Id] = employees[i]
	}
	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			now := queue[i]
			e := m[now]
			ans += e.Importance
			for _,sub := range e.Subordinates {
				queue = append(queue, sub)
			}
		}
		queue = queue[size:]
	}
	return ans
}