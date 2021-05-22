/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type NestedIterator struct {
	List []*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{List: nestedList}
}

func (this *NestedIterator) Next() int {
	first := this.List[0]
	if first.IsInteger() {
			this.List = this.List[1:]
			return first.GetInteger()
	}
	list := first.GetList()
	this.List = append(list, this.List[1:]...)
	return this.Next()
}

func (this *NestedIterator) HasNext() bool {
	if len(this.List) == 0 {
			return false
	}
	first := this.List[0]
	if first.IsInteger() {
			return true
	}
	list := first.GetList()
	this.List = append(list, this.List[1:]...)
	return this.HasNext()
}