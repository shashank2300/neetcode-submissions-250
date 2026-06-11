// 1. Define a custom type for the heap elements
// Each element is an []int representing: [count, tweetId, followeeId, index]
type tweetHeap [][]int

// 2. Implement the 5 methods required by container/heap (sort.Interface + Push/Pop)
func (h tweetHeap) Len() int           { return len(h) }
func (h tweetHeap) Less(i, j int) bool { return h[i][0] < h[j][0] } // Min-heap on count
func (h tweetHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *tweetHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *tweetHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// ---------------------------------------------------------

type Twitter struct {
	count     int
	tweetMap  map[int][][]int      // userId -> list of [count, tweetId]
	followMap map[int]map[int]bool // userId -> set of followeeId
}

func Constructor() Twitter {
	return Twitter{
		count:     0,
		tweetMap:  make(map[int][][]int),
		followMap: make(map[int]map[int]bool),
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	if this.tweetMap[userId] == nil {
		this.tweetMap[userId] = make([][]int, 0)
	}
	// Note: Your count decrements, making recent tweets more negative. 
	// This works perfectly with our Min-Heap (Less function).
	this.tweetMap[userId] = append(this.tweetMap[userId], []int{this.count, tweetId})
	this.count--
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	res := make([]int, 0)

	// Initialize our standard Go heap
	minHeap := &tweetHeap{}
	heap.Init(minHeap)

	if this.followMap[userId] == nil {
		this.followMap[userId] = make(map[int]bool)
	}
	this.followMap[userId][userId] = true

	// Seed the heap with the most recent tweet from the user and everyone they follow
	for followeeId := range this.followMap[userId] {
		tweets := this.tweetMap[followeeId]
		if len(tweets) > 0 {
			index := len(tweets) - 1
			count, tweetId := tweets[index][0], tweets[index][1]
			heap.Push(minHeap, []int{count, tweetId, followeeId, index - 1})
		}
	}

	// Extract up to 10 most recent tweets
	for minHeap.Len() > 0 && len(res) < 10 {
		item := heap.Pop(minHeap).([]int)
		tweetId, followeeId, index := item[1], item[2], item[3]

		res = append(res, tweetId)

		// If this user has more tweets, push the next most recent one into the heap
		if index >= 0 {
			tweets := this.tweetMap[followeeId]
			nextCount, nextTweetId := tweets[index][0], tweets[index][1]
			heap.Push(minHeap, []int{nextCount, nextTweetId, followeeId, index - 1})
		}
	}

	return res
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if this.followMap[followerId] == nil {
		this.followMap[followerId] = make(map[int]bool)
	}
	this.followMap[followerId][followeeId] = true
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if this.followMap[followerId] != nil {
		delete(this.followMap[followerId], followeeId)
	}
}
