package main

func main() {

}

func countQuadruplets(nums []int) int {
	// sort.Ints(nums)
	n := len(nums)
	ans := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				for l := k + 1; l < n; l++ {
					if nums[i]+nums[j]+nums[k] == nums[l] {
						ans++
					}
				}
			}
		}
	}
	return ans
}
