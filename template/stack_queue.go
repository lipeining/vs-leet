package template

import "math"

type MinStack struct {
	stack    []int
	minStack []int
}

func NewMinStack() *MinStack {
	return &MinStack{stack: make([]int, 0), minStack: make([]int, 0)}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	if len(this.minStack) != 0 && this.minStack[len(this.minStack)-1] >= x {
		this.minStack = append(this.minStack, x)
	}
}
func (this *MinStack) Pop(x int) {
	n := len(this.stack)
	if len(this.minStack) != 0 && this.minStack[len(this.minStack)-1] == this.stack[n-1] {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
	this.stack = this.stack[:n-1]
}
func (this *MinStack) Top(x int) int {
	return this.stack[len(this.stack)-1]
}
func (this *MinStack) GetMin(x int) int {
	return this.minStack[len(this.minStack)-1]
}

// 单调递增栈
// for(int i = 0; i < T.size(); i++){
//   while(! stk.empty() && stk.top() > T[i]){
//     ​stk.pop();
//   }
//   stk.push(A[i]);
// }
// 单调递减栈
// for(int i = T.size() - 1; i >= 0; i--){
//   while(! stk.empty() && T[i] >= stk.top()){
//     stk.pop();
//   }
//   stk.push(i);
// }

// 单调栈，
func dailyTemperatures(temperatures []int) []int {
	// 单调递减的栈，记录需要升温的日期数
	n := len(temperatures)
	indices := make([]int, 0)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		for len(indices) != 0 {
			top := indices[len(indices)-1]
			if temperatures[top] >= temperatures[i] {
				break
			}
			indices = indices[:len(indices)-1]
			ans[top] = i - top
		}
		indices = append(indices, i)
	}
	return ans
}

// 实现单调队列，主要分为三个部分：

// 去尾操作 ：队尾元素出队列。当队列有新元素待入队，需要从队尾开始，
// 删除影响队列单调性的元素，维护队列的单调性。(删除一个队尾元素后，就重新判断新的队尾元素)
// 去尾操作结束后，将该新元素入队列。

// 删头操作 ：队头元素出队列。判断队头元素是否在待求解的区间之内，
// 如果不在，就将其删除。（这个很好理解呀，因为单调队列的队头元素就是待求解区间的极值）

// 取解操作 ：经过上面两个操作，取出 队列的头元素 ，就是 当前区间的极值 。
// 单调队列
// 头部维持窗口的开头区间内，
// 尾部需要不断 popBack and push 窗口的结尾区间。
func lestQueue(nums []int, m int) int {
	n := len(nums)
	queue := make([]int, 0)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = math.MaxInt32
	}
	for i := 0; i < n; i++ {
		// queue的尾部进行计算
		for len(queue) != 0 && dp[len(queue)-1] >= dp[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
		for len(queue) != 0 && queue[0] < i-m {
			queue = queue[1:]
		}
		dp[i] = dp[queue[0]] + nums[i]
	}
	return dp[n-1]
}
