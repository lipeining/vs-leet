/*
 * @lc app=leetcode.cn id=307 lang=golang
 *
 * [307] 区域和检索 - 数组可修改
 */

// @lc code=start
type NumArray struct {
	tree []int
	n int
}

func Constructor(nums []int) NumArray {
	tree := buildTree(nums)
	return NumArray{tree: tree, n: len(nums)}
}
func buildTree(nums []int) []int {
	length := len(nums)
	tree := make([]int, length*2)
	if length == 0 {
		return tree
	}
	for i:=length,j:=0;i<length*2;i++,j++ {
		tree[i] = nums[j]
	}
	for i:=length-1;i>=0;i-- {
		tree[i] = tree[i*2]+tree[i*2+1]
	}
	return tree
}

func (this *NumArray) Update(i int, val int) {
	i+= this.n
	this.tree[i] = val
	for i!= 0 {
		left := i
		right := i
		if i%2 == 0 {
			right = i+1
		} else {
			left = i-1
		}
		this.tree[i/2] = this.tree[left] + this.tree[right]
		i /= 2 
	}
}

func (this *NumArray) SumRange(i int, j int) int {
	l := i+this.n
	r := j+this.n
	sum := 0
	for l<=r {
		if l%2 == 1 {
			sum += this.tree[l]
			l++
		}
		if r%2 == 0 {
			sum += this.tree[r]
			r--
		}
		l /= 2
		r /= 2
	}
	return sum
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(i,val);
 * param_2 := obj.SumRange(i,j);
 */
// @lc code=end

// type NumArray struct {
// 	nums []int
// }

// func Constructor(nums []int) NumArray {
// 	return NumArray{nums: nums}
// }

// func (this *NumArray) Update(i int, val int) {
// 	this.nums[i] = val
// }

// func (this *NumArray) SumRange(i int, j int) int {
// 	sum := 0
// 	for l := i; l <= j; l++ {
// 		sum += this.nums[l]
// 	}
// 	return sum
// }

// type NumArray struct {
// 	sum []int
// }

// func Constructor(nums []int) NumArray {
// 	sum := make([]int, len(nums))
// 	for i := 0; i < len(nums); i++ {
// 		if i == 0 {
// 			sum[i] = nums[i]
// 		} else {
// 			sum[i] = sum[i-1] + nums[i]
// 		}
// 	}
// 	return NumArray{sum: sum}
// }

// func (this *NumArray) Update(i int, val int) {
// 	origin := this.sum[i] // 0
// 	if i > 0 {
// 		origin = this.sum[i] - this.sum[i-1]
// 	}
// 	diff := val - origin
// 	// fmt.Println(diff, this.sum)
// 	for start := i; start < len(this.sum); start++ {
// 		this.sum[start] += diff
// 	}
// 	// fmt.Println(diff, this.sum)
// }

// func (this *NumArray) SumRange(i int, j int) int {
// 	// fmt.Println("sum range", this.sum)
// 	sum := this.sum
// 	less := 0
// 	if i > 0 {
// 		less = sum[i-1]
// 	}
// 	return sum[j] - less
// }