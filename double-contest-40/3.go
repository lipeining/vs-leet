package main

import "fmt"

type FrontMiddleBackQueue struct {
	queue []int
}

func Constructor() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{queue: make([]int, 0)}
}

func (this *FrontMiddleBackQueue) PushFront(val int) {
	this.queue = append([]int{val}, this.queue...)
}

func (this *FrontMiddleBackQueue) PushMiddle(val int) {
	n := len(this.queue)
	half := n / 2
	left := make([]int, half)
	copy(left, this.queue[:half])
	left = append(left, val)
	this.queue = append(left, this.queue[half:]...)
	fmt.Println(this.queue)
}

func (this *FrontMiddleBackQueue) PushBack(val int) {
	this.queue = append(this.queue, val)
}

func (this *FrontMiddleBackQueue) PopFront() int {
	if len(this.queue) == 0 {
		return -1
	}
	front := this.queue[0]
	this.queue = this.queue[1:]
	return front
}

func (this *FrontMiddleBackQueue) PopMiddle() int {
	n := len(this.queue)
	if n == 0 {
		return -1
	}
	half := n / 2
	if n%2 == 0 {
		half--
	}
	mid := this.queue[half]
	right := this.queue[half+1:]
	this.queue = this.queue[:half]
	fmt.Println("popmid", this.queue, right)
	this.queue = append(this.queue, right...)
	return mid
}

func (this *FrontMiddleBackQueue) PopBack() int {
	n := len(this.queue)
	if n == 0 {
		return -1
	}
	end := this.queue[n-1]
	this.queue = this.queue[:n-1]
	return end
}

/**
 * Your FrontMiddleBackQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PushFront(val);
 * obj.PushMiddle(val);
 * obj.PushBack(val);
 * param_4 := obj.PopFront();
 * param_5 := obj.PopMiddle();
 * param_6 := obj.PopBack();
 */
