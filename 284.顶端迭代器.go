/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */

// @lc code=start

type PeekingIterator struct {
	now *Iterator
	_hasNext bool
	_next int
}

func Constructor(iter *Iterator) *PeekingIterator {
	_hasNext := iter.hasNext()
	_next := 0
	if _hasNext {
		_next = iter.next()
	}	
	return &PeekingIterator{iter, _hasNext, _next}
}

func (this *PeekingIterator) hasNext() bool {
	return this._hasNext
}

func (this *PeekingIterator) next() int {
	origin := this._next
	to := 0
	this._hasNext = this.now.hasNext()
	if this._hasNext {
		to = this.now.next()
	}
	this._next = to
	return origin
}

func (this *PeekingIterator) peek() int {
	return this._next
}

// @lc code=end