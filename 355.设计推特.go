/*
 * @lc app=leetcode.cn id=355 lang=golang
 *
 * [355] 设计推特
 */

// @lc code=start

// An Item is something we manage in a priority queue.
type Item struct {
    value    int // The value of the item; arbitrary.
    priority int    // The priority of the item in the queue.
    // The index is needed by update and is maintained by the heap.Interface methods.
    index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
    // We want Pop to give us the highest, not lowest, priority so we use greater than here.
    // return pq[i].priority > pq[j].priority
    return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
    pq[i].index = i
    pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
    n := len(*pq)
    item := x.(*Item)
    item.index = n
    *pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    item.index = -1 // for safety
    *pq = old[0 : n-1]
    return item
}
type Post struct {
	tweetId int
	time int
}
type Twitter struct {
	posts map[int][]Post
	following map[int][]int
	// follower map[int][]int
}


/** Initialize your data structure here. */
func Constructor() Twitter {
	posts := make(map[int][]Post)
	following := make(map[int][]int)
    return Twitter{posts, following}
}


var Gcounter int = 1

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int)  {
	_,ok := this.posts[userId]
	post := Post{tweetId, Gcounter}
	Gcounter++
	if ok {
		this.posts[userId] = append(this.posts[userId], post)
	} else {
		this.posts[userId] = make([]Post, 1)
		this.posts[userId][0] = post
	}
}


/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	following,_ := this.following[userId]
	following = append(following, userId)
	K := 10
	for i:=0;i<len(following);i++ {
		posts,_ := this.posts[following[i]]
		for j:=0;j<len(posts);j++ {
			priority := posts[j].time
			value := posts[j].tweetId
			index := i*len(following)+j
			item := &Item{
				value:    value,
				priority: priority,
				index:    index,
			}
			heap.Push(&pq, item)
			if pq.Len() > K {
				heap.Pop(&pq)
			}
		}
	}
	max := 10 
	if max > pq.Len() {
		max = pq.Len()
	}
	ans := make([]int, max)
	i:=max-1
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		ans[i] = item.value
		i--
	}
	return ans
}


/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int)  {
	if followerId == followeeId {
		return
	}
	following,ok := this.following[followerId]
	if !ok {
		following = make([]int, 0)
	}
	i:=0;
	for  i<len(following) {
		if following[i] == followeeId {
			break
		}
		i++
	}
	if i != len(following) {
		return
	}
	this.following[followerId] = append(following, followeeId)
	fmt.Println("follow", this.following[followerId])
}


/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int)  {
	if followerId == followeeId {
		return
	}
	following,ok := this.following[followerId]
	if !ok {
		following = make([]int, 0)
	}
	i:=0;
	for  i<len(following) {
		if following[i] == followeeId {
			break
		}
		i++
	}
	if i == len(following) {
		return
	}
	this.following[followerId] = append(following[:i], following[i+1:]...)
	fmt.Println("unfollow", this.following[followerId])
}


/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
// @lc code=end

