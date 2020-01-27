/*
 * @lc app=leetcode.cn id=901 lang=golang
 *
 * [901] 股票价格跨度
 */

// @lc code=start
type StockSpanner struct {
    stack []int
}


func Constructor() StockSpanner {
    return StockSpanner{}
}


func (this *StockSpanner) Next(price int) int {
	ans := 0
	for i:=len(this.stack)-1;i>=0;i-- {
		if this.stack[i] <= price {
			ans++
		} else {
			break
		}
	}
	// fmt.Println(this.stack, price, ans)
	this.stack = append(this.stack, price)
	return ans+1
}


/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
// @lc code=end

